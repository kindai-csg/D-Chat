package gateway

import (
	"crypto/md5"
	"encoding/base64"

	"github.com/kindai-csg/D-Chat/domain"
)

type UserInput struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

func (user *UserInput) GetUser() domain.User {
	md5 := md5.Sum([]byte(user.Password))
	password := base64.StdEncoding.EncodeToString(md5[:])
	return domain.User{
		UserId:   user.UserId,
		Name:     user.Name,
		Password: "{MD5}" + password,
		Mail:     user.Mail,
	}
}
