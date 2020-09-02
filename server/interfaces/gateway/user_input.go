package gateway

import "github.com/kindai-csg/D-Chat/domain"

type UserInput struct {
	UserId   string `json:"user_id"`
	Name     string `json:"name"`
	Password string `json:"password"`
	Mail     string `json:"mail"`
}

func CreateUserFromUserInput(user UserInput) domain.User {
	return domain.User{
		UserId:   user.UserId,
		Name:     user.Name,
		Password: user.Password,
		Mail:     user.Mail,
	}
}
