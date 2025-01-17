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
	IRedis interface {
		// 保存jwt令牌
		SaveJwtToken(ctx context.Context, in model.SaveJwtTokenInput) (err error)
		SaveStudent(ctx context.Context, in model.SaveStudentInput) (err error)
		// 验证码校验
		CheckEmailCode(ctx context.Context, in model.CheckEmailCodeInput) (err error)
		// 保存邮箱验证码
		SaveEmailCode(ctx context.Context, in model.SaveEmailCodeInput) (err error)
	}
)

var (
	localRedis IRedis
)

func Redis() IRedis {
	if localRedis == nil {
		panic("implement not found for interface IRedis, forgot register?")
	}
	return localRedis
}

func RegisterRedis(i IRedis) {
	localRedis = i
}
