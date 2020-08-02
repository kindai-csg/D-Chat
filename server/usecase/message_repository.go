package usecase

import "github.com/kindai-csg/D-Chat/domain"

type MessageRepository interface {
	StoreMessage(domain.Message) (domain.Message, error)
	DeleteMessage(string) error
	UpdateText(domain.Message) (domain.Message, error)
}
