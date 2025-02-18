// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package microservice

import (
	"context"

	"backend/api/microservice/v1"
)

type IMicroserviceV1 interface {
	DeleteMicroserviceByIds(ctx context.Context, req *v1.DeleteMicroserviceByIdsReq) (res *v1.DeleteMicroserviceByIdsRes, err error)
	CreateMicroservice(ctx context.Context, req *v1.CreateMicroserviceReq) (res *v1.CreateMicroserviceRes, err error)
	UpdateMicroservice(ctx context.Context, req *v1.UpdateMicroserviceReq) (res *v1.UpdateMicroserviceRes, err error)
	GetMicroserviceByPageUser(ctx context.Context, req *v1.GetMicroserviceByPageUserReq) (res *v1.GetMicroserviceByPageUserRes, err error)
}
