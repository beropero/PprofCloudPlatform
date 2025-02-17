package upload

import (
	"context"

	v1 "backend/api/upload/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) Upload(ctx context.Context, req *v1.UploadReq) (res *v1.UploadRes, err error) {
	err = service.Upload().Upload(ctx, model.UploadInput{
		Block:     req.Block,
		Cpu:       req.Cpu,
		Goroutine: req.Goroutine,
		Memory:    req.Memory,
		Mutex:     req.Mutex,
		Types:     req.Types,
	})
	if err != nil {
		return nil, err
	}
	return &v1.UploadRes{
		Msg: "success",
	}, err
}
