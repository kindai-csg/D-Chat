package controller

import (
	"github.com/kindai-csg/D-Chat/interfaces"
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/interfaces/gateway"
	"github.com/kindai-csg/D-Chat/usecase"
)

type GroupController struct {
	logger     interfaces.Logger
	interactor usecase.GroupInteractor
}

func NewGroupController(logger interfaces.Logger, mongoHandler database.MongoHandler) *GroupController {
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
		c.JSON(400, gateway.StatusMessageOutput{false, "Parameter is vailed."})
		return
	}
	group, err = controller.interactor.Create(input.GetGroup())
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(500, gateway.StatusMessageOutput{false, "Failed to create group"})
		return
	}
	c.JSON(200, gateway.CreateGroupInfoOutputFromGroup(group))
}
