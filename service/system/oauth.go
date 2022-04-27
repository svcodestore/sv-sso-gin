package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
	"github.com/svcodestore/sv-sso-gin/utils"
	"github.com/thanhpk/randstr"
	"log"
	"strconv"
	"strings"
	"time"
)

const (
	GrantedCodeRedisKey        = "grantedCode"
	IssuedAccessTokenRedisKey  = "issuedAccessToken"
	IssuedRefreshTokenRedisKey = "issuedRefreshToken"
)

var (
	applicationService = ApplicationService{}
	jwtService         = JwtService{}
	userService        = UserService{}
)

type OauthService struct {
}

func (s *OauthService) DoGenerateGrantCode(userId, clientId string) (string, error) {
	application, err := applicationService.ApplicationWithClientId(&model.Applications{ClientID: clientId})
	if err != nil {
		return "", err
	}
	if application.ClientID == clientId {
		grantedCode := oauthService.GenerateGrantCode()
		_, err = oauthService.SaveGrantedCodeToRedis(userId, clientId, grantedCode)
		return grantedCode, nil
	} else {
		return "", errors.New("application nonexistent")
	}
}

func (s *OauthService) GenerateGrantCode() string {
	return randstr.Hex(10)
}

func (s *OauthService) SaveGrantedCodeToRedis(userId, clientId, grantedCode string) (affected int64, err error) {
	ctx := context.Background()
	v := userId + ":" + grantedCode + fmt.Sprintf(":%d", time.Now().Add(10*time.Minute).Unix())
	affected, err = global.REDIS.HSet(ctx, GrantedCodeRedisKey, clientId, v).Result()
	if err != redis.Nil {
		err = nil
	}
	return
}

func (s *OauthService) GetGrantedCodeFromRedisByClientId(clientId string) (userId, grantedCode string, expireAt int64, err error) {
	ctx := context.Background()
	grantedCode, err = global.REDIS.HGet(ctx, GrantedCodeRedisKey, clientId).Result()
	if err != redis.Nil {
		err = nil
	}
	if grantedCode == "" {
		err = errors.New("nonexistent code")
		return
	}
	strs := strings.Split(grantedCode, ":")
	userId = strs[0]
	grantedCode = strs[1]
	expireAt, err = strconv.ParseInt(strs[2], 10, 64)
	return
}

func (s *OauthService) DeleteGrantCodeByClientId(clientId string) (isDeleted bool) {
	ctx := context.Background()
	affected, err := global.REDIS.HDel(ctx, GrantedCodeRedisKey, clientId).Result()
	if err != nil {
		log.Fatalln(err)
	}
	isDeleted = affected != 0
	return
}

func (s *OauthService) saveTokenToRedis(key, token string) (affected int64, err error) {
	ctx := context.Background()
	affected, err = global.REDIS.SAdd(ctx, key, token).Result()
	if err != redis.Nil {
		err = nil
	}
	return
}

func (s *OauthService) SaveAccessTokenToRedis(userId, token string) (affected int64, err error) {
	k := IssuedAccessTokenRedisKey + ":" + userId
	affected, err = s.saveTokenToRedis(k, token)
	return
}

func (s *OauthService) SaveRefreshTokenToRedis(userId, token string) (affected int64, err error) {
	k := IssuedRefreshTokenRedisKey + ":" + userId
	affected, err = s.saveTokenToRedis(k, token)
	return
}

func (s OauthService) deleteTokenFromRedis(key string) (affected int64, err error) {
	ctx := context.Background()
	affected, err = global.REDIS.Del(ctx, key).Result()
	if err != redis.Nil {
		err = nil
	}
	return
}

func (s *OauthService) DeleteAccessTokenFromRedis(userId string) (affected int64, err error) {
	k := IssuedAccessTokenRedisKey + ":" + userId
	affected, err = s.deleteTokenFromRedis(k)
	return
}

func (s *OauthService) DeleteRefreshTokenFromRedis(userId string) (affected int64, err error) {
	k := IssuedRefreshTokenRedisKey + ":" + userId
	affected, err = s.deleteTokenFromRedis(k)
	return
}

func (s *OauthService) IsUserLogin(userId string) (isLogin bool) {
	ctx := context.Background()
	k := IssuedAccessTokenRedisKey + ":" + userId
	count, err := global.REDIS.Exists(ctx, k).Result()
	if err == nil {
		if count == 1 {
			isLogin = true
			return
		}
	}
	return
}

func (s *OauthService) DoGenerateOauthCode(clientId, clientSecret, code, redirectUri string) (accessToken, refreshToken string, user model.Users, err error) {
	application, err := applicationService.ApplicationWithClientId(&model.Applications{ClientID: clientId})
	if err != nil {
		return
	}
	if redirectUri != application.RedirectURIs {
		err = errors.New("redirect uri " + redirectUri + " error")
		return
	}

	if clientId == application.ClientID && clientSecret == application.ClientSecret {
		userId, grantedCode, expireAt, e := oauthService.GetGrantedCodeFromRedisByClientId(clientId)
		if e != nil {
			err = e
			return
		}

		if grantedCode != code {
			err = errors.New("code err")
			return
		}

		if utils.IsExpire(expireAt) {
			oauthService.DeleteGrantCodeByClientId(clientId)
			err = errors.New("expired code")
			return
		}
		user, err = userService.UserWithId(&model.Users{ID: userId})
		if err != nil {
			return
		}
		accessToken, refreshToken, err = jwtService.GenerateToken(request.BaseClaims{
			UserId:   user.ID,
			UUID:     user.UUID,
			LoginId:  user.LoginID,
			Username: user.Name,
			ClientId: clientId,
		})

		if err == nil {
			oauthService.DeleteGrantCodeByClientId(clientId)
			return
		}
	}
	return
}

func (s *OauthService) DoOauthLogin(username, password, loginType, clientId string) (accessToken, refreshToken string, user *model.Users, err error) {
	if loginType == "login" {
		user, err = userService.Login(username, password)
		if err != nil {
			return
		}

		accessToken, refreshToken, err = jwtService.GenerateToken(request.BaseClaims{
			UUID:     user.UUID,
			UserId:   user.ID,
			Username: user.Name,
			LoginId:  user.LoginID,
			ClientId: clientId,
		})
		return
	}
	return
}
