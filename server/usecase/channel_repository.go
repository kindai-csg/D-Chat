package usecase

import "github.com/kindai-csg/D-Chat/domain"

type ChannelRepository interface {
	Create(domain.Channel) (domain.Channel, error)
	/*Delete(domain.Channel) error
	Update(domain.Channel) (domain.Channel, error)
	GetAll() ([]domain.Channel, error)*/
}
