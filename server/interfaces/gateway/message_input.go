package gateway

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type MessageInput struct {
	UserId string `json:"user_id"`
	Body   string `json:"body"`
}

func (message *MessageInput) GetMessage() domain.Message {
	return domain.Message{
		UserId: message.UserId,
		Body:   message.Body,
	}
}
