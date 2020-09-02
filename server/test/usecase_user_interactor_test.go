package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kindai-csg/D-Chat/domain"
	mock "github.com/kindai-csg/D-Chat/test/mock_usecase"
	"github.com/kindai-csg/D-Chat/usecase"
)

func TestUserInteractorCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock.NewMockUserRepository(ctrl)
	interactor := usecase.NewUserInteractor(userRepository)

	user := domain.User{}
	userRepository.EXPECT().Create(user).Return(user, errors.New(""))
	_, err := interactor.Create(user)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	user = domain.User{Password: "test"}
	userRepository.EXPECT().Create(user).Return(user, nil)
	_, err = interactor.Create(user)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock.NewMockUserRepository(ctrl)
	interactor := usecase.NewUserInteractor(userRepository)

	user := domain.User{Id: "test"}
	userRepository.EXPECT().Delete(user.Id).Return(nil)
	err := interactor.DeleteUser(user)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
	userRepository.EXPECT().Delete(user.Id).Return(errors.New(""))
	err = interactor.DeleteUser(user)
	if err == nil {
		t.Errorf("Expectation: return error")
	}
}

func TestUpdateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock.NewMockUserRepository(ctrl)
	interactor := usecase.NewUserInteractor(userRepository)

	user := domain.User{}
	userRepository.EXPECT().Update(user).Return(user, errors.New(""))
	_, err := interactor.UpdateUser(user)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	user = domain.User{Password: "test"}
	userRepository.EXPECT().Update(user).Return(user, nil)
	_, err = interactor.UpdateUser(user)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
}

func TestAuthenticateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock.NewMockUserRepository(ctrl)
	interactor := usecase.NewUserInteractor(userRepository)

	user := domain.User{}
	userRepository.EXPECT().Authenticate(user).Return(nil)
	err := interactor.AuthenticateUser(user)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
	userRepository.EXPECT().Authenticate(user).Return(errors.New(""))
	err = interactor.AuthenticateUser(user)
	if err == nil {
		t.Errorf("Expectation: return error")
	}
}

func TestGetAllUsers(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock.NewMockUserRepository(ctrl)
	interactor := usecase.NewUserInteractor(userRepository)

	users := []domain.User{domain.User{Password: "test"}, domain.User{Password: "test"}, domain.User{Password: "test"}}
	userRepository.EXPECT().GetAll().Return(users, errors.New(""))
	_, err := interactor.GetAllUsers()
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	userRepository.EXPECT().GetAll().Return(users, nil)
	_, err = interactor.GetAllUsers()
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
}
