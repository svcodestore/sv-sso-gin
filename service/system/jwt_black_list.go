package system

import (
	"time"

	"github.com/golang-jwt/jwt/v4"

	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
	"github.com/svcodestore/sv-sso-gin/utils"
)

type JwtService struct{}

func (jwtService *JwtService) GenerateToken(c request.BaseClaims) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(168 * time.Hour)
	refreshExpireTime := nowTime.Add(720 * time.Hour)
	origin := global.CONFIG.JWT.Issuer

	j := &utils.JWT{
		SigningKey: []byte(global.CONFIG.JWT.SigningKey),
	}
	claims := request.CustomClaims{
		BaseClaims: c,
		RegisteredClaims: jwt.RegisteredClaims{
			ID:        c.LoginId,
			Issuer:    origin,
			Subject:   c.LoginId,
			Audience:  []string{c.ClientId},
			ExpiresAt: jwt.NewNumericDate(expireTime),
			NotBefore: jwt.NewNumericDate(nowTime),
			IssuedAt:  jwt.NewNumericDate(nowTime),
		},
	}
	accessToken, err = j.CreateToken(claims)
	if err != nil {
		return
	}

	claims.ExpiresAt = jwt.NewNumericDate(refreshExpireTime)
	refreshToken, err = j.CreateToken(claims)
	if err != nil {
		return
	}

	return
}
