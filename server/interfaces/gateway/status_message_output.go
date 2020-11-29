package gateway

type StatusMessageOutput struct {
	Status  bool   `json:"status"`
	Message string `json:"message"`
}

func CreateStatusMessageOutputFromMessage(user domain.User) UserInfoOutput {
	return StatusMessageOutput{
		message.status,
		message.message,
	}
}