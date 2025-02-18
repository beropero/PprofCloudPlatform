// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Project is the golang structure of table project for DAO operations like Where/Data.
type Project struct {
	g.Meta             `orm:"table:project, do:true"`
	ProjectId          interface{} //
	CreatorId          interface{} //
	ProjectName        interface{} //
	ProjectDescription interface{} //
	ProjectToken       interface{} //
	CreatedAt          *gtime.Time //
	UpdatedAt          *gtime.Time //
	DeletedAt          *gtime.Time //
}
