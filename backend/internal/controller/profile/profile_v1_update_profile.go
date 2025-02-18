package profile

import (
	"context"

	v1 "backend/api/profile/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) UpdateProfile(ctx context.Context, req *v1.UpdateProfileReq) (res *v1.UpdateProfileRes, err error) {
	err = service.Profile().UpdateProfile(ctx, model.UpdateProfileInput{
		Comment:   req.Comment,
		ProfileId: req.ProfileId,
	})
	if err != nil {
		return &v1.UpdateProfileRes{
			Msg: "failed",
		}, err
	}
	return &v1.UpdateProfileRes{
		Msg: "success",
	}, nil
}
