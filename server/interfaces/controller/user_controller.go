package controller

import (
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/interfaces/gateway"
	"github.com/kindai-csg/D-Chat/usecase"
)

type UserController struct {
	interactor usecase.UserInteractor
}

func NewUserController(mongoHandler database.MongoHandler) *UserController {
	userController := UserController{
		interactor: *usecase.NewUserInteractor(database.NewUserRepository(mongoHandler)),
	}
	return &userController
}

func (controller *UserController) CreateUser(c Context) {
	input := gateway.UserInput{}
	c.Bind(&input)
	user, err := controller.interactor.Create(gateway.CreateUserFromUserInput(input))
	if err != nil {
		c.JSON(500, gateway.StatusMessageOutput{false, "Failed to create account."})
		return
	}
	c.JSON(200, gateway.CreateUserInfoOutputFromUser(user))
}
