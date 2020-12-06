package gateway

type CreateGroupInfoOutputFromGroup struct {
	Name         string `json:"name"`
	About        string `json:"about"`
	Participants string `json:"participants"`
}
