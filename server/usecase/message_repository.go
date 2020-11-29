package usecase

import "github.com/kindai-csg/D-Chat/domain"

type MessageRepository interface{
	Create(domain.User) (domain.User, error)
}