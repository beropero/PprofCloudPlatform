// =================================================================================
// Code generated and maintained by GoFrame CLI tool. DO NOT EDIT.
// =================================================================================

package upload

import (
	"context"

	"backend/api/upload/v1"
)

type IUploadV1 interface {
	UploadTotal(ctx context.Context, req *v1.UploadTotalReq) (res *v1.UploadTotalRes, err error)
}
