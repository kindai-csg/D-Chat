package gateway

import "github.com/kindai-csg/D-Chat/domain"

type UserInfoOutput struct {
	Id         string          `json:"id"`
	UserId     string          `json:"user_id"`
	Name       string          `json:"name"`
	Mail       string          `json:"mail"`
	Bio        string          `json:"bio"`
	Status     string          `json:"status"`
	StatusText string          `json:"statusText"`
	Auth       domain.AuthType `json:"auth"`
}

func CreateUserInfoOutputFromUser(user domain.User) UserInfoOutput {
	return UserInfoOutput{
		user.Id,
		user.UserId,
		user.Name,
		user.Mail,
		user.Bio,
		user.Status,
		user.StatusText,
		user.Auth,
	}
}
