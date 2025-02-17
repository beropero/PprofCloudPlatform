// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SampleLocation is the golang structure for table sample_location.
type SampleLocation struct {
	SampleId   uint64 `json:"sampleId"   orm:"sample_id"   description:""` //
	LocationId uint64 `json:"locationId" orm:"location_id" description:""` //
	OrderIndex uint   `json:"orderIndex" orm:"order_index" description:""` //
}
