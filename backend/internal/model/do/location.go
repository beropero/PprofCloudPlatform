// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Location is the golang structure of table location for DAO operations like Where/Data.
type Location struct {
	g.Meta    `orm:"table:location, do:true"`
	Id        interface{} //
	ProfileId interface{} //
	MappingId interface{} //
	Address   interface{} //
	IsFolded  interface{} //
}
