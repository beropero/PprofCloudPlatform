package model

type GenerateJwtInput struct {
	UserId         int
	UserPermission string
}

type GenerateJwtOutput struct {
	Token string
}

type DecodeJwtInput struct {
	Token string
}

type DecodeJwtOutput struct {
	Userid   string
	UserRole string
}
