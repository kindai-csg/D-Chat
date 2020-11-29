package database

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/kindai-csg/D-Chat/domain"
	"github.com/kindai-csg/D-Chat/domain/enum"
)

type MessageRepository struct {
	mongoHandler   MongoHandler
	collectionName string
}

func NewMessageRepository(mongoHandler MongoHandler) *MessageRepository {
	messageRepository := MessageRepository{
		mongoHandler:   mongoHandler,
		collectionName: "Messages",
	}
	messageRepository.createIndex()
	return &messageRepository
}

func (repository *UserRepository) createIndex() {
	repository.mongoHandler.CreateIndex(repository.collectionName, []KV{{"user_id", 1}}, []KV{{"unique", true}})
	repository.mongoHandler.CreateIndex(repository.collectionName, []KV{{"mail", 1}}, []KV{{"unique", true}})
}

func (repository *MessageRepository) Create(message domain.Message) (domain.Message, error) {
	doc := []KV{
		{"user_id", message.UserId},
		{"name", message.Name},
		{"password", message.Password},
		{"mail", message.Mail},
		{"bio", message.Bio},
		{"status", message.Status},
		{"status_text", message.StatusText},
		{"auth", message.Auth},
	}
	id, err := repository.mongoHandler.Insert(repository.collectionName, doc)
	message.Id = id
	return message, err
}