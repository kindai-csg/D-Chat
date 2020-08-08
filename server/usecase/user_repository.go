package usecase

import "github.com/kindai-csg/D-Chat/domain"

type UserRepository interface {
	Create(domain.User) (domain.User, error)
	Update(domain.User) (domain.User, error)
	Delete(string) error
	GetAll() ([]domain.User, error)
	Authenticate(domain.User) error
}
