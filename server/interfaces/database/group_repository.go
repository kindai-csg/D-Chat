package database

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type GroupRepository struct {
	mongoHandler   MongoHandler
	collectionName string
}

func NewGroupRepository(mongoHandler MongoHandler) *GroupRepository {
	groupRepository := GroupRepository{
		mongoHandler:   mongoHandler,
		collectionName: "Groups",
	}
	return &groupRepository
}

func (repository *GroupRepository) Create(group domain.Group) (domain.Group, error) {
	doc := []KV{
		{"name", group.Name},
		{"about", group.About},
		{"participants", group.Participants},
	}

	_, err := repository.mongoHandler.Insert(repository.collectionName, doc)
	return group, err
}
