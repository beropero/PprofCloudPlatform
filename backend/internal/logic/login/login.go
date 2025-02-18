package login

import (
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"golang.org/x/crypto/bcrypt"
)

type sLogin struct{}

func init() {
	service.RegisterLogin(&sLogin{})
}

// 用户登录
func (s *sLogin) UserLogin(ctx context.Context, in model.UserLoginInput) (out model.UserLoginOutput, err error) {
	var (
		info = entity.Users{}
	)
	err = dao.Users.Ctx(ctx).Where(g.Map{
		"email": in.Email,
	}).Scan(&info)
	if err != nil || info.UserId == 0 {
		return out, gerror.New("用户不存在")
	}
	// 比较密码哈希值与输入的密码
	err = bcrypt.CompareHashAndPassword([]byte(info.HashedPassword), []byte(in.Password))
	if err != nil {
		fmt.Println("密码错误")
		return
	}
	// 生成JWT
	jwtout, err := service.Jwt().GenerateJwt(ctx, model.GenerateJwtInput{
		UserId:         info.UserId,
		UserPermission: info.Permission,
	})
	if err != nil {
		return
	}
	out = model.UserLoginOutput{
		Permission: info.Permission,
		Token:      jwtout.Token,
	}
	return
}

// 用户注册
func (s *sLogin) UserRegister(ctx context.Context, in model.UserRegisterInput) (out model.UserRegisterOutput, err error) {
	var (
		info = entity.Users{}
	)
	// 检查验证码
	err = service.Redis().CheckEmailCode(ctx, model.CheckEmailCodeInput{
		Code:  in.Code,
		Email: in.Email,
	})
	if err != nil {
		return out, gerror.New("验证码错误")
	}
	count, err := dao.Users.Ctx(ctx).Where(g.Map{
		"email": in.Email,
	}).Count()
	if err != nil || count > 0 {
		return out, gerror.New("用户已存在")
	}
	// 生成密码哈希值
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return
	}
	_, err = dao.Users.Ctx(ctx).Insert(g.Map{
		"email":           in.Email,
		"hashed_password": string(hashedPassword),
		"permission":      "user",
	})
	if err != nil {
		return
	}
	// 生成JWT
	jwtout, err := service.Jwt().GenerateJwt(ctx, model.GenerateJwtInput{
		UserId:         info.UserId,
		UserPermission: info.Permission,
	})
	if err != nil {
		return
	}
	out.Token = jwtout.Token
	return
}
