package login

import (
	"backend/internal/consts"
	"backend/internal/dao"
	"backend/internal/model"
	"backend/internal/model/entity"
	"backend/internal/service"
	"context"
	"fmt"

	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"golang.org/x/crypto/bcrypt"
	"gopkg.in/gomail.v2"
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
	err = dao.Users.Ctx(ctx).Where(g.Map{
		"email": in.Email,
	}).Scan(&info)
	if err != nil || info.UserId != 0 {
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

func (s *sLogin) EmailSendCode(ctx context.Context, in model.EmailSendCodeInput) (err error) {
	// 生成6位随机验证码
	var (
		code = grand.Digits(6)
	)

	m := gomail.NewMessage()
	m.SetHeader("From", consts.EmailCfg.Username)
	m.SetHeader("To", in.Email)
	m.SetHeader("Subject", "验证码")

	msg := fmt.Sprintf("你的验证码为: %s", code)
	m.SetBody("text/html", msg)

	d := gomail.NewDialer(consts.EmailCfg.Host, consts.EmailCfg.Port, consts.EmailCfg.Username, consts.EmailCfg.Password)
	err = d.DialAndSend(m)
	if err != nil {
		return
	}
	// 保存到Redis
	err = service.Redis().SaveEmailCode(ctx, model.SaveEmailCodeInput{
		Code:  code,
		Email: in.Email,
	})
	return
}
