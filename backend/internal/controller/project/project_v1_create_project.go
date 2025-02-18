package project

import (
	"context"

	v1 "backend/api/project/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) CreateProject(ctx context.Context, req *v1.CreateProjectReq) (res *v1.CreateProjectRes, err error) {
	err = service.Project().CreateProject(ctx, model.CreateProjectInput{
		ProjectDescription: req.ProjectDescription,
		ProjectName:        req.ProjectName,
	})
	if err != nil {
		return &v1.CreateProjectRes{
			Msg: "failed",
		}, err
	}
	return &v1.CreateProjectRes{
		Msg: "success",
	}, err
}
