package controller

import (
	"github.com/kindai-csg/D-Chat/interfaces"
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/interfaces/gateway"
	"github.com/kindai-csg/D-Chat/usecase"
)

type ChannelController struct {
	logger     interfaces.Logger
	interactor usecase.UserInteractor
}

func NewChannelController(logger interfaces.Logger, mongoHandler database.MongoHandler) *ChannelController {
	channelContoroller := ChannelContoroller{
		logger : logger,
		interactor: *usecase.NewChannelInteractor(database.NewChannelRepository(mongoHandler)),
	}
	return &channelContoroller
}

func (contoroller *ChannelController) CreateChannel(c Context){
	func (controller *UserController) CreateUser(c Context) {
		input := gateway.UserInput{}
		err := c.ShouldBindJSON(&input)
		if err != nil {
			controller.logger.Error(err.Error())
			c.JSON(400, gateway.StatusMessageOutput{false, "Parameter is invalid."})
			return
		}
		user, err := controller.interactor.Create(input.GetUser())
		if err != nil {
			controller.logger.Error(err.Error())
			c.JSON(500, gateway.StatusMessageOutput{false, "Failed to create account."})
			return
		}
		c.JSON(200, gateway.CreateUserInfoOutputFromUser(user))
	}
}