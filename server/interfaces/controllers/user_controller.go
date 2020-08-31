package controllers

import (
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/usecase"
)

type UserController struct {
	interactor usecase.UserInteractor
}

func NewUserController(mongoHandler database.MongoHandler) *UserController {
	userController := UserController{
		interactor: *usecase.NewUserInteractor(database.NewUserRepository(mongoHandler)),
	}
}
