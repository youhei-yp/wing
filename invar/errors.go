// Copyright (c) 2018-2019 WING All Rights Reserved.
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

	// ErrOverTimes over retry times
	ErrOverTimes = errors.New("Over retry times")

	// ErrSetFrameNil cannot set frame meta to nil
	ErrSetFrameNil = errors.New("cannot set frame meta to nil")

	// ErrOperationNotSupport operation not support
	ErrOperationNotSupport = errors.New("operation not support")

	// ErrSendHeadBytes
	ErrSendHeadBytes = errors.New("error send head bytes")

	// ErrSendBodyBytes
	ErrSendBodyBytes = errors.New("error send body bytes")

	// ErrRead error read bytes
	ErrRead = errors.New("error read bytes")

	// ErrFileNotFound file not found
	ErrFileNotFound = errors.New("file not found")

	// ErrInternalServer internal server error
	ErrInternalServer = errors.New("internal server error")

	// ErrDownloadFile error download file
	ErrDownloadFile = errors.New("error download file")

	// ErrCreateByte cannot create bytes: system protection
	ErrCreateByte = errors.New("cannot create bytes: system protection")

	// ErrAlrdyVon already connected
	ErrAlrdyVon = errors.New("already connected")

	// ErrOpenSourceFile open source file failed
	ErrOpenSourceFile = errors.New("open source file failed")

	// ErrEmptyRepons receive empty response from server
	ErrEmptyRepons = errors.New("receive empty response from server")

	// ErrReadConf error read from configuration file
	ErrReadConf = errors.New("error read from configuration file")

	// ErrDirectoryPath expect file path not directory path
	ErrDirectoryPath = errors.New("expect file path not directory path")

	// ErrWriteMd5 error write md
	ErrWriteMd5 = errors.New("error write md")

	// ErrWriteOut error write out
	ErrWriteOut = errors.New("error write out")

	// ErrHandleDownload error handle download file
	ErrHandleDownload = errors.New("error handle download file")

	// ErrFullConnectionPool connection pool is full
	ErrFullConnectionPool = errors.New("connection pool is full")

	// ErrPoolSize thread pool size value must be positive
	ErrPoolSize = errors.New("thread pool size value must be positive")

	// ErrPoolFull pool is full, can not take any more
	ErrPoolFull = errors.New("pool is full, can not take any more")

	// ErrCheckDb error check db: failed retry many times
	ErrCheckDb = errors.New("error check db: failed retry many times")

	// ErrFetchDb cannot fetch db connection from pool: wait time out
	ErrFetchDb = errors.New("cannot fetch db connection from pool: wait time out")

	// ErrReadFileBody read file body failed
	ErrReadFileBody = errors.New("read file body failed")

	// ErrNilFrame frame is null
	ErrNilFrame = errors.New("frame is null")

	// ErrNoStorage no storage server available
	ErrNoStorage = errors.New("no storage server available")

	// ErrDownloadWrongLen download return wrong file length
	ErrDownloadWrongLen = errors.New("download return wrong file length")
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

	// WErrOverTimes 0x101A over try times
	WErrOverTimes = &WingErr{0x101A, ErrOverTimes}

	// WErrSetFrameNil 0x101B cannot set frame meta to nil
	WErrSetFrameNil = &WingErr{0x101B, ErrSetFrameNil}

	// WErrOperationNotSupport 0x101C operation not support
	WErrOperationNotSupport = &WingErr{0x101C, ErrOperationNotSupport}

	// WErrSendHeadBytes 0x101D error send head bytes
	WErrSendHeadBytes = &WingErr{0x101D, ErrSendHeadBytes}

	// WErrSendBodyBytes 0x101E error send body bytes
	WErrSendBodyBytes = &WingErr{0x101E, ErrSendBodyBytes}

	// WErrRead 0x101F error read bytes
	WErrRead = &WingErr{0x101F, ErrRead}

	// WErrFileNotFound 0x1020 file not found
	WErrFileNotFound = &WingErr{0x1020, ErrFileNotFound}

	// WErrInternalServer 0x1021 internal server error
	WErrInternalServer = &WingErr{0x1021, ErrInternalServer}

	// WErrDownloadFile 0x1022 error download file
	WErrDownloadFile = &WingErr{0x1022, ErrDownloadFile}

	// WErrCreateByte 0x1023 cannot create bytes: system protection
	WErrCreateByte = &WingErr{0x1023, ErrCreateByte}

	// WErrAlrdyVon 0x1024 already connected
	WErrAlrdyVon = &WingErr{0x1024, ErrAlrdyVon}

	// WErrOpenSourceFile 0x1025 open source file failed
	WErrOpenSourceFile = &WingErr{0x1025, ErrOpenSourceFile}

	// WErrEmptyRepons 0x1026 receive empty response from server
	WErrEmptyRepons = &WingErr{0x1026, ErrEmptyRepons}

	// WErrReadConf 0x1027 error read from configuration file
	WErrReadConf = &WingErr{0x1027, ErrReadConf}

	// WErrDirectoryPath 0x1028 expect file path not directory path
	WErrDirectoryPath = &WingErr{0x1028, ErrDirectoryPath}

	// WErrWriteMd5 0x1029 error write md
	WErrWriteMd5 = &WingErr{0x1029, ErrWriteMd5}

	// WErrWriteOut 0x102A error write out
	WErrWriteOut = &WingErr{0x102A, ErrWriteOut}

	// WErrHandleDownload 0x102B error handle download file
	WErrHandleDownload = &WingErr{0x102B, ErrHandleDownload}

	// WErrFullConnectionPool 0x102C connection pool is full
	WErrFullConnectionPool = &WingErr{0x102C, ErrFullConnectionPool}

	// WErrPoolSize 0x102D thread pool size value must be positive
	WErrPoolSize = &WingErr{0x102D, ErrPoolSize}

	// WErrPoolFull 0x102E pool is full, can not take any more
	WErrPoolFull = &WingErr{0x102E, ErrPoolFull}

	// WErrCheckDb 0x102F error check db: failed retry many times
	WErrCheckDb = &WingErr{0x102F, ErrCheckDb}

	// WErrFetchDb 0x1030 cannot fetch db connection from pool: wait time out
	WErrFetchDb = &WingErr{0x1030, ErrFetchDb}

	// WErrReadFileBody 0x1031 read file body failed
	WErrReadFileBody = &WingErr{0x1031, ErrReadFileBody}

	// WErrNilFrame 0x1032 frame is null
	WErrNilFrame = &WingErr{0x1032, ErrNilFrame}

	// WErrNoStorage 0x1033 no storage server available
	WErrNoStorage = &WingErr{0x1033, ErrNoStorage}

	// WErrDownloadWrongLen 0x1034 download return wrong file length
	WErrDownloadWrongLen = &WingErr{0x1034, ErrDownloadWrongLen}
)
