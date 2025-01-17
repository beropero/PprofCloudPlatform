// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// TestRecords is the golang structure for table test_records.
type TestRecords struct {
	RecordId       int         `json:"recordId"       orm:"record_id"       description:""` //
	CpuUsage       float64     `json:"cpuUsage"       orm:"cpu_usage"       description:""` //
	MemoryUsage    float64     `json:"memoryUsage"    orm:"memory_usage"    description:""` //
	DiskUsage      float64     `json:"diskUsage"      orm:"disk_usage"      description:""` //
	NetworkUsage   float64     `json:"networkUsage"   orm:"network_usage"   description:""` //
	MicroserviceId int         `json:"microserviceId" orm:"microservice_id" description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""` //
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:""` //
}
