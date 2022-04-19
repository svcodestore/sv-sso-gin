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

func (s *OauthService) SaveAccessTokenToRedis(token string) (affected int64, err error) {
	affected, err = s.saveTokenToRedis(IssuedAccessTokenRedisKey, token)
	return
}

func (s *OauthService) SaveRefreshTokenToRedis(token string) (affected int64, err error) {
	affected, err = s.saveTokenToRedis(IssuedRefreshTokenRedisKey, token)
	return
}

func (s OauthService) deleteTokenFromRedis(key string, tokens ...string) (affected int64, err error) {
	ctx := context.Background()
	affected, err = global.REDIS.SRem(ctx, key, tokens).Result()
	if err != redis.Nil {
		err = nil
	}
	return
}

func (s *OauthService) DeleteAccessTokenFromRedis(token string) (affected int64, err error) {
	affected, err = s.deleteTokenFromRedis(IssuedAccessTokenRedisKey, token)
	return
}

func (s *OauthService) DeleteRefreshTokenFromRedis(token string) (affected int64, err error) {
	affected, err = s.deleteTokenFromRedis(IssuedRefreshTokenRedisKey, token)
	return
}

func (s *OauthService) IsUserLogin(token string) (isLogin bool) {
	ctx := context.Background()
	isLogin, _ = global.REDIS.SIsMember(ctx, IssuedAccessTokenRedisKey, token).Result()
	return
}