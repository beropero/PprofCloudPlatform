package v1

import (
	"backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
)

type DeleteProfileByIdsReq struct {
	g.Meta `path:"/deleteprofilebyids" method:"post" tags:"Profile" summary:"删除Profile"`
	Ids    []int `json:"ids" v:"required#ID不能为空"`
}

type DeleteProfileByIdsRes struct {
	Msg string `json:"msg"`
}

type CreateProfileReq struct {
	g.Meta             `path:"/createprofile" method:"post" tags:"Profile" summary:"创建Profile"`
	ProfileId          int    `json:"profileId"  `
	ProfileName        string `json:"profileName"  `
	ProfileDescription string `json:"profileDescription"  `
}

type CreateProfileRes struct {
	Msg string `json:"msg"`
}

type UpdateProfileReq struct {
	g.Meta    `path:"/updateprofile" method:"post" tags:"Profile" summary:"更新Profile"`
	ProfileId int    `json:"projectId"   `
	Comment   string `json:"comment"    `
}

type UpdateProfileRes struct {
	Msg string `json:"msg"`
}

type GetProfileByPageUserReq struct {
	g.Meta         `path:"/getprofilebypageuser" method:"post" tags:"Profile" summary:"分页获取Profile"`
	Page           int    `json:"page"   `
	PageSize       int    `json:"pageSize"   `
	ProfileId      int    `json:"profileId"   `
	MicroserviceId int    `json:"microserviceId"` //
	ProjectId      int    `json:"projectId"     ` //
	Ptype          string `json:"ptype"      `    //
	Comment        string `json:"comment"    `
}

type GetProfileByPageUserRes struct {
	Msg      string           `json:"msg"`
	Total    int              `json:"total"   `
	Profiles []entity.Profile `json:"profiles"   `
}
