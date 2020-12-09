package database

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/kindai-csg/D-Chat/domain"
	"github.com/kindai-csg/D-Chat/domain/enum"
)

type UserRepository struct {
	mongoHandler   MongoHandler
	collectionName string
}

func NewUserRepository(mongoHandler MongoHandler) *UserRepository {
	userRepository := UserRepository{
		mongoHandler:   mongoHandler,
		collectionName: "Users",
	}
	userRepository.createIndex()
	return &userRepository
}

func (repository *UserRepository) createIndex() {
	repository.mongoHandler.CreateIndex(repository.collectionName, []KV{{"user_id", 1}}, []KV{{"unique", true}})
	repository.mongoHandler.CreateIndex(repository.collectionName, []KV{{"mail", 1}}, []KV{{"unique", true}})
}

func (repository *UserRepository) Create(user domain.User) (domain.User, error) {
	doc := []KV{
		{"user_id", user.UserId},
		{"name", user.Name},
		{"password", user.Password},
		{"mail", user.Mail},
		{"bio", user.Bio},
		{"status", user.Status},
		{"status_text", user.StatusText},
		{"auth", user.Auth},
	}
	id, err := repository.mongoHandler.Insert(repository.collectionName, doc)
	user.Id = id
	return user, err
}

func (repository *UserRepository) Update(user domain.User) (domain.User, error) {
	query := []KV{
		{"_id", user.Id},
	}
	update := []KV{
		{"$set", []KV{
			{"user_id", user.UserId},
			{"name", user.Name},
			{"mail", user.Mail},
			{"bio", user.Bio},
			{"status", user.Status},
			{"status_text", user.StatusText},
		}},
	}
	err := repository.mongoHandler.Update(repository.collectionName, query, update)
	return user, err
}

func (repository *UserRepository) Delete(id string) error {
	query := []KV{
		{"_id", id},
	}
	_, err := repository.mongoHandler.Delete(repository.collectionName, query)
	return err
}

func (repository *UserRepository) GetAll() ([]domain.User, error) {
	query := []KV{}
	raw, err := repository.mongoHandler.Find(repository.collectionName, query)
	if err != nil {
		return nil, err
	}
	users := []domain.User{}
	for _, kv := range raw {
		for _, u := range kv {
			user := domain.User{}
			switch u.Key {
			case "_id":
				user.Id = u.Value.(string)
			case "user_id":
				user.UserId = u.Value.(string)
			case "password":
				user.Password = u.Value.(string)
			case "mail":
				user.Mail = u.Value.(string)
			case "bio":
				user.Bio = u.Value.(string)
			case "status":
				user.Status = u.Value.(string)
			case "status_text":
				user.StatusText = u.Value.(string)
			case "auth":
				user.Auth = u.Value.(domain.AuthType)
			}
			users = append(users, user)
		}
	}
	return users, nil
}

func (repository *UserRepository) Authenticate(user_id string, password string) error {
	query := []KV{
		{"user_id", user_id},
	}
	user, err := repository.mongoHandler.FindOne(repository.collectionName, query)
	if err != nil {
		return err
	}
	for _, kv := range user {
		if kv.Key == "password" {
			hashTypeStrIndex := strings.Index(kv.Value.(string), "}") + 1
			if kv.Value.(string) == repository.hashPassword(password, kv.Value.(string)[0:hashTypeStrIndex]) {
				return nil
			}
			break
		}
	}
	return errors.New("wrong id or password")
}

func (repository *UserRepository) hashPassword(password string, hahsTypeStr string) string {
	switch hahsTypeStr {
	case enum.MD5.String():
		md5 := md5.Sum([]byte(password))
		password = enum.MD5.String() + base64.StdEncoding.EncodeToString(md5[:])
	default:
		md5 := md5.Sum([]byte(password))
		password = enum.MD5.String() + base64.StdEncoding.EncodeToString(md5[:])
	}
	return password
}
