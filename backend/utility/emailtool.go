package utility

import (
	"backend/internal/model"
	"backend/internal/service"
	"backend/resource/template"
	"context"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/grand"
	"gopkg.in/gomail.v2"
)

// 发送邮箱验证码
func EmailSendCode(ctx context.Context, email string) (code string, err error) {

	// 生成6位随机验证码
	code = grand.Digits(6)
	msg, _ := g.View().ParseContent(ctx, template.EmailCodeTemplate, g.Map{
		"code": code,
	})
	err = SendEmail(ctx, email, "验证码", msg)
	if err != nil {
		return
	}
	// 保存到Redis
	err = service.Redis().SaveEmailCode(ctx, model.SaveEmailCodeInput{
		Code:  code,
		Email: email,
	})
	return
}

// 发送通用邮件
func SendEmail(ctx context.Context, email string, subject string, msg string) (err error) {
	var (
		usernameVar = g.Cfg().MustGet(ctx, "email.username")
		passwordVar = g.Cfg().MustGet(ctx, "email.password")
		hostVar     = g.Cfg().MustGet(ctx, "email.host")
		portVar     = g.Cfg().MustGet(ctx, "email.port")
	)
	m := gomail.NewMessage()
	m.SetHeader("From", usernameVar.String())
	m.SetHeader("To", email)
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", msg)
	d := gomail.NewDialer(hostVar.String(), portVar.Int(), usernameVar.String(), passwordVar.String())
	err = d.DialAndSend(m)
	return err
}
