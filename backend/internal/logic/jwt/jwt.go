package jwt

import (
	"backend/internal/consts"
	"backend/internal/model"
	"backend/internal/service"
	"context"

	"time"

	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/guid"
	"github.com/golang-jwt/jwt/v5"
)

type sJwt struct{}

func init() {
	service.RegisterJwt(&sJwt{})
}

func (s *sJwt) GenerateJwt(ctx context.Context, in model.GenerateJwtInput) (out model.GenerateJwtOutput, err error) {
	var (
		secret = []byte(consts.JwtCfg.Secret)
		expire = consts.JwtCfg.Expired
	)

	data := jwt.MapClaims{
		"userid":  in.UserId,
		"userole": in.UserPermission,
		"exp":     jwt.NewNumericDate(time.Now().Add(time.Duration(expire) * time.Second)),
		"jti":     guid.S(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)
	tokenString, err := token.SignedString(secret)

	if err != nil {
		return model.GenerateJwtOutput{}, err
	}
	// redis 存储
	err = service.Redis().SaveJwtToken(ctx, model.SaveJwtTokenInput{
		UserId: gconv.String(in.UserId),
		Token:  tokenString,
	})
	if err != nil {
		return model.GenerateJwtOutput{}, err
	}
	out = model.GenerateJwtOutput{
		Token: tokenString,
	}

	return
}

func (s *sJwt) DecodeJwt(ctx context.Context, decodeJwtInput model.DecodeJwtInput) (decodeJwtOutput model.DecodeJwtOutput, err error) {
	var (
		secret      = []byte(consts.JwtCfg.Secret)
		tokenString = decodeJwtInput.Token
	)

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})

	if err != nil || !token.Valid {
		return model.DecodeJwtOutput{}, err
	}

	claims := token.Claims.(jwt.MapClaims)
	userid := gconv.String(claims["userid"])
	userole := gconv.String(claims["userole"])

	decodeJwtOutput = model.DecodeJwtOutput{
		Userid:   userid,
		UserRole: userole,
	}
	return
}
