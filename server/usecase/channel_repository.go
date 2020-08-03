package usecase

import "github.com/kindai-csg/D-Chat/domain"

type ChannelRepository interface {
	StoreChannel(domain.Channel) (domain.Channel, error)
	DeleteChannel(string) error
}
