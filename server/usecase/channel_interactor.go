package usecase

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type ChannelInteractor struct {
	ChannelRepository ChannelRepository
}

func NewChannelInteractor(channelRepository ChannelRepository) *ChannelInteractor {
	channelInteractor := ChannelInteractor{
		ChannelRepository: channelRepository,
	}
	return &channelInteractor
}

func (interactor *ChannelInteractor) CreateChannel(channel domain.Channel) (domain.Channel, error) {
	u, err := interactor.ChannelRepository.Create(channel)
	if err != nil {
		return domain.Channel{}, err
	}
	return u, nil
}

/*
func (interactor *ChannelInteractor) DeleteChannel(channel domain.Channel) error {
	err := interactor.ChannelRepository.Delete(channel)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *ChannelInteractor) UpdateChannel(channel domain.Channel) (domain.Channel, error) {
	u, err := interactor.ChannelRepository.Update(channel)
	if err != nil {
		return domain.Channel{}, err
	}
	return u, nil
}
*/
