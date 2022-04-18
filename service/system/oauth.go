package system

import (
	"context"
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
	GrantedCodeRedisKey = "grantedCode"
	IssuedAccessTokenRedisKey = "issuedAccessToken"
	IssuedRefreshTokenRedisKey = "issuedRefreshToken"
)

type OauthService struct {
}

func (s *OauthService) GenerateGrantCode() string {
	return randstr.Hex(10)
}

func (s *OauthService) SaveGrantedCodeToRedis(clientId, grantedCode string) (affected int64, err error) {
	ctx := context.Background()
	affected, err = global.REDIS.HSet(ctx, GrantedCodeRedisKey, clientId, grantedCode+fmt.Sprintf(":%d", time.Now().Add(10*time.Minute).Unix())).Result()
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