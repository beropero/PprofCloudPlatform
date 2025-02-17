package model

import "backend/internal/model/entity"

type CreateProjectInput struct {
	ProjectName        string `json:"projectName"   `
	ProjectDescription string `json:"projectDescription" `
}

type UpdateProjectInput struct {
	ProjectId          int    `json:"projectId"   `
	ProjectName        string `json:"projectName"   `
	ProjectDescription string `json:"projectDescription" `
}

type DeleteProjectByIdsInput struct {
	ProjectIds []int `json:"projectIds"   `
}

type GetProjectByPageUserInput struct {
	Page               int    `json:"page"   `
	PageSize           int    `json:"pageSize"   `
	ProjectId          int    `json:"projectId"   `
	ProjectName        string `json:"projectName"   `
	ProjectDescription string `json:"projectDescription" `
}

type GetProjectByPageUserOutput struct {
	Total    int              `json:"total"   `
	Projects []entity.Project `json:"projects"   `
}
