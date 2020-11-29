package gateway

import (
	"github.com/kindai-csg/D-Chat/domain"
)

type GroupInput struct {
	Name         string `json:"name"`
	About        string `json:"about"`
	Participants string `json:"participants"`
}

func (group *GroupInput) GetGroup() domain.Group {
	return domain.Group{
		Name:         group.Name,
		About:        group.About,
		Participants: group.Participants,
	}
}
