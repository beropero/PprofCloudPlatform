package microservice

import (
	"context"

	v1 "backend/api/microservice/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) UpdateMicroservice(ctx context.Context, req *v1.UpdateMicroserviceReq) (res *v1.UpdateMicroserviceRes, err error) {
	if err := service.Microservice().UpdateMicroservice(ctx, model.UpdateMicroserviceInput{
		Ip:                      req.Ip,
		MicroserviceDescription: req.MicroserviceDescription,
		MicroserviceId:          req.MicroserviceId,
		MicroserviceName:        req.MicroserviceName,
		Port:                    req.Port,
		ProjectId:               req.ProjectId,
	}); err != nil {
		return &v1.UpdateMicroserviceRes{
			Msg: "failed",
		}, err
	}
	return &v1.UpdateMicroserviceRes{
		Msg: "success",
	}, err
}
