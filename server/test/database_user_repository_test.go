package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kindai-csg/D-Chat/domain"
	"github.com/kindai-csg/D-Chat/interfaces/database"
	mock "github.com/kindai-csg/D-Chat/test/mock_database"
)

func createUsersIndex(m *mock.MockMongoHandler) {
	m.EXPECT().CreateIndex(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
	m.EXPECT().CreateIndex(gomock.Any(), gomock.Any(), gomock.Any()).Return(nil)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	mongoHandler := mock.NewMockMongoHandler(ctrl)
	createUsersIndex(mongoHandler)
	repository := database.NewUserRepository(mongoHandler)

	collectionName := "Users"
	argUser := domain.User{}
	id := "test"
	mongoHandler.EXPECT().Insert(collectionName, gomock.Any()).Return(id, nil)
	u, err := repository.Create(argUser)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
	if u.Id != id {
		t.Errorf("Expectation: id is test")
	}

	mongoHandler.EXPECT().Insert(collectionName, gomock.Any()).Return(id, errors.New(""))
	_, err = repository.Create(argUser)
	if err == nil {
		t.Errorf("Expectation: return error")
	}
}
