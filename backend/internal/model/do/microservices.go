// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Microservices is the golang structure of table microservices for DAO operations like Where/Data.
type Microservices struct {
	g.Meta                  `orm:"table:microservices, do:true"`
	MicroserviceId          interface{} //
	ProjectId               interface{} //
	MicroserviceName        interface{} //
	Ip                      interface{} //
	Port                    interface{} //
	MicroserviceDescription interface{} //
	CreatedAt               *gtime.Time //
	UpdatedAt               *gtime.Time //
	DeletedAt               *gtime.Time //
}
