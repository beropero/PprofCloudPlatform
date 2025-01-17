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
	IJwt interface {
		GenerateJwt(ctx context.Context, in model.GenerateJwtInput) (out model.GenerateJwtOutput, err error)
		DecodeJwt(ctx context.Context, decodeJwtInput model.DecodeJwtInput) (decodeJwtOutput model.DecodeJwtOutput, err error)
	}
)

var (
	localJwt IJwt
)

func Jwt() IJwt {
	if localJwt == nil {
		panic("implement not found for interface IJwt, forgot register?")
	}
	return localJwt
}

func RegisterJwt(i IJwt) {
	localJwt = i
}
