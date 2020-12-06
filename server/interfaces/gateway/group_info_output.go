package gateway

import "github.com/kindai-csg/D-Chat/domain"

type UserInfoOutput struct {
	Name         string `json:"name"`
	About        string `json:"about"`
	Participants string `json:"participants"`
}

func CreateGroupInfoOutputFromGroup(group domain.Group) GroupInfoOutput {
	return GroupInfoOutput{
		group.Name,
		group.About,
		group.Participants,
	}
}
