// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Function is the golang structure for table function.
type Function struct {
	Id         uint64 `json:"id"         orm:"id"          description:""` //
	ProfileId  uint64 `json:"profileId"  orm:"profile_id"  description:""` //
	Name       string `json:"name"       orm:"name"        description:""` //
	SystemName string `json:"systemName" orm:"system_name" description:""` //
	Filename   string `json:"filename"   orm:"filename"    description:""` //
	StartLine  int64  `json:"startLine"  orm:"start_line"  description:""` //
}
