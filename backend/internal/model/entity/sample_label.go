// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SampleLabel is the golang structure for table sample_label.
type SampleLabel struct {
	Id       uint64 `json:"id"       orm:"id"        description:""` //
	SampleId uint64 `json:"sampleId" orm:"sample_id" description:""` //
	KeyStr   string `json:"keyStr"   orm:"key_str"   description:""` //
	ValueStr string `json:"valueStr" orm:"value_str" description:""` //
}
