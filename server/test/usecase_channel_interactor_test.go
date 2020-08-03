package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kindai-csg/D-Chat/domain"
	mock "github.com/kindai-csg/D-Chat/test/mock_usecase"
	"github.com/kindai-csg/D-Chat/usecase"
)

func TestCreateChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	channelRepository := mock.NewMockChannelRepository(ctrl)
	interactor := usecase.NewChannelInteractor(channelRepository)

	channel := domain.Channel{}
	channelRepository.EXPECT().StoreChannel(channel).Return(channel, errors.New(""))
	_, err := interactor.CreateChannel(channel)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	channelRepository.EXPECT().StoreChannel(channel).Return(channel, nil)
	_, err = interactor.CreateChannel(channel)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
}

func TestDeleteChannel(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	channelRepository := mock.NewMockChannelRepository(ctrl)
	interactor := usecase.NewChannelInteractor(channelRepository)

	channel := domain.Channel{Id: "test"}
	channelRepository.EXPECT().DeleteChannel(channel.Id).Return(errors.New(""))
	err := interactor.DeleteChannel(channel)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	channelRepository.EXPECT().DeleteChannel(channel.Id).Return(nil)
	err = interactor.DeleteChannel(channel)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
}
