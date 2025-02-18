// ================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// You can delete these comments if you wish manually maintain this interface file.
// ================================================================================

package service

import (
	"backend/internal/model"
	"context"
)

type (
	IProfile interface {
		CreateProfile(ctx context.Context, in model.CreateProfileInput) (err error)
		DeleteProfileByIds(ctx context.Context, in model.DeleteProfileByIdsInput) (err error)
		UpdateProfile(ctx context.Context, in model.UpdateProfileInput) (err error)
		GetProfileByPageUser(ctx context.Context, in model.GetProfileByPageUserInput) (out model.GetProfileByPageUserOutput, err error)
	}
)

var (
	localProfile IProfile
)

func Profile() IProfile {
	if localProfile == nil {
		panic("implement not found for interface IProfile, forgot register?")
	}
	return localProfile
}

func RegisterProfile(i IProfile) {
	localProfile = i
}
