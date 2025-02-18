package model

import (
	"backend/internal/model/entity"
)

type CreateProfileInput struct {
	Id             uint64 `json:"id"            `  //
	MicroserviceId int    `json:"microserviceId" ` //
	ProjectId      int    `json:"projectId"    `   //
	Ptype          string `json:"ptype"       `    //
	Comment        string `json:"comment"   `      //
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
