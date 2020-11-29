package database

import (
	"crypto/md5"
	"encoding/base64"
	"errors"
	"strings"

	"github.com/kindai-csg/D-Chat/domain"
	"github.com/kindai-csg/D-Chat/domain/enum"
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
	messageRepository.createIndex()
	return &messageRepository
}

func (repository *MessageRepository) Create(message domain.Message) (domain.Message, error) {
	doc := []KV{
		{"status", message.Status},
		{"message", message.Message}
	}
	id, err := repository.mongoHandler.Insert(repository.collectionName, doc)
	message.Id = id
	return message, err
}