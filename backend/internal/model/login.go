package model

type UserLoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type UserLoginOutput struct {
	Token      string `json:"token"`
	Permission string `json:"permission"`
}

type UserRegisterInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
	Code     string `json:"code"`
}

type UserRegisterOutput struct {
	Token string `json:"token"`
}

type EmailSendCodeInput struct {
	Email string `json:"email"`
}

type EmailSendCodeOutput struct {
	Code string `json:"code"`
}
