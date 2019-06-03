// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package secure

import (
	"github.com/dgrijalva/jwt-go"
	"time"
)

// Claims jwt claims data
type Claims struct {
	Keyword string `json:"keyword"`
	jwt.StandardClaims
}

// ObatinJwtToken generate a jwt token with keyword and salt string,
// the token will expired after the given duration
func GenJwtToken(keyword, salt string, dur time.Duration) (string, error) {
	expireAt := time.Now().Add(dur).Unix()
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
	signedToken, err := token.SignedString([]byte(salt))
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

// ViaLoginToken verify the encoded jwt token witch salt string
func ViaJwtToken(signedToken, salt string) (string, error) {
	token, err := jwt.ParseWithClaims(signedToken, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(salt), nil
	})
	if claims, ok := token.Claims.(*Claims); ok && token.Valid {
		return claims.Keyword, nil
	}
	return "", err
}
