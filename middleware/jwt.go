package middleware

import (
	"github.com/gin-gonic/gin"
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model/common/response"
	"github.com/svcodestore/sv-sso-gin/model/system"
	"github.com/svcodestore/sv-sso-gin/service"
	"github.com/svcodestore/sv-sso-gin/utils"
	"go.uber.org/zap"
	"strconv"
	"strings"
	"time"
)

var jwtService = service.ServiceGroup.JwtService

func JWTCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := c.Request.Header.Get("Authorization")
		if token == "" {
			response.UnAuth(c)
			c.Abort()
			return
		}
		if strings.HasPrefix(token, "Bearer") {
			t := strings.Split(token, " ")
			if len(t) != 2 || t[1] == "" {
				response.UnAuth(c)
				c.Abort()
				return
			}
			token = t[1]
		}
		j := utils.NewJWT()
		claims, err := j.ParseToken(token)
		if err != nil {
			if err == utils.TokenExpired {
				response.UnAuthWithMessage(err.Error(), c)
				c.Abort()
				return
			}
			response.UnAuth(c)
			c.Abort()
			return
		}
		if claims.ExpiresAt.Time.Unix()-time.Now().Unix() < claims.BufferTime {
			claims.ExpiresAt.Time = time.Unix(time.Now().Unix()+global.CONFIG.JWT.ExpiresTime, 0)
			newToken, _ := j.CreateTokenByOldToken(token, *claims)
			newClaims, _ := j.ParseToken(newToken)
			c.Header("new-token", newToken)
			c.Header("new-expires-at", strconv.FormatInt(newClaims.ExpiresAt.Time.Unix(), 10))
			if global.CONFIG.System.UseMultipoint {
				err, RedisJwtToken := jwtService.GetRedisJWT(newClaims.Username)
				if err != nil {
					global.LOGGER.Error("get redis jwt failed", zap.Error(err))
				} else { // 当之前的取成功时才进行拉黑操作
					_ = jwtService.JsonInBlacklist(system.JwtBlacklist{Jwt: RedisJwtToken})
				}
				// 无论如何都要记录当前的活跃状态
				_ = jwtService.SetRedisJWT(newToken, newClaims.Username)
			}
		}
		c.Set("claims", claims)
		c.Next()
	}
}
