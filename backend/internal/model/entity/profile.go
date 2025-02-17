// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Profile is the golang structure for table profile.
type Profile struct {
	Id                uint64      `json:"id"                orm:"id"                  description:""` //
	MicroserviceId    int         `json:"microserviceId"    orm:"microservice_id"     description:""` //
	ProjectId         int         `json:"projectId"         orm:"project_id"          description:""` //
	DefaultSampleType string      `json:"defaultSampleType" orm:"default_sample_type" description:""` //
	DocUrl            string      `json:"docUrl"            orm:"doc_url"             description:""` //
	DropFrames        string      `json:"dropFrames"        orm:"drop_frames"         description:""` //
	KeepFrames        string      `json:"keepFrames"        orm:"keep_frames"         description:""` //
	TimeNanos         int64       `json:"timeNanos"         orm:"time_nanos"          description:""` //
	DurationNanos     int64       `json:"durationNanos"     orm:"duration_nanos"      description:""` //
	Period            int64       `json:"period"            orm:"period"              description:""` //
	PeriodTypeType    string      `json:"periodTypeType"    orm:"period_type_type"    description:""` //
	PeriodTypeUnit    string      `json:"periodTypeUnit"    orm:"period_type_unit"    description:""` //
	CreatedAt         *gtime.Time `json:"createdAt"         orm:"created_at"          description:""` //
	UpdatedAt         *gtime.Time `json:"updatedAt"         orm:"updated_at"          description:""` //
	DeletedAt         *gtime.Time `json:"deletedAt"         orm:"deleted_at"          description:""` //
}
