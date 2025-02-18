package microservice

import (
	"context"

	v1 "backend/api/microservice/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) DeleteMicroserviceByIds(ctx context.Context, req *v1.DeleteMicroserviceByIdsReq) (res *v1.DeleteMicroserviceByIdsRes, err error) {
	err = service.Microservice().DeleteMicroserviceByIds(ctx, model.DeleteMicroserviceByIdsInput{
		MicroserviceIds: req.Ids,
	})
	if err != nil {
		return &v1.DeleteMicroserviceByIdsRes{
			Msg: "failed",
		}, err
	}
	return &v1.DeleteMicroserviceByIdsRes{
		Msg: "success",
	}, err
}
