// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package project

import (
	"context"

	"backend/api/project/v1"
)

type IProjectV1 interface {
	DeleteProjectByIds(ctx context.Context, req *v1.DeleteProjectByIdsReq) (res *v1.DeleteProjectByIdsRes, err error)
	CreateProject(ctx context.Context, req *v1.CreateProjectReq) (res *v1.CreateProjectRes, err error)
	UpdateProject(ctx context.Context, req *v1.UpdateProjectReq) (res *v1.UpdateProjectRes, err error)
	GetProjectByPageUser(ctx context.Context, req *v1.GetProjectByPageUserReq) (res *v1.GetProjectByPageUserRes, err error)
}
