package gateway

type TokenOutput struct {
	Token string `json:"token"`
}

func NewTokenOutput(token string) TokenOutput {
	return TokenOutput{
		Token: token,
	}
}
