package profile

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	v1 "backend/api/profile/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) CreateProfile(ctx context.Context, req *v1.CreateProfileReq) (res *v1.CreateProfileRes, err error) {
	if err := service.Profile().CreateProfile(ctx, model.CreateProfileInput{
		ServiceToken: req.ServiceToken,
		Comment:      req.Comment,
		Ptype:        req.Ptype,
	}); err != nil {
		return &v1.CreateProfileRes{
			Msg: "failed",
		}, err
	}
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
