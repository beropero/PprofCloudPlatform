// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// TestRecords is the golang structure of table test_records for DAO operations like Where/Data.
type TestRecords struct {
	g.Meta         `orm:"table:test_records, do:true"`
	RecordId       interface{} //
	CpuUsage       interface{} //
	MemoryUsage    interface{} //
	DiskUsage      interface{} //
	NetworkUsage   interface{} //
	MicroserviceId interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
}
