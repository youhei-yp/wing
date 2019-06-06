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
	"bytes"
	"crypto/md5"
	crypto "crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"fmt"
	"github.com/bwmarrin/snowflake"
	"github.com/youhei-yp/wing/invar"
	"golang.org/x/crypto/scrypt"
	"io"
	"math/rand"
	"strings"
	"time"
)

const (
	oauthCodeSeedsNum   = "0123456789"
	oauthCodeSeedsLower = "abcdefghijklmnopqrstuvwxyz"
	oauthCodeSeedsUpper = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	radixCodeCharMap    = "01234567890abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	passwordHashBytes   = 64 // password hash length
)

// uuidNode : generate uuid string
var uuidNode *snowflake.Node

// init uuid generater
func init() {
	if uuidNode == nil {
		node, err := snowflake.NewNode(1)
		if err != nil {
			panic(err)
		}
		uuidNode = node
	}
}

// GenUUID generate a new uuid in int64
func GenUUID() int64 {
	return uuidNode.Generate().Int64()
}

// GenUUIDString generate a new uuid in string
func GenUUIDString() string {
	return uuidNode.Generate().String()
}

// GenCode generate a code by using current nanosecond
func GenCode() string {
	now := time.Now().UnixNano()
	radix := (int64)(len(radixCodeCharMap))

	code := []byte{}
	for v := now; v > 0; v /= radix {
		i := v % radix
		code = append(code, radixCodeCharMap[i])
	}
	return (string)(code)
}

// GenToken convert to lower string and encode by base64 -> md5
func GenToken(original string) string {
	return EncodeB64MD5(strings.ToLower(original))
}

// GenOAuthCode generate a random OAuth code
func GenOAuthCode(length int, randomType string) (string, error) {
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
		return "", invar.ErrUnkownCharType
	}

	// random OAuth code
	buf.Reset()
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < length; i++ {
		buf.WriteByte(str[rand.Intn(len)])
	}
	return buf.String(), nil
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

// HashMD5 hash string by md5
func HashMD5(original []byte) []byte {
	h := md5.New()
	h.Write(original)
	return h.Sum(nil)
}

// HashSHA256 hash string by sha256
func HashSHA256(original []byte) []byte {
	// h := sha256.New()
	// h.Write([]byte(original))
	// hashed := h.Sum(nil)
	hashed := sha256.Sum256(original)
	return hashed[:]
}

// DecodeBase64 decode from base64 string
func DecodeBase64(ciphertext string) (string, error) {
	original, err := base64.StdEncoding.DecodeString(ciphertext)
	if err != nil {
		return ciphertext, err
	}
	return string(original), nil
}

// EncodeBase64 encode string by base64
func EncodeBase64(original string) string {
	return base64.StdEncoding.EncodeToString([]byte(original))
}

// EncodeMD5 encode string by md5
func EncodeMD5(original string) string {
	return hex.EncodeToString(HashMD5([]byte(original)))
}

// EncodeB64MD5 encode string to base64, and then encode by md5
func EncodeB64MD5(original string) string {
	return EncodeMD5(EncodeBase64(original))
}

// EncodeMD5B64 encode string to md5, and then encode by base64
func EncodeMD5B64(original string) string {
	return EncodeBase64(EncodeMD5(original))
}
