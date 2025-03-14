package model

import (
	"backend/internal/model/entity"
	"github.com/google/pprof/profile"
)

type CreateProfileInput struct {
	ServiceToken string `json:"serviceToken"       `
	Ptype        string `json:"ptype"       ` //
	Comment      string `json:"comment"   `   //
	OssPath      string `json:"ossPath"   `
}

type UpdateProfileInput struct {
	ProfileId int    `json:"projectId"   `
	Comment   string `json:"comment"      `
}

type DeleteProfileByIdsInput struct {
	ProfileIds []int `json:"projectIds"   `
}

type GetProfileByPageUserInput struct {
	Page           int    `json:"page"   `
	PageSize       int    `json:"pageSize"   `
	ProfileId      int    `json:"profileId"   `
	MicroserviceId int    `json:"microserviceId"` //
	ProjectId      int    `json:"projectId"     ` //
	Ptype          string `json:"ptype"      `    //
	Comment        string `json:"comment"    `
}

type GetProfileByPageUserOutput struct {
	Total    int              `json:"total"   `
	Profiles []entity.Profile `json:"profiles"   `
}

type GetProfileContentInput struct {
	Id int `json:"id"   `
}
type GetProfileContentOutput struct {
	FileContent *profile.Profile `json:"fileContent"   `
}