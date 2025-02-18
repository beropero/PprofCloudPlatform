package project

import (
	"context"

	v1 "backend/api/project/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) UpdateProject(ctx context.Context, req *v1.UpdateProjectReq) (res *v1.UpdateProjectRes, err error) {
	if err := service.Project().UpdateProject(ctx, model.UpdateProjectInput{
		ProjectDescription: req.ProjectDescription,
		ProjectId:          req.ProjectId,
		ProjectName:        req.ProjectName,
	}); err != nil {
		return &v1.UpdateProjectRes{
			Msg: "failed",
		}, err
	}
	return &v1.UpdateProjectRes{
		Msg: "success",
	}, err
}
