package usecase

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type UserInteractor struct {
	UserRepository UserRepository
}

func NewUserInteractor(userRepository UserRepository) *UserInteractor {
	userInteractor := UserInteractor{
		UserRepository: userRepository,
	}
	return &userInteractor
}

func (interactor *UserInteractor) AddNewUser(user domain.User) (domain.User, error) {
	u, err := interactor.UserRepository.StoreUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return interactor.hidePassword(u), nil
}

func (interactor *UserInteractor) DeleteUser(user domain.User) error {
	err := interactor.UserRepository.DeleteUser(user.Id)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *UserInteractor) UpdateUser(user domain.User) (domain.User, error) {
	u, err := interactor.UserRepository.UpdateUser(user)
	if err != nil {
		return domain.User{}, err
	}
	return interactor.hidePassword(u), nil
}

func (interactor *UserInteractor) AuthenticateUser(user domain.User) error {
	err := interactor.UserRepository.AuthenticateUser(user)
	if err != nil {
		return err
	}
	return nil
}

func (interactor *UserInteractor) GetAllUsers() ([]domain.User, error) {
	users, err := interactor.UserRepository.GetAllUsers()
	if err != nil {
		return nil, err
	}
	for i, u := range users {
		users[i] = interactor.hidePassword(u)
	}
	return users, nil
}

func (interactor *UserInteractor) hidePassword(user domain.User) domain.User {
	user.Password = ""
	return user
}
