package gateway

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type ChannelInput struct {
	Id             string `json:"id"`
	ChannelId      string `json:"channel_id"`
	ChannelName    string `jason:"channel_name"`
	ChannelConcept string `json:"channel_concept"`
}

func (channel *ChannelInput) GetChannel() domain.Channel {
	return domain.Channel{
		Id:             channel.Id,
		ChannelId:      channel.ChannelId,
		ChannelName:    channel.ChannelName,
		ChannelConcept: channel.ChannelConcept,
	}
}
