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

