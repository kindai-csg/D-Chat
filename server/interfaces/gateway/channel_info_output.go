package gateway

import "github.com/kindai-csg/D-Chat/domain"

type ChannelInfoOutput struct {
	Id      string `josn:"id"`
	Name    string `json:"name"`
	Concept string `json:"concept"`
}

func CreateChannelInfoOutputFromChannel(channel domain.Channel) ChannelInfoOutput {
	return ChannelInfoOutput{
		channel.Id,
		channel.Name,
		channel.Concept,
	}
}
