package gateway

type LoginInput struct {
	UserId   string `json:"user_id"`
	Password string `json:"password"`
}
