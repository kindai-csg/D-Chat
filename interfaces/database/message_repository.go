package database

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type MessageRepository struct {
	mongoHandler   MongoHandler
	collectionName string
}

func NewMessageRepository(mongoHandler MongoHandler) *MessageRepository {
	messageRepository := MessageRepository{
		mongoHandler:   mongoHandler,
		collectionName: "Messages",
	}
	return &messageRepository
}

func (repository *MessageRepository) Create(message domain.Message) (domain.Message, error) {
	doc := []KV{
		{"user_id", message.UserId},
		{"body", message.Body},
	}

	_, err := repository.mongoHandler.Insert(repository.collectionName, doc)
	return message, err
}
