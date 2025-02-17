// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// ProfileComment is the golang structure for table profile_comment.
type ProfileComment struct {
	Id         uint64 `json:"id"         orm:"id"          description:""` //
	ProfileId  uint64 `json:"profileId"  orm:"profile_id"  description:""` //
	Comment    string `json:"comment"    orm:"comment"     description:""` //
	OrderIndex uint   `json:"orderIndex" orm:"order_index" description:""` //
}
