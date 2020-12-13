package controller

import (
	"github.com/kindai-csg/D-Chat/interfaces"
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/interfaces/gateway"
	"github.com/kindai-csg/D-Chat/usecase"
)

type UserController struct {
	logger     interfaces.Logger
	interactor usecase.UserInteractor
}

func NewUserController(logger interfaces.Logger, mongoHandler database.MongoHandler) *UserController {
	userController := UserController{
		logger:     logger,
		interactor: *usecase.NewUserInteractor(database.NewUserRepository(mongoHandler)),
	}
	return &userController
}

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
