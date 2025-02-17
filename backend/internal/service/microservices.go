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
	IMicroservice interface {
		CreateMicroservice(ctx context.Context, in model.CreateMicroserviceInput) (err error)
		DeleteMicroserviceByIds(ctx context.Context, in model.DeleteMicroserviceByIdsInput) (err error)
		UpdateMicroservice(ctx context.Context, in model.UpdateMicroserviceInput) (err error)
		GetMicroserviceByPageUser(ctx context.Context, in model.GetMicroserviceByPageUserInput) (out model.GetMicroserviceByPageUserOutput, err error)
	}
)

var (
	localMicroservice IMicroservice
)

func Microservice() IMicroservice {
	if localMicroservice == nil {
		panic("implement not found for interface IMicroservice, forgot register?")
	}
	return localMicroservice
}

func RegisterMicroservice(i IMicroservice) {
	localMicroservice = i
}
