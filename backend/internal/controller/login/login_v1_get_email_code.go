package login

import (
	"context"

	v1 "backend/api/login/v1"
	"backend/utility"
)

func (c *ControllerV1) GetEmailCode(ctx context.Context, req *v1.GetEmailCodeReq) (res *v1.GetEmailCodeRes, err error) {
	_, err = utility.EmailSendCode(ctx, req.Email)
	if err != nil {
		return &v1.GetEmailCodeRes{
			Msg: "failed",
		}, err
	}
	return &v1.GetEmailCodeRes{
		Msg: "success",
	}, nil
}
