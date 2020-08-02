package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kindai-csg/D-Chat/domain"
	mock "github.com/kindai-csg/D-Chat/test/mock_usecase"
	"github.com/kindai-csg/D-Chat/usecase"
)

func TestAddNewMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	messageRepository := mock.NewMockMessageRepository(ctrl)
	interactor := usecase.NewMessageInteractor(messageRepository)

	message := domain.Message{}
	messageRepository.EXPECT().StoreMessage(message).Return(message, errors.New(""))
	_, err := interactor.AddNewMessage(message)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	messageRepository.EXPECT().StoreMessage(message).Return(message, nil)
	_, err = interactor.AddNewMessage(message)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
}

func TestUpdateText(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	messageRepository := mock.NewMockMessageRepository(ctrl)
	interactor := usecase.NewMessageInteractor(messageRepository)

	message := domain.Message{}
	messageRepository.EXPECT().UpdateText(message).Return(message, errors.New(""))
	_, err := interactor.UpdateText(message)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	messageRepository.EXPECT().UpdateText(message).Return(message, nil)
	_, err = interactor.UpdateText(message)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
}

func TestDeleteMessage(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	messageRepository := mock.NewMockMessageRepository(ctrl)
	interactor := usecase.NewMessageInteractor(messageRepository)

	message := domain.Message{Id: "test"}
	messageRepository.EXPECT().DeleteMessage(message.Id).Return(errors.New(""))
	err := interactor.DeleteMessage(message)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	messageRepository.EXPECT().DeleteMessage(message.Id).Return(nil)
	err = interactor.DeleteMessage(message)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
}
