package login

import (
	"context"

	v1 "backend/api/login/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) Register(ctx context.Context, req *v1.RegisterReq) (res *v1.RegisterRes, err error) {
	out, err := service.Login().UserRegister(ctx, model.UserRegisterInput{
		Code:     req.Code,
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return &v1.RegisterRes{
			Msg: "failed",
		}, err
	}
	return &v1.RegisterRes{
		Msg:   "success",
		Token: out.Token,
	}, nil
}
