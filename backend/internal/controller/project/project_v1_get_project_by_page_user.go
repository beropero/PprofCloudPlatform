package project

import (
	"context"

	"github.com/gogf/gf/v2/util/gconv"

	v1 "backend/api/project/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) GetProjectByPageUser(ctx context.Context, req *v1.GetProjectByPageUserReq) (res *v1.GetProjectByPageUserRes, err error) {
	out, err := service.Project().GetProjectByPageUser(ctx, model.GetProjectByPageUserInput{
		Page:               req.Page,
		PageSize:           req.PageSize,
		ProjectDescription: req.ProjectDescription,
		ProjectId:          req.ProjectId,
		ProjectName:        req.ProjectName,
	})
	if err != nil {
		return &v1.GetProjectByPageUserRes{
			Msg: "failed",
		}, err
	}
	gconv.Scan(out, &res)
	res.Msg = "success"
	return res, err
}
