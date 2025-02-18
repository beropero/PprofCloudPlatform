package microservice

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "backend/api/microservice/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) GetMicroserviceByPageUser(ctx context.Context, req *v1.GetMicroserviceByPageUserReq) (res *v1.GetMicroserviceByPageUserRes, err error) {
	out, err := service.Microservice().GetMicroserviceByPageUser(ctx, model.GetMicroserviceByPageUserInput{
		Ip:                      req.Ip,
		MicroserviceDescription: req.MicroserviceDescription,
		MicroserviceId:          req.MicroserviceId,
		MicroserviceName:        req.MicroserviceName,
		Page:                    req.Page,
		PageSize:                req.PageSize,
	})
	if err != nil {
		return &v1.GetMicroserviceByPageUserRes{
			Msg: "failed",
		}, err
	}
	gconv.Scan(out, &res)
	res.Msg = "success"
	return res, err
}
