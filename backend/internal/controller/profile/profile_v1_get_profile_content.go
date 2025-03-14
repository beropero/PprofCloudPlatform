package profile

import (
	"context"

	v1 "backend/api/profile/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) GetProfileContent(ctx context.Context, req *v1.GetProfileContentReq) (res *v1.GetProfileContentRes, err error) {
	out, err := service.Profile().GetProfileContent(ctx, model.GetProfileContentInput{
		Id: req.Id,
	})
	if err != nil {
		return &v1.GetProfileContentRes{
			Msg: "failed",
		}, err
	}
	res = &v1.GetProfileContentRes{
		FileContent: out.FileContent,
		Msg:         "success",
	}
	return
}
