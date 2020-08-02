package usecase

import "github.com/kindai-csg/D-Chat/domain"

type MessageInteractor struct {
	MessageRepository MessageRepository
}

func NewMessageInteractor(messageRepository MessageRepository) *MessageInteractor {
	messageInteractor := MessageInteractor{
		MessageRepository: messageRepository,
	}
	return &messageInteractor
}

func (interactor *MessageInteractor) AddNewMessage(message domain.Message) (domain.Message, error) {
	m, err := interactor.MessageRepository.StoreMessage(message)
	if err != nil {
		return domain.Message{}, err
	}
	return m, nil
}

func (interactor *MessageInteractor) UpdateText(message domain.Message) (domain.Message, error) {
	m, err := interactor.MessageRepository.UpdateText(message)
	if err != nil {
		return domain.Message{}, err
	}
	return m, nil
}

func (interactor *MessageInteractor) DeleteMessage(message domain.Message) error {
	err := interactor.MessageRepository.DeleteMessage(message.Id)
	if err != nil {
		return err
	}
	return nil
}
