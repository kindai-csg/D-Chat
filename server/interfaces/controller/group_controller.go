package controller

import (
	"github.com/kindai-csg/D-Chat/interfaces"
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/interfaces/gateway"
	"github.com/kindai-csg/D-Chat/usecase"
)

type GroupController struct {
	logger     interfaces.logger
	interactor usecase.GroupInteractor
}

func NewGroupController(logger interfaces.Logger, mongoHandler database.ongoHandler) *GroupController {
	groupController := GroupController{
		logger:     logger,
		interactor: *usecase.NewGroupInteractor(database.NewGroupRepository(mongoHandler)),
	}
	return &groupController
}

func (controller *GroupController) CreateGroup(c Context) {
	input := gateway.GroupInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(400, gateway.StatusGroupOutput{false, "Parameter is vailed."})
		return
	}
	_, err = controller.interactor.Create(input.GetGroup())
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(500, gateway.StatusGroupOutput{false, "Failed to create group"})
		return
	}
	c.JSON(200, gateway.StatusGroupOutput{true, "success!"})
}
