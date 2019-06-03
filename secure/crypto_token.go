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
	"fmt"
	"strconv"
	"strings"
	"time"
	"wing/invar"
)

// GenLoginToken generate a login token with account and password
// --------------------------------------------------------
//  account   password
//      |- + -|
//         |
//      base64          current nanosecode
//         |                    |
//        md5                base64
//         +------- "."---------|
//                   |
//                base64 => token
// --------------------------------------------------------
func GenLoginToken(acc, pwd string) string {
	timestamp := fmt.Sprintf("%v", time.Now().UnixNano())
	origin := EncodeB64MD5(acc+"."+pwd) + "." + EncodeBase64(timestamp)
	return EncodeBase64(origin)
}

// ViaLoginToken verify login token
// --------------------------------------------------------
//        token => base64
//                   |
//         +------- "."---------|
//        md5                base64
//         |                    |
//      base64          current nanosecode
//         |
//      |- + -|
//  account   password
// --------------------------------------------------------
func ViaLoginToken(acc, pwd, token string, duration int64) (bool, error) {
	origin, err := DecodeBase64(token)
	if err != nil {
		return false, err
	}

	segments := strings.Split(string(origin), ".")
	if segments != nil && len(segments) == 2 {
		if segments[0] != EncodeB64MD5(acc+"."+pwd) {
			return false, nil
		}

		latestByte, err := DecodeBase64(segments[1])
		if err != nil {
			return false, err
		}
		latest, err := strconv.ParseInt(string(latestByte), 10, 64)
		if err != nil {
			return false, err
		}

		// check token period
		if time.Now().UnixNano()-latest <= duration {
			return true, nil
		}
		return false, invar.ErrTokenExpired.Err
	}
	return false, nil
}
