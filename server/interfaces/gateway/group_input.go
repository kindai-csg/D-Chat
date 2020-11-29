package gateway

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type GroupInput struct {
	groupName    string `json:"groupname"`
	aboutGroup   string `json;"aboutgroup"`
	participants string `json:"participants"`
}

func (group *GroupInput) GetGroup() domain.Group {
	return domain.Group{
		groupName:    group.groupName,
		aboutGroup:   group.aboutGroup,
		participants: group.participants,
	}
}
