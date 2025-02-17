// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package entity

import (
	"github.com/gogf/gf/v2/os/gtime"
)

// Users is the golang structure for table users.
type Users struct {
	UserId         int         `json:"userId"         orm:"user_id"         description:""` //
	Email          string      `json:"email"          orm:"email"           description:""` //
	HashedPassword string      `json:"hashedPassword" orm:"hashed_password" description:""` //
	Permission     string      `json:"permission"     orm:"permission"      description:""` //
	Status         string      `json:"status"         orm:"status"          description:""` //
	CreatedAt      *gtime.Time `json:"createdAt"      orm:"created_at"      description:""` //
	UpdatedAt      *gtime.Time `json:"updatedAt"      orm:"updated_at"      description:""` //
	DeletedAt      *gtime.Time `json:"deletedAt"      orm:"deleted_at"      description:""` //
}
