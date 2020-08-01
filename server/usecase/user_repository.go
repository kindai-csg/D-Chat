package usecase

import "github.com/kindai-csg/D-Chat/domain"

type UserRepository interface {
	StoreUser(domain.User) (domain.User, error)
	UpdateUser(domain.User) (domain.User, error)
	DeleteUser(string) error
	GetAllUsers() ([]domain.User, error)
	FindUserById(string) (domain.User, error)
	AuthenticateUser(domain.User) error
}
