package controller

import (
	"github.com/kindai-csg/D-Chat/domain"
	"github.com/kindai-csg/D-Chat/interfaces"
	"github.com/kindai-csg/D-Chat/interfaces/database"
	"github.com/kindai-csg/D-Chat/interfaces/gateway"
	"github.com/kindai-csg/D-Chat/usecase"
)

type TokenGenerator interface {
	Create(domain.TokenData) string
	Authentication(string) (domain.TokenData, error)
}

type TokenController struct {
	logger         interfaces.Logger
	tokenGenerator TokenGenerator
	interactor     usecase.UserInteractor
}

func NewTokenController(logger interfaces.Logger, tokenGenerator TokenGenerator, mongoHandler database.MongoHandler) *TokenController {
	tokenController := TokenController{
		logger:         logger,
		tokenGenerator: tokenGenerator,
		interactor:     *usecase.NewUserInteractor(database.NewUserRepository(mongoHandler)),
	}
	return &tokenController
}

func (controller *TokenController) CreateToken(c Context) {
	input := gateway.LoginInput{}
	err := c.ShouldBindJSON(&input)
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(400, gateway.StatusMessageOutput{false, "Parameter is invalid."})
		return
	}
	err = controller.interactor.AuthenticateUser(input.UserId, input.Password)
	if err != nil {
		controller.logger.Error(err.Error())
		c.JSON(401, gateway.StatusMessageOutput{false, "wrong id or password"})
		return
	}
	tokenData := domain.TokenData{
		UserId: input.UserId,
	}
	token := controller.tokenGenerator.Create(tokenData)
	c.JSON(200, gateway.NewTokenOutput(token))
}
