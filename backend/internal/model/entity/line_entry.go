// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// LineEntry is the golang structure for table line_entry.
type LineEntry struct {
	Id         uint64 `json:"id"         orm:"id"          description:""` //
	LocationId uint64 `json:"locationId" orm:"location_id" description:""` //
	FunctionId uint64 `json:"functionId" orm:"function_id" description:""` //
	Line       int64  `json:"line"       orm:"line"        description:""` //
	Column     int64  `json:"column"     orm:"column"      description:""` //
}
