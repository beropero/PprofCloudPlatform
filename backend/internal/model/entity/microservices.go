// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Microservices is the golang structure for table microservices.
type Microservices struct {
	MicroserviceId          int         `json:"microserviceId"          orm:"microservice_id"          description:""` //
	ProjectId               int         `json:"projectId"               orm:"project_id"               description:""` //
	MicroserviceName        string      `json:"microserviceName"        orm:"microservice_name"        description:""` //
	Ip                      string      `json:"ip"                      orm:"ip"                       description:""` //
	Port                    int         `json:"port"                    orm:"port"                     description:""` //
	MicroserviceDescription string      `json:"microserviceDescription" orm:"microservice_description" description:""` //
	CreatedAt               *gtime.Time `json:"createdAt"               orm:"created_at"               description:""` //
	UpdatedAt               *gtime.Time `json:"updatedAt"               orm:"updated_at"               description:""` //
	DeletedAt               *gtime.Time `json:"deletedAt"               orm:"deleted_at"               description:""` //
}
