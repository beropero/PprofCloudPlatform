// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package do

import (
	"github.com/gogf/gf/v2/frame/g"
)

// ProfileComment is the golang structure of table profile_comment for DAO operations like Where/Data.
type ProfileComment struct {
	g.Meta     `orm:"table:profile_comment, do:true"`
	Id         interface{} //
	ProfileId  interface{} //
	Comment    interface{} //
	OrderIndex interface{} //
}
