package v1

import (
	"backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type DeleteProjectByIdsReq struct {
	g.Meta `path:"/deleteprojectbyids" method:"post" tags:"Project" summary:"删除Project"`
	Ids    []int `json:"ids" v:"required#ID不能为空"`
}

type DeleteProjectByIdsRes struct {
	Msg string `json:"msg"`
}

type CreateProjectReq struct {
	g.Meta             `path:"/createproject" method:"post" tags:"Project" summary:"创建Project"`
	ProjectId          int    `json:"profileId"  `
	ProjectName        string `json:"profileName"  `
	ProjectDescription string `json:"profileDescription"  `
}

type CreateProjectRes struct {
	Msg string `json:"msg"`
}

type UpdateProjectReq struct {
	g.Meta             `path:"/updateproject" method:"post" tags:"Project" summary:"更新Project"`
	ProjectId          int    `json:"profileId"  `
	ProjectName        string `json:"profileName"  `
	ProjectDescription string `json:"profileDescription"  `
}

type UpdateProjectRes struct {
	Msg string `json:"msg"`
}

type GetProjectByPageUserReq struct {
	g.Meta             `path:"/getprojectbypageuser" method:"post" tags:"Project" summary:"分页获取Project"`
	Page               int    `json:"page" d:"1"`
	PageSize           int    `json:"pageSize" d:"10"`
	ProjectId          int    `json:"projectId"   `
	ProjectName        string `json:"projectName"   `
	ProjectDescription string `json:"projectDescription" `
}

type GetProjectByPageUserRes struct {
	Msg      string           `json:"msg"`
	Total    int              `json:"total"   `
	Projects []entity.Project `json:"projects"   `
}
