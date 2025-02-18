// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package profile

import (
	"context"

	"backend/api/profile/v1"
)

type IProfileV1 interface {
	DeleteProfileByIds(ctx context.Context, req *v1.DeleteProfileByIdsReq) (res *v1.DeleteProfileByIdsRes, err error)
	CreateProfile(ctx context.Context, req *v1.CreateProfileReq) (res *v1.CreateProfileRes, err error)
	UpdateProfile(ctx context.Context, req *v1.UpdateProfileReq) (res *v1.UpdateProfileRes, err error)
	GetProfileByPageUser(ctx context.Context, req *v1.GetProfileByPageUserReq) (res *v1.GetProfileByPageUserRes, err error)
}
