package gateway

import "github.com/kindai-csg/D-Chat/domain"

type ChannelInfoOutput struct {
	Name    string `json:"name"`
	Concept string `json:"concept"`
}

func CreateChannelInfoOutputFromChannel(channel domain.Channel) ChannelInfoOutput {
	return ChannelInfoOutput{
		channel.Name,
		channel.Concept,
	}
}
