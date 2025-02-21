// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gtime"
)

// Microservice is the golang structure of table microservice for DAO operations like Where/Data.
type Microservice struct {
	g.Meta                  `orm:"table:microservice, do:true"`
	MicroserviceId          interface{} //
	ProjectId               interface{} //
	CreatorId               interface{} //
	MicroserviceName        interface{} //
	Ip                      interface{} //
	Port                    interface{} //
	MicroserviceDescription interface{} //
	MicroserviceToken       interface{} //
	CreatedAt               *gtime.Time //
	UpdatedAt               *gtime.Time //
	DeletedAt               *gtime.Time //
}
