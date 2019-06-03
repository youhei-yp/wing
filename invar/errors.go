// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------
package invar

import (
	"errors"
)

// WingErr const error with code
type WingErr struct {
	Code int
	Err  error
}

var (
	// ErrNotFound 0x1000, not fount
	ErrNotFound = &WingErr{0x1000, errors.New("Not fount")}

	// ErrInvalidNum 0x1001, invalid number
	ErrInvalidNum = &WingErr{0x1001, errors.New("Invalid number")}

	// ErrInvalidAccount 0x1002, invalid account
	ErrInvalidAccount = &WingErr{0x1002, errors.New("Invalid account")}

	// ErrInvalidToken 0x1003, invalid token
	ErrInvalidToken = &WingErr{0x1003, errors.New("Invalid token")}

	// ErrInvalidClient 0x1004, invalid client
	ErrInvalidClient = &WingErr{0x1004, errors.New("Invalid client")}

	// ErrInvalidData 0x1005, invalid data
	ErrInvalidData = &WingErr{0x1005, errors.New("Invalid data")}

	// ErrInvalidState 0x1006, invalid state
	ErrInvalidState = &WingErr{0x1006, errors.New("Invalid state")}

	// ErrTagOffline 0x1007, target offline
	ErrTagOffline = &WingErr{0x1007, errors.New("Target offline")}

	// ErrClientOffline 0x1008, client offline
	ErrClientOffline = &WingErr{0x1008, errors.New("Client offline")}

	// ErrDupRegister 0x1009, duplicated registration
	ErrDupRegister = &WingErr{0x1009, errors.New("Duplicated registration")}

	// ErrDupLogin 0x100A, duplicated admin login
	ErrDupLogin = &WingErr{0x100A, errors.New("Duplicated admin login")}

	// ErrTokenExpired 0x100B, token expired
	ErrTokenExpired = &WingErr{0x100B, errors.New("Token expired")}

	// ErrBadPublicKey 0x100C, invalid public key
	ErrBadPublicKey = &WingErr{0x100C, errors.New("Invalid public key")}

	// ErrBadPrivateKey 0x100D, invalid private key
	ErrBadPrivateKey = &WingErr{0x100D, errors.New("Invalid private key")}

	// ErrUnkownCharType 0x100E, unkown chars type
	ErrUnkownCharType = &WingErr{0x100E, errors.New("Unkown chars type")}

	// ErrUnperparedState 0x100F, unperpared state
	ErrUnperparedState = &WingErr{0x100F, errors.New("Unperpared state")}

	// ErrOrmNotUsing 0x10F0, orm not using
	ErrOrmNotUsing = &WingErr{0x10F0, errors.New("Orm not using")}

	// ErrNoneRowFound 0x10F1, none row found
	ErrNoneRowFound = &WingErr{0x10F1, errors.New("None row found")}

	// ErrSendFailed 0x10F2, failed to send(sms or mail)
	ErrSendFailed = &WingErr{0x10F2, errors.New("Failed to send")}

	// ErrAuthDenied 0x10F3, permission denied
	ErrAuthDenied = &WingErr{0x10F3, errors.New("Permission denied")}

	// ErrKeyLenSixteen require sixteen-length secret key
	ErrKeyLenSixteen = &WingErr{0x10F4, errors.New("Require sixteen-length secret key")}
)
