package usecase

import "github.com/kindai-csg/D-Chat/domain"

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
	c, err := interactor.ChannelRepository.StoreChannel(channel)
	if err != nil {
		return domain.Channel{}, err
	}
	return c, err
}

func (interactor *ChannelInteractor) DeleteChannel(channel domain.Channel) error {
	err := interactor.ChannelRepository.DeleteChannel(channel.Id)
	if err != nil {
		return err
	}
	return nil
}
