// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// Mapping is the golang structure of table mapping for DAO operations like Where/Data.
type Mapping struct {
	g.Meta          `orm:"table:mapping, do:true"`
	Id              interface{} //
	ProfileId       interface{} //
	Start           interface{} //
	Limit           interface{} //
	Offset          interface{} //
	File            interface{} //
	BuildId         interface{} //
	HasFunctions    interface{} //
	HasFilenames    interface{} //
	HasLineNumbers  interface{} //
	HasInlineFrames interface{} //
}
