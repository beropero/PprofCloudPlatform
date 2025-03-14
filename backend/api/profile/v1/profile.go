package v1

import (
	"backend/internal/model/entity"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/google/pprof/profile"
)

type DeleteProfileByIdsReq struct {
	g.Meta `path:"/deleteprofilebyids" method:"delete" tags:"Profile" summary:"删除Profile"`
	Ids    []int `json:"ids" v:"required#ID不能为空"`
}

type DeleteProfileByIdsRes struct {
	Msg string `json:"msg"`
}

type CreateProfileReq struct {
	g.Meta       `path:"/createprofile" method:"post" tags:"Profile" summary:"创建Profile"`
	ServiceToken string `json:"serviceToken"       v:"required#ServiceToken不能为空"`
	Comment      string `json:"comment"    `
	Ptype        string `json:"ptype"       v:"required#Ptype不能为空"`
}

type CreateProfileRes struct {
	Msg string `json:"msg"`
}

type UpdateProfileReq struct {
	g.Meta    `path:"/updateprofile" method:"put" tags:"Profile" summary:"更新Profile"`
	ProfileId int    `json:"profileId"   `
	Comment   string `json:"comment"    `
}

type UpdateProfileRes struct {
	Msg string `json:"msg"`
}

type GetProfileByPageUserReq struct {
	g.Meta         `path:"/getprofilebypageuser" method:"post" tags:"Profile" summary:"分页获取Profile"`
	Page           int    `json:"page" d:"1"`
	PageSize       int    `json:"pageSize" d:"10"`
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

type GetProfileContentReq struct {
	g.Meta `path:"/getprofilecontent" method:"get" tags:"Profile" summary:"获取Profile内容"`
	Id     int `json:"id"   `
}
type GetProfileContentRes struct {
	Msg         string          `json:"msg"`
	FileContent *profile.Profile `json:"fileContent"   `
}
