package handler

type GetPasswordMessage struct {
	Host string `json:"host"`
}

type ErrorResponseMessage struct {
	Error string `json:"error"`
}

type SuccessReponseMessage struct {
	Username string `json:"username"`
	Password string `json:"password"`
}
