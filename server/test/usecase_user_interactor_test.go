package test

import (
	"errors"
	"testing"

	"github.com/golang/mock/gomock"
	"github.com/kindai-csg/D-Chat/domain"
	mock "github.com/kindai-csg/D-Chat/test/mock_usecase"
	"github.com/kindai-csg/D-Chat/usecase"
)

func TestAddNewUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock.NewMockUserRepository(ctrl)
	interactor := usecase.NewUserInteractor(userRepository)

	user := domain.User{}
	userRepository.EXPECT().StoreUser(user).Return(user, errors.New(""))
	_, err := interactor.AddNewUser(user)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	user = domain.User{Password: "test"}
	userRepository.EXPECT().StoreUser(user).Return(user, nil)
	rUser, err := interactor.AddNewUser(user)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
	if rUser.Password != "" {
		t.Errorf("Expectation: hide password")
	}
}

func TestDeleteUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock.NewMockUserRepository(ctrl)
	interactor := usecase.NewUserInteractor(userRepository)

	user := domain.User{Id: "test"}
	userRepository.EXPECT().DeleteUser(user.Id).Return(nil)
	err := interactor.DeleteUser(user)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
	userRepository.EXPECT().DeleteUser(user.Id).Return(errors.New(""))
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
	userRepository.EXPECT().UpdateUser(user).Return(user, errors.New(""))
	_, err := interactor.UpdateUser(user)
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	user = domain.User{Password: "test"}
	userRepository.EXPECT().UpdateUser(user).Return(user, nil)
	rUser, err := interactor.UpdateUser(user)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
	if rUser.Password != "" {
		t.Errorf("Expectation: hide password")
	}
}

func TestAuthenticateUser(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	userRepository := mock.NewMockUserRepository(ctrl)
	interactor := usecase.NewUserInteractor(userRepository)

	user := domain.User{}
	userRepository.EXPECT().AuthenticateUser(user).Return(nil)
	err := interactor.AuthenticateUser(user)
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
	userRepository.EXPECT().AuthenticateUser(user).Return(errors.New(""))
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
	userRepository.EXPECT().GetAllUsers().Return(users, errors.New(""))
	_, err := interactor.GetAllUsers()
	if err == nil {
		t.Errorf("Expectation: return error")
	}

	userRepository.EXPECT().GetAllUsers().Return(users, nil)
	users, err = interactor.GetAllUsers()
	if err != nil {
		t.Errorf("Expectation: return nil")
	}
	for _, user := range users {
		if user.Password != "" {
			t.Errorf("Expectation: hide password")
		}
	}
}
