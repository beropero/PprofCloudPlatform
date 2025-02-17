// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SampleType is the golang structure for table sample_type.
type SampleType struct {
	Id        uint64 `json:"id"        orm:"id"         description:""` //
	ProfileId uint64 `json:"profileId" orm:"profile_id" description:""` //
	Type      string `json:"type"      orm:"type"       description:""` //
	Unit      string `json:"unit"      orm:"unit"       description:""` //
}
