package database

import (
	"github.com/kindai-csg/D-chat/domain"
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
		{"groupname", group.groupName},
		{"aboutgroup", group.aboutGroup},
		{"participants", group.participants},
	}

	_, err := repository.mongoHandler.Insert(repository.collectionName, doc)
	return group, err
}
