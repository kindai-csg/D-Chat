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

func (controller *MessageController) CreateMessage(c Context) {
	input := gateway.MessageInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(400, gateway.StatusMessageOutput{false, "Parameter is invalid."})
		return
	}
	_, err = controller.interactor.Create(input.GetMessage())
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(500, gateway.StatusMessageOutput{false, "Failed to create message."})
		return
	}
	c.JSON(200, gateway.StatusMessageOutput{true, "success!"})
}
