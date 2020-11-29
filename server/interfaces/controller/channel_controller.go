package controller

import (
	"github.com/kindai-csg/D-Chat/interfaces"
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/interfaces/gateway"
	"github.com/kindai-csg/D-Chat/usecase"
)

type ChannelController struct {
	logger     interfaces.Logger
	interactor usecase.ChannelInteractor
}

func NewChannelController(logger interfaces.Logger, mongoHandler database.MongoHandler) *ChannelController {
	channelContoroller := ChannelController{
		logger:     logger,
		interactor: *usecase.NewChannelInteractor(database.NewChannelRepository(mongoHandler)),
	}
	return &channelContoroller
}

func (controller *ChannelController) CreateChannel(c Context) {
	input := gateway.ChannelInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(400, gateway.StatusMessageOutput{false, "Parameter is invalid."})
		return
	}
	channel, err := controller.interactor.CreateChannel(input.GetChannel())
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(500, gateway.StatusMessageOutput{false, "Failed to create account."})
		return
	}
	c.JSON(200, gateway.CreateChannelInfoOutputFromChannel(channel))
}
