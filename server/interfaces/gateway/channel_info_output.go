package gateway

import "github.com/kindai-csg/D-Chat/domain"

type ChannelInfoOutput struct {
	Id             string `josn:"id"`
	ChannelId      string `json:"channel_id"`
	ChannelName    string `json:"channel_name"`
	ChannelConcept string `json:"channel_concept"`
}

func CreateChannelInfooutputFormChannel(channel domain.Channel) ChannelInfoOutput {
	return ChannelInfoOutput{
		channel.Id,
		channel.Channelid,
		channel.ChannelName,
		channel.ChannelConcept,
	}
}
