// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Projects is the golang structure for table projects.
type Projects struct {
	ProjectId          int         `json:"projectId"          orm:"project_id"          description:""` //
	CreatorId          int         `json:"creatorId"          orm:"creator_id"          description:""` //
	ProjectName        string      `json:"projectName"        orm:"project_name"        description:""` //
	ProjectDescription string      `json:"projectDescription" orm:"project_description" description:""` //
	CreatedAt          *gtime.Time `json:"createdAt"          orm:"created_at"          description:""` //
	UpdatedAt          *gtime.Time `json:"updatedAt"          orm:"updated_at"          description:""` //
	DeletedAt          *gtime.Time `json:"deletedAt"          orm:"deleted_at"          description:""` //
}
