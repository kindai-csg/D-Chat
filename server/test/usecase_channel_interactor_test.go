package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kindai-csg/D-Chat/domain"
	mock "github.com/kindai-csg/D-Chat/test/mock_usecase"
	"github.com/kindai-csg/D-Chat/usecase"
)

func TestChannelInteractorCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	channelRepository := mock.NewMockChannelRepository(ctrl)
	interactor := usecase.NewChannelInteractor(channelRepository)

	channel := domain.Channel{}
	channelRepository.EXPECT().Create(channel).Return(channel, errors.New(""))
	_, err := interactor.CreateChannel(channel)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	channelRepository.EXPECT().Create(channel).Return(channel, nil)
	_, err = interactor.CreateChannel(channel)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
}
