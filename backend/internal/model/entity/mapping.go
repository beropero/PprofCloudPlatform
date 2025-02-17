// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

// Mapping is the golang structure for table mapping.
type Mapping struct {
	Id              uint64 `json:"id"              orm:"id"                description:""` //
	ProfileId       uint64 `json:"profileId"       orm:"profile_id"        description:""` //
	Start           uint64 `json:"start"           orm:"start"             description:""` //
	Limit           uint64 `json:"limit"           orm:"limit"             description:""` //
	Offset          uint64 `json:"offset"          orm:"offset"            description:""` //
	File            string `json:"file"            orm:"file"              description:""` //
	BuildId         string `json:"buildId"         orm:"build_id"          description:""` //
	HasFunctions    int    `json:"hasFunctions"    orm:"has_functions"     description:""` //
	HasFilenames    int    `json:"hasFilenames"    orm:"has_filenames"     description:""` //
	HasLineNumbers  int    `json:"hasLineNumbers"  orm:"has_line_numbers"  description:""` //
	HasInlineFrames int    `json:"hasInlineFrames" orm:"has_inline_frames" description:""` //
}
