// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/internal/model"
	"context"
)

type (
	ILogin interface {
		// 用户登录
		UserLogin(ctx context.Context, in model.UserLoginInput) (out model.UserLoginOutput, err error)
		// 用户注册
		UserRegister(ctx context.Context, in model.UserRegisterInput) (out model.UserRegisterOutput, err error)
		EmailSendCode(ctx context.Context, in model.EmailSendCodeInput) (err error)
	}
)

var (
	localLogin ILogin
)

func Login() ILogin {
	if localLogin == nil {
		panic("implement not found for interface ILogin, forgot register?")
	}
	return localLogin
}

func RegisterLogin(i ILogin) {
	localLogin = i
}
