// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package crypto

import (
	"bytes"
	"crypto/md5"
	crypto "crypto/rand"
	"encoding/base64"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/scrypt"
	"io"
	"math/rand"
	"strings"
	"time"
	"wing/logger"
	"wing/utils"
)

const (
	oauthCodeSeedsNum   = "0123456789"
	oauthCodeSeedsLower = "abcdefghijklmnopqrstuvwxyz"
	oauthCodeSeedsUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	passwordHashBytes   = 64 // password hash length
)

// Claims jwt claims data
type Claims struct {
	Keyword string `json:"keyword"`
	jwt.StandardClaims
}

// ToBase64Bytes encodes string to base64 bytes
func ToBase64Bytes(src []byte) []byte {
	return []byte(ToBase64String(src))
}

// ToBase64String encodes string to base64 string
func ToBase64String(src []byte) string {
	return base64.StdEncoding.EncodeToString(src)
}

// ToMD5 encodes string to md5 string
func ToMD5(src string) string {
	ctx := md5.New()
	ctx.Write([]byte(src))
	cipher := ctx.Sum(nil)
	return ToBase64String(cipher)
}

// GenSalt generates a random salt
func GenSalt() (string, error) {
	buf := make([]byte, passwordHashBytes)
	if _, err := io.ReadFull(crypto.Reader, buf); err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", buf), nil
}

// GenHash hash the given source with salt
func GenHash(src, salt string) (string, error) {
	hex, err := scrypt.Key([]byte(src), []byte(salt), 16384, 8, 1, passwordHashBytes)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", hex), nil
}

// VerifyJwtToken verify the encoded jwt token witch salt string
func VerifyJwtToken(signedToken, salt string) (string, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		logger.I("Verified JWT token:", signedToken, "by salt:", salt)
		return claims.Keyword, err
	}
	logger.E("Invalid JWT token:", signedToken)
	return "", err
}

// ObatinJwtToken create a jwt token with account id and salt string,
// the token will expired one hour later
func ObatinJwtToken(keyword string, salt string) (string, int64) {
	expireAt := time.Now().Add(time.Hour * 1).Unix()
	claims := Claims{
		keyword,
		jwt.StandardClaims{
			ExpiresAt: expireAt,
			Issuer:    keyword,
		},
	}

	// create the token using your claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// signs the token with a salt.
	signedToken, _ := token.SignedString([]byte(salt))

	if logger.IsEnableLevel(logger.Informational) {
		at := time.Unix(expireAt, 0).Format(utils.TimeLayout)
		logger.I("Obatin JWT token:", signedToken, "expire at:", at)
	}
	return signedToken, expireAt
}

// ObatinOAuthCode create a random OAuth code
func ObatinOAuthCode(randomLength int, randomType string) string {
	// fill random seeds chars
	buf := bytes.Buffer{}
	if strings.Contains(randomType, "0") {
		buf.WriteString(oauthCodeSeedsNum)
	}
	if strings.Contains(randomType, "a") {
		buf.WriteString(oauthCodeSeedsLower)
	}
	if strings.Contains(randomType, "A") {
		buf.WriteString(oauthCodeSeedsUpper)
	}

	// check random seeds if empty
	str := buf.String()
	len := len(str)
	if len == 0 {
		logger.E("Unknown code chars type:", randomType)
		return ""
	}

	// random OAuth code
	buf.Reset()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < randomLength; i++ {
		buf.WriteByte(str[rand.Intn(len)])
	}
	code := buf.String()
	logger.I("Obatin OAuth code:", code)
	return code
}
