package system

import (
	"context"
	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
	"time"

	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model/system"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
	"github.com/svcodestore/sv-sso-gin/utils"
)

type JwtService struct{}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: JsonInBlacklist
//@description: 拉黑jwt
//@param: jwtList model.JwtBlacklist
//@return: err error

func (jwtService *JwtService) JsonInBlacklist(jwtList system.JwtBlacklist) (err error) {
	err = global.DB.Create(&jwtList).Error
	if err != nil {
		return
	}
	//global.BlackCache.SetDefault(jwtList.Jwt, struct{}{})
	return
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: IsBlacklist
//@description: 判断JWT是否在黑名单内部
//@param: jwt string
//@return: bool

//func (jwtService *JwtService) IsBlacklist(jwt string) bool {
//	_, ok := global.BlackCache.Get(jwt)
//	return ok
//	// err := global.GVA_DB.Where("jwt = ?", jwt).First(&system.JwtBlacklist{}).Error
//	// isNotFound := errors.Is(err, gorm.ErrRecordNotFound)
//	// return !isNotFound
//}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: GetRedisJWT
//@description: 从redis取jwt
//@param: userName string
//@return: err error, redisJWT string

func (jwtService *JwtService) GetRedisJWT(userName string) (err error, redisJWT string) {
	redisJWT, err = global.REDIS.Get(context.Background(), userName).Result()
	return err, redisJWT
}

//@author: [piexlmax](https://github.com/piexlmax)
//@function: SetRedisJWT
//@description: jwt存入redis并设置过期时间
//@param: jwt string, userName string
//@return: err error

func (jwtService *JwtService) SetRedisJWT(jwt string, userName string) (err error) {
	// 此处过期时间等于jwt过期时间
	timer := time.Duration(global.CONFIG.JWT.ExpiresTime) * time.Second
	err = global.REDIS.Set(context.Background(), userName, jwt, timer).Err()
	return err
}

func (jwtService *JwtService) GenerateToken(c request.BaseClaims) (accessToken, refreshToken string, err error) {
	nowTime := time.Now()
	expireTime := nowTime.Add(time.Duration(168) * time.Hour)
	refreshExpireTime := nowTime.Add(time.Duration(720) * time.Hour)
	origin := global.CONFIG.JWT.Issuer

	j := &utils.JWT{
		SigningKey: []byte(global.CONFIG.JWT.SigningKey),
	}
	claims := request.CustomClaims{
		BaseClaims: c,
		BufferTime: 86400,
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
	return
}

func LoadAll() {
	var data []string
	err := global.DB.Model(&system.JwtBlacklist{}).Select("jwt").Find(&data).Error
	if err != nil {
		global.LOGGER.Error("加载数据库jwt黑名单失败!", zap.Error(err))
		return
	}
	for i := 0; i < len(data); i++ {
		//global.BlackCache.SetDefault(data[i], struct{}{})
	} // jwt黑名单 加入 BlackCache 中
}
