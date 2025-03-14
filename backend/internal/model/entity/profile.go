// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Profile is the golang structure for table profile.
type Profile struct {
	Id             int         `json:"id"             orm:"id"              description:""` //
	MicroserviceId int         `json:"microserviceId" orm:"microservice_id" description:""` //
	ProjectId      int         `json:"projectId"      orm:"project_id"      description:""` //
	Ptype          string      `json:"ptype"          orm:"ptype"           description:""` //
	OssPath        string      `json:"ossPath"        orm:"oss_path"        description:""` //
	Comment        string      `json:"comment"        orm:"comment"         description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""` //
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:""` //
}
