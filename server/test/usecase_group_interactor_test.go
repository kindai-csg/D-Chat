package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kindai-csg/D-Chat/domain"
	mock "github.com/kindai-csg/D-Chat/test/mock_usecase"
	"github.com/kindai-csg/D-Chat/usecase"
)

func TestGroupInteractorCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	groupRepository := mock.NewMockGroupRepository(ctrl)
	interactor := usecase.NewGroupInteractor(groupRepository)

	group := domain.Group{}
	groupRepository.EXPECT().Create(group).Return(group, errors.New(""))

	_, err := interactor.Create(group)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	group = domain.Group{}
}
