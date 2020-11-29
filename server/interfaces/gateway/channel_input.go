package gateway

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type ChannelInput struct {
	Id      string `json:"id"`
	Name    string `jason:"name"`
	Concept string `json:"concept"`
}

func (channel *ChannelInput) GetChannel() domain.Channel {
	return domain.Channel{
		Id:      channel.Id,
		Name:    channel.Name,
		Concept: channel.Concept,
	}
}
