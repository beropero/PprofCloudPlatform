package profile

import (
	"context"

	v1 "backend/api/profile/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) DeleteProfileByIds(ctx context.Context, req *v1.DeleteProfileByIdsReq) (res *v1.DeleteProfileByIdsRes, err error) {
	if err := service.Profile().DeleteProfileByIds(ctx, model.DeleteProfileByIdsInput{
		ProfileIds: req.Ids,
	}); err != nil {
		return &v1.DeleteProfileByIdsRes{
			Msg: "failed",
		}, err
	}
	return &v1.DeleteProfileByIdsRes{
		Msg: "success",
	}, err
}
