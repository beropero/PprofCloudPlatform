// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/internal/model"
	"context"
)

type (
	IProject interface {
		CreateProject(ctx context.Context, in model.CreateProjectInput) (id int64, err error)
		DeleteProjectByIds(ctx context.Context, in model.DeleteProjectByIdsInput) (err error)
		UpdateProject(ctx context.Context, in model.UpdateProjectInput) (err error)
		GetProjectByPageUser(ctx context.Context, in model.GetProjectByPageUserInput) (out model.GetProjectByPageUserOutput, err error)
	}
)

var (
	localProject IProject
)

func Project() IProject {
	if localProject == nil {
		panic("implement not found for interface IProject, forgot register?")
	}
	return localProject
}

func RegisterProject(i IProject) {
	localProject = i
}
