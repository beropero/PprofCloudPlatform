package microservice

import (
	"context"

	v1 "backend/api/microservice/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) CreateMicroservice(ctx context.Context, req *v1.CreateMicroserviceReq) (res *v1.CreateMicroserviceRes, err error) {
	err = service.Microservice().CreateMicroservice(ctx, model.CreateMicroserviceInput{
		Ip:                      req.Ip,
		MicroserviceDescription: req.MicroserviceDescription,
		MicroserviceName:        req.MicroserviceName,
		Port:                    req.Port,
		ProjectId:               req.ProjectId,
	})
	if err != nil {
		return &v1.CreateMicroserviceRes{
			Msg: "failed",
		}, err
	}
	return &v1.CreateMicroserviceRes{
		Msg: "success",
	}, err
}
