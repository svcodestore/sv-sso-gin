package request

import (
	jwt "github.com/golang-jwt/jwt/v4"
	uuid "github.com/satori/go.uuid"
)

type CustomClaims struct {
	BaseClaims
	jwt.RegisteredClaims
}

type BaseClaims struct {
	UserId string
	UUID     uuid.UUID
	LoginId  string
	Username string
	NickName string
	ClientId string
}
