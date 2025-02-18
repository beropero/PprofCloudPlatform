package profile

import (
	"context"

	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"

	"backend/api/profile/v1"
)

func (c *ControllerV1) CreateProfile(ctx context.Context, req *v1.CreateProfileReq) (res *v1.CreateProfileRes, err error) {
	return nil, gerror.NewCode(gcode.CodeNotImplemented)
}
