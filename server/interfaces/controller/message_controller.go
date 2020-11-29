package controller

import (
	"github.com/kindai-csg/D-Chat/interfaces"
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/interfaces/gateway"
	"github.com/kindai-csg/D-Chat/usecase"
)

type MessageController struct {
	logger     interfaces.Logger
	interactor usecase.MessageInteractor
}

func NewMessageController(logger interfaces.Logger, mongoHandler database.MongoHandler) *MessageController {
	messageController := MessageController{
		logger:     logger,
		interactor: *usecase.NewMessageInteractor(database.NewMessageRepository(mongoHandler)),
	}
	return &messageController
}