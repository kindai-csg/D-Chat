package usecase

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type GroupInteractor struct {
	GroupRepository GroupRepository
}

func NewGroupInteractor(groupRepository GroupRepository) *GroupInteractor {
	groupInteractor := GroupInteractor{
		GroupRepository: groupRepository,
	}
	return &groupInteractor
}

func (interactor *GroupInteractor) Create(group domain.Group) (domain.Group, error) {
	g, err := interactor.GroupRepository.Create(group)
	if err != nil {
		return domain.Group{}, err
	}
	return g, nil
}
