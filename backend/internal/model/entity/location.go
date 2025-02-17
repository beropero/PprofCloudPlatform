// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Location is the golang structure for table location.
type Location struct {
	Id        uint64 `json:"id"        orm:"id"         description:""` //
	ProfileId uint64 `json:"profileId" orm:"profile_id" description:""` //
	MappingId uint64 `json:"mappingId" orm:"mapping_id" description:""` //
	Address   uint64 `json:"address"   orm:"address"    description:""` //
	IsFolded  int    `json:"isFolded"  orm:"is_folded"  description:""` //
}
