package usecase

import "github.com/kindai-csg/D-Chat/domain"

type MessageRepository interface{
	Create(domain.Message) (domain.Message, error)
}