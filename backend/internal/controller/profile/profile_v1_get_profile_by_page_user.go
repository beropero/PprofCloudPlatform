package profile

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/util/gconv"

	v1 "backend/api/profile/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) GetProfileByPageUser(ctx context.Context, req *v1.GetProfileByPageUserReq) (res *v1.GetProfileByPageUserRes, err error) {
	out, err := service.Profile().GetProfileByPageUser(ctx, model.GetProfileByPageUserInput{
		Comment:        req.Comment,
		MicroserviceId: req.MicroserviceId,
		Page:           req.Page,
		PageSize:       req.PageSize,
		ProfileId:      req.ProfileId,
		ProjectId:      req.ProjectId,
		Ptype:          req.Ptype,
	})
	if err != nil {
		return &v1.GetProfileByPageUserRes{
			Msg: "failed",
		}, err
	}
	gconv.Scan(out, &res)
	res.Msg = "success"
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
