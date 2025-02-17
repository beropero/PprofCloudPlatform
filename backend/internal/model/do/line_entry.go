// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// LineEntry is the golang structure of table line_entry for DAO operations like Where/Data.
type LineEntry struct {
	g.Meta     `orm:"table:line_entry, do:true"`
	Id         interface{} //
	LocationId interface{} //
	FunctionId interface{} //
	Line       interface{} //
	Column     interface{} //
}
