// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// SampleNumlabel is the golang structure for table sample_numlabel.
type SampleNumlabel struct {
	Id         uint64 `json:"id"         orm:"id"          description:""` //
	SampleId   uint64 `json:"sampleId"   orm:"sample_id"   description:""` //
	KeyStr     string `json:"keyStr"     orm:"key_str"     description:""` //
	ValueIndex uint   `json:"valueIndex" orm:"value_index" description:""` //
	NumValue   int64  `json:"numValue"   orm:"num_value"   description:""` //
	UnitStr    string `json:"unitStr"    orm:"unit_str"    description:""` //
}
