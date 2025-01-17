package login

import (
	"context"

	v1 "backend/api/login/v1"
	"backend/internal/model"
	"backend/internal/service"
)

func (c *ControllerV1) Login(ctx context.Context, req *v1.LoginReq) (res *v1.LoginRes, err error) {
	out, err := service.Login().UserLogin(ctx, model.UserLoginInput{
		Email:    req.Email,
		Password: req.Password,
	})
	if err != nil {
		return &v1.LoginRes{}, err
	}
	return &v1.LoginRes{
		Token: out.Token,
	}, nil
}
