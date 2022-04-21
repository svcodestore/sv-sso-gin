package system

import (
	"context"
	"errors"
	"fmt"
	"github.com/go-redis/redis/v8"
	"github.com/svcodestore/sv-sso-gin/global"
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

type OauthService struct {
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
