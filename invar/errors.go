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
	// ErrNotFound not fount
	ErrNotFound = errors.New("Not fount")

	// ErrInvalidNum invalid number
	ErrInvalidNum = errors.New("Invalid number")

	// ErrInvalidAccount invalid account
	ErrInvalidAccount = errors.New("Invalid account")

	// ErrInvalidToken invalid token
	ErrInvalidToken = errors.New("Invalid token")

	// ErrInvalidClient invalid client
	ErrInvalidClient = errors.New("Invalid client")

	// ErrInvalidDevice invalid device
	ErrInvalidDevice = errors.New("Invalid device")

	// ErrInvalidParams invalid params
	ErrInvalidParams = errors.New("Invalid params")

	// ErrInvalidData invalid data
	ErrInvalidData = errors.New("Invalid data")

	// ErrInvalidState invalid state
	ErrInvalidState = errors.New("Invalid state")

	// ErrTagOffline target offline
	ErrTagOffline = errors.New("Target offline")

	// ErrClientOffline client offline
	ErrClientOffline = errors.New("Client offline")

	// ErrDupRegister duplicated registration
	ErrDupRegister = errors.New("Duplicated registration")

	// ErrDupLogin duplicated admin login
	ErrDupLogin = errors.New("Duplicated admin login")

	// ErrDupData duplicate data
	ErrDupData = errors.New("Duplicated data")

	// ErrTokenExpired token expired
	ErrTokenExpired = errors.New("Token expired")

	// ErrBadPublicKey invalid public key
	ErrBadPublicKey = errors.New("Invalid public key")

	// ErrBadPrivateKey invalid private key
	ErrBadPrivateKey = errors.New("Invalid private key")

	// ErrUnkownCharType unkown chars type
	ErrUnkownCharType = errors.New("Unkown chars type")

	// ErrUnperparedState unperpared state
	ErrUnperparedState = errors.New("Unperpared state")

	// ErrOrmNotUsing orm not using
	ErrOrmNotUsing = errors.New("Orm not using")

	// ErrNoneRowFound none row found
	ErrNoneRowFound = errors.New("None row found")

	// ErrNotChanged not changed
	ErrNotChanged = errors.New("Not changed")

	// ErrNotInserted not inserted
	ErrNotInserted = errors.New("Not inserted")

	// ErrSendFailed failed to send(sms or mail)
	ErrSendFailed = errors.New("Failed to send")

	// ErrAuthDenied permission denied
	ErrAuthDenied = errors.New("Permission denied")

	// ErrKeyLenSixteen require sixteen-length secret key
	ErrKeyLenSixteen = errors.New("Require sixteen-length secret key")
)

var (
	// WErrNotFound 0x1000, not fount
	WErrNotFound = &WingErr{0x1000, ErrNotFound}

	// WErrInvalidNum 0x1001, invalid number
	WErrInvalidNum = &WingErr{0x1001, ErrInvalidNum}

	// WErrInvalidAccount 0x1002, invalid account
	WErrInvalidAccount = &WingErr{0x1002, ErrInvalidAccount}

	// WErrInvalidToken 0x1003, invalid token
	WErrInvalidToken = &WingErr{0x1003, ErrInvalidToken}

	// WErrInvalidClient 0x1004, invalid client
	WErrInvalidClient = &WingErr{0x1004, ErrInvalidClient}

	// WErrInvalidDevice 0x1005, invalid device
	WErrInvalidDevice = &WingErr{0x1005, ErrInvalidDevice}

	// WErrInvalidParams 0x1006, invalid params
	WErrInvalidParams = &WingErr{0x1006, ErrInvalidParams}

	// WErrInvalidData 0x1007, invalid data
	WErrInvalidData = &WingErr{0x1007, ErrInvalidData}

	// WErrInvalidState 0x1008, invalid state
	WErrInvalidState = &WingErr{0x1008, ErrInvalidState}

	// WErrTagOffline 0x1009, target offline
	WErrTagOffline = &WingErr{0x1009, ErrTagOffline}

	// WErrClientOffline 0x100A, client offline
	WErrClientOffline = &WingErr{0x100A, ErrClientOffline}

	// WErrDupRegister 0x100B, duplicated registration
	WErrDupRegister = &WingErr{0x100B, ErrDupRegister}

	// WErrDupLogin 0x100C, duplicated admin login
	WErrDupLogin = &WingErr{0x100C, ErrDupLogin}

	// WErrDupData 0x100D, duplicated data
	WErrDupData = &WingErr{0x100D, ErrDupData}

	// WErrTokenExpired 0x100E, token expired
	WErrTokenExpired = &WingErr{0x100E, ErrTokenExpired}

	// WErrBadPublicKey 0x100F, invalid public key
	WErrBadPublicKey = &WingErr{0x100F, ErrBadPublicKey}

	// WErrBadPrivateKey 0x1010, invalid private key
	WErrBadPrivateKey = &WingErr{0x1010, ErrBadPrivateKey}

	// WErrUnkownCharType 0x1011, unkown chars type
	WErrUnkownCharType = &WingErr{0x1011, ErrUnkownCharType}

	// WErrUnperparedState 0x1012, unperpared state
	WErrUnperparedState = &WingErr{0x1012, ErrUnperparedState}

	// WErrOrmNotUsing 0x1013, orm not using
	WErrOrmNotUsing = &WingErr{0x1013, ErrOrmNotUsing}

	// WErrNoneRowFound 0x1014, none row found
	WErrNoneRowFound = &WingErr{0x1014, ErrNoneRowFound}

	// WErrNotChanged 0x1015, not changed
	WErrNotChanged = &WingErr{0x1015, ErrNotChanged}

	// WErrNotInserted 0x1016, not inserted
	WErrNotInserted = &WingErr{0x1016, ErrNotInserted}

	// WErrSendFailed 0x1017, failed to send(sms or mail)
	WErrSendFailed = &WingErr{0x1017, ErrSendFailed}

	// WErrAuthDenied 0x1018, permission denied
	WErrAuthDenied = &WingErr{0x1018, ErrAuthDenied}

	// WErrKeyLenSixteen 0x1019, require sixteen-length secret key
	WErrKeyLenSixteen = &WingErr{0x1019, ErrKeyLenSixteen}
)
