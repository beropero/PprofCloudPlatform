package upload

import (
	"context"

	v1 "backend/api/upload/v1"
	"backend/internal/model"
	"backend/internal/service"

	"github.com/gogf/gf/v2/frame/g"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	err = service.Upload().Upload(ctx, model.UploadInput{
		File:    req.File,
		Type:    req.Type,
		Token:   g.RequestFromCtx(ctx).Header.Get("X-Profile-Token"),
		Comment: req.Comment,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UploadRes{
		Msg: "success",
	}, err
}
