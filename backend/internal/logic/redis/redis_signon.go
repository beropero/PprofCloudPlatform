package redis

import (
	"backend/internal/consts"
	"backend/internal/model"
	"context"
	"fmt"
	"time"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
)

// // 验证码校验
// func (s *sRedis) CheckEmailCode(ctx context.Context, in model.CheckEmailCodeInput) (err error) {
// 	var (
// 		rediskey = fmt.Sprintf("email:code:%s", in.Email)
// 	)
// 	// 读取验证码
// 	code, _ := g.Redis().Do(ctx, "GET", rediskey)
// 	if code.String() != in.Code {
// 		return gerror.New("验证码错误")
// 	}
// 	// 删除验证码
// 	_, err = g.Redis().Do(ctx, "DEL", rediskey)

// 	return err
// }

// // 保存邮箱验证码
// func (s *sRedis) SaveEmailCode(ctx context.Context, in model.SaveEmailCodeInput) (err error) {
// 	var (
// 		rediskey = fmt.Sprintf("email:code:%s", in.Email)
// 	)
// 	_, err = g.Redis().Do(ctx, "SETEX", rediskey, consts.EmailCfg.Expired, in.Code)
// 	return err
// }

// 保存jwt令牌
func (s *sRedis) SaveJwtToken(ctx context.Context, in model.SaveJwtTokenInput) (err error) {
	var (
		rediskey = fmt.Sprintf("token:user:%s", in.UserId)
	)
	_, err = g.Redis().Do(ctx, "SETEX", rediskey, consts.JwtCfg.Expired, in.Token)
	return err
}

func (s *sRedis) SaveStudent(ctx context.Context, in model.SaveStudentInput) (err error) {
	var (
		rediskey = fmt.Sprintf("student:code:%s", in.Code)
	)
	_, err = g.Redis().Do(ctx, "SETEX", rediskey, 30*24*time.Hour, in.StudentSnoList)
	return err
}

// 验证码校验
func (s *sRedis) CheckEmailCode(ctx context.Context, in model.CheckEmailCodeInput) (err error) {
	var (
		rediskey = fmt.Sprintf("email:code:%s", in.Email)
	)
	// 读取验证码
	code, _ := g.Redis().Do(ctx, "GET", rediskey)
	if code.String() != in.Code {
		return gerror.New("验证码错误")
	}
	// 删除验证码
	_, err = g.Redis().Do(ctx, "DEL", rediskey)

	return err
}

// 保存邮箱验证码
func (s *sRedis) SaveEmailCode(ctx context.Context, in model.SaveEmailCodeInput) (err error) {
	var (
		rediskey = fmt.Sprintf("email:code:%s", in.Email)
	)
	_, err = g.Redis().Do(ctx, "SETEX", rediskey, consts.EmailCfg.Expired, in.Code)
	return err
}
