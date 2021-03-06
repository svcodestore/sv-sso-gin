package system

import (
	"context"
	"errors"
	"fmt"
	"log"
	"strconv"
	"strings"
	"time"

	"github.com/go-redis/redis/v8"
	"github.com/thanhpk/randstr"

	"github.com/svcodestore/sv-sso-gin/global"
	"github.com/svcodestore/sv-sso-gin/model"
	"github.com/svcodestore/sv-sso-gin/model/system/request"
	"github.com/svcodestore/sv-sso-gin/utils"
)

const (
	GrantedCodeRedisKey        = "grantedCode"
	IssuedAccessTokenRedisKey  = "issuedAccessToken"
	IssuedRefreshTokenRedisKey = "issuedRefreshToken"
	OnlineUsers                = "onlineUsers"
)

type OauthService struct {
}

func (s *OauthService) DoGenerateGrantCode(userId, clientId string) (grantedCode string, err error) {
	grantedCode = oauthService.GenerateGrantCode()
	_, err = oauthService.SaveGrantedCodeToRedis(userId, clientId, grantedCode)
	return
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
	application, err := applicationService.ApplicationWithClientId(clientId)
	if err != nil {
		return
	}

	redirectUris := strings.Split(application.RedirectURIs, "|")
	flag := false
	for _, uris := range redirectUris {
		if redirectUri == uris {
			flag = true
			break
		}
	}

	if !flag {
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
		user, err = userService.UserWithId(userId)
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
		if err != nil {
			return
		}
		_, err = oauthService.SaveAccessTokenToRedis(user.ID, accessToken)
		if err != nil {
			return
		}
		_, err = oauthService.SaveRefreshTokenToRedis(user.ID, refreshToken)

		if err == nil {
			oauthService.DeleteGrantCodeByClientId(clientId)
			return
		}
	}
	return
}

func (s *OauthService) DoOauthLogin(username, password, loginType, clientId string) (accessToken, refreshToken string, user model.Users, err error) {
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
		if err != nil {
			return
		}
		_, err = oauthService.SaveAccessTokenToRedis(user.ID, accessToken)
		if err != nil {
			return
		}
		_, err = oauthService.SaveRefreshTokenToRedis(user.ID, refreshToken)

		return
	}
	return
}

func (s *OauthService) AllOnlineUser() (users map[string]string) {
	ctx := context.Background()
	k := OnlineUsers
	users, _ = global.REDIS.HGetAll(ctx, k).Result()

	return
}

func (s *OauthService) IsUserOnline(userId string) (isOnline bool) {
	ctx := context.Background()
	k := OnlineUsers
	isOnline, err := global.REDIS.HExists(ctx, k, userId).Result()

	if err != nil {
		isOnline = false
		return
	}

	return
}

func (s *OauthService) SetUserOnline(userId, userName string) (isOk bool) {
	ctx := context.Background()
	k := OnlineUsers
	err := global.REDIS.HSet(ctx, k, userId, userName).Err()

	if err == nil {
		isOk = true
		return
	}

	return
}

func (s *OauthService) UnsetUserOnline(userId string) (isOk bool) {
	ctx := context.Background()
	k := OnlineUsers
	err := global.REDIS.HDel(ctx, k, userId).Err()

	if err == nil {
		isOk = true
		return
	}

	return
}
