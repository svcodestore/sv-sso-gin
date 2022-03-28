package system

import (
	"encoding/base64"
	"fmt"
	"time"

	"golang.org/x/crypto/bcrypt"

	"github.com/svcodestore/sv-sso-gin/utils"
)

type CryptoService struct {
	Key       string
	Iv        string
	timeRange uint8
}

func (c *CryptoService) init() {
	rag := 2

	ts := time.Now().Unix()
	h := fmt.Sprintf("%x", ts/int64(rag))
	c.Key = utils.Pad(h, 32, h)
	iv := utils.Reverse(h)
	c.Iv = utils.Pad(iv, 16, iv)
}

func (c *CryptoService) AesEncrypt(data string) (string, error) {
	c.init()
	encrypted, err := utils.AesEncrypt([]byte(data), []byte(c.Key), []byte(c.Iv))
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(encrypted), nil
}

func (c *CryptoService) AesDecrypt(data string) (string, error) {
	c.init()
	d, _ := base64.StdEncoding.DecodeString(data)
	decrypted, err := utils.AesDecrypt(d, []byte(c.Key), []byte(c.Iv))
	if err != nil {
		return "", err
	}
	return string(decrypted), nil
}

func (c *CryptoService) PasswordHash(data string) (string, error) {
	b, err := bcrypt.GenerateFromPassword([]byte(data), bcrypt.DefaultCost)
	return string(b), err
}

func (c *CryptoService) PasswordVerify(data, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(data))
	return err == nil
}
