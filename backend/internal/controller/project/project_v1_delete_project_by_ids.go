package project

import (
	"context"

	v1 "backend/api/project/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) DeleteProjectByIds(ctx context.Context, req *v1.DeleteProjectByIdsReq) (res *v1.DeleteProjectByIdsRes, err error) {
	err = service.Project().DeleteProjectByIds(ctx, model.DeleteProjectByIdsInput{
		ProjectIds: req.Ids,
	})
	if err != nil {
		return &v1.DeleteProjectByIdsRes{
			Msg: "failed",
		}, err
	}
	return &v1.DeleteProjectByIdsRes{
		Msg: "success",
	}, err
}
