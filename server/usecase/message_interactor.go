package usecase

import {
	"github.com/kindai-csg/D-Chat/domain"
}

type MessageInteractor struct{
	MessageRepository MessageRepository
}

func NewMessageInteractor(messageRepository MessageRepository) *MessageInteractor {
	messageInteractor := MessageInteractor{
		MessageRepository: messageRepository,
	}
	return &messageInteractor
}

func (interactor *MessageInteractor) Create(message domain.Message) (domain.Message, error) {
	u, err := interactor.MessageRepository.Create(message)
	if err != nil {
		return domain.Message{}, err
	}
	return u, nil
}