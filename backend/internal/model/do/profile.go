// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Profile is the golang structure of table profile for DAO operations like Where/Data.
type Profile struct {
	g.Meta         `orm:"table:profile, do:true"`
	Id             interface{} //
	MicroserviceId interface{} //
	ProjectId      interface{} //
	Ptype          interface{} //
	OssPath        interface{} //
	CreatedAt      *gtime.Time //
	UpdatedAt      *gtime.Time //
	DeletedAt      *gtime.Time //
}
