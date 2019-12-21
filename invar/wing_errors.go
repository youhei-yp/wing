// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// 00002       2019/06/30   zhaixing       Add function from godfs
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

	// ErrSetFrameNil failed clear frame meta
	ErrSetFrameNil = errors.New("Failed clear frame meta")

	// ErrOperationNotSupport operation not support
	ErrOperationNotSupport = errors.New("Operation not support")

	// ErrSendHeadBytes failed send head bytes
	ErrSendHeadBytes = errors.New("Failed send head bytes")

	// ErrSendBodyBytes failed send body bytes
	ErrSendBodyBytes = errors.New("Failed send body bytes")

	// ErrReadBytes error read bytes
	ErrReadBytes = errors.New("Error read bytes")

	// ErrFileNotFound file not found
	ErrFileNotFound = errors.New("File not found")

	// ErrInternalServer internal server error
	ErrInternalServer = errors.New("Internal server error")

	// ErrDownloadFile failed download file
	ErrDownloadFile = errors.New("Failed download file")

	// ErrCreateByte failed create bytes: system protection
	ErrCreateByte = errors.New("Failed create bytes: system protection")

	// ErrAlreadyConn already connected
	ErrAlreadyConn = errors.New("Already connected")

	// ErrOpenSourceFile failed open source file
	ErrOpenSourceFile = errors.New("Failed open source file")

	// ErrEmptyReponse received empty response
	ErrEmptyReponse = errors.New("Received empty response")

	// ErrReadConf failed load config file
	ErrReadConf = errors.New("Failed load config file")

	// ErrUnexpectedDir expect file path not directory
	ErrUnexpectedDir = errors.New("Expect file path not directory")

	// ErrWriteMD5 failed write to md5
	ErrWriteMD5 = errors.New("Failed write to md5")

	// ErrWriteOut failed write out
	ErrWriteOut = errors.New("Failed write out")

	// ErrHandleDownload failed handle download file
	ErrHandleDownload = errors.New("Failed handle download file")

	// ErrFullConnPool connection pool is full
	ErrFullConnPool = errors.New("Connection pool is full")

	// ErrPoolSize thread pool size value must be positive
	ErrPoolSize = errors.New("Thread pool size value must be positive")

	// ErrPoolFull pool is full, can not take any more
	ErrPoolFull = errors.New("Pool is full, can not take any more")

	// ErrCheckDB check database: failed retry many times
	ErrCheckDB = errors.New("Check database: failed retry many times")

	// ErrFetchDB fetch database connection time out from pool
	ErrFetchDB = errors.New("Fetch database connection time out from pool")

	// ErrReadFileBody failed read file content
	ErrReadFileBody = errors.New("Failed read file content")

	// ErrNilFrame frame is null
	ErrNilFrame = errors.New("Frame is null")

	// ErrNoStorage no storage server available
	ErrNoStorage = errors.New("No storage server available")

	// ErrUnmatchLen unmatch download file length
	ErrUnmatchLen = errors.New("Unmatch download file length")

	// ErrCopyFile failed copy file
	ErrCopyFile = errors.New("Failed copy file")
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

	// WErrSetFrameNil 0x101B failed clear frame meta
	WErrSetFrameNil = &WingErr{0x101B, ErrSetFrameNil}

	// WErrOperationNotSupport 0x101C operation not support
	WErrOperationNotSupport = &WingErr{0x101C, ErrOperationNotSupport}

	// WErrSendHeadBytes 0x101D failed send head bytes
	WErrSendHeadBytes = &WingErr{0x101D, ErrSendHeadBytes}

	// WErrSendBodyBytes 0x101E failed send body bytes
	WErrSendBodyBytes = &WingErr{0x101E, ErrSendBodyBytes}

	// WErrReadBytes 0x101F error read bytes
	WErrReadBytes = &WingErr{0x101F, ErrReadBytes}

	// WErrFileNotFound 0x1020 file not found
	WErrFileNotFound = &WingErr{0x1020, ErrFileNotFound}

	// WErrInternalServer 0x1021 internal server error
	WErrInternalServer = &WingErr{0x1021, ErrInternalServer}

	// WErrDownloadFile 0x1022 failed download file
	WErrDownloadFile = &WingErr{0x1022, ErrDownloadFile}

	// WErrCreateByte 0x1023 failed create bytes: system protection
	WErrCreateByte = &WingErr{0x1023, ErrCreateByte}

	// WErrAlrdyVon 0x1024 already connected
	WErrAlreadyConn = &WingErr{0x1024, ErrAlreadyConn}

	// WErrOpenSourceFile 0x1025 failed open source file
	WErrOpenSourceFile = &WingErr{0x1025, ErrOpenSourceFile}

	// WErrEmptyRepons 0x1026 received empty response
	WErrEmptyReponse = &WingErr{0x1026, ErrEmptyReponse}

	// WErrReadConf 0x1027 failed load config file
	WErrReadConf = &WingErr{0x1027, ErrReadConf}

	// WErrDirectoryPath 0x1028 expect file path not directory
	WErrUnexpectedDir = &WingErr{0x1028, ErrUnexpectedDir}

	// WErrWriteMD5 0x1029 failed write to md5
	WErrWriteMD5 = &WingErr{0x1029, ErrWriteMD5}

	// WErrWriteOut 0x102A failed write out
	WErrWriteOut = &WingErr{0x102A, ErrWriteOut}

	// WErrHandleDownload 0x102B failed handle download file
	WErrHandleDownload = &WingErr{0x102B, ErrHandleDownload}

	// ErrFullConnPool 0x102C connection pool is full
	WErrFullConnPool = &WingErr{0x102C, ErrFullConnPool}

	// WErrPoolSize 0x102D thread pool size value must be positive
	WErrPoolSize = &WingErr{0x102D, ErrPoolSize}

	// WErrPoolFull 0x102E pool is full, can not take any more
	WErrPoolFull = &WingErr{0x102E, ErrPoolFull}

	// WErrCheckDB 0x102F check database: failed retry many times
	WErrCheckDB = &WingErr{0x102F, ErrCheckDB}

	// WErrFetchDB 0x1030 fetch database connection time out from pool
	WErrFetchDB = &WingErr{0x1030, ErrFetchDB}

	// WErrReadFileBody 0x1031 failed read file content
	WErrReadFileBody = &WingErr{0x1031, ErrReadFileBody}

	// WErrNilFrame 0x1032 frame is null
	WErrNilFrame = &WingErr{0x1032, ErrNilFrame}

	// WErrNoStorage 0x1033 no storage server available
	WErrNoStorage = &WingErr{0x1033, ErrNoStorage}

	// WErrUnmatchLen 0x1034 unmatch download file length
	WErrUnmatchLen = &WingErr{0x1034, ErrUnmatchLen}

	// WErrCopyFile failed copy file
	WErrCopyFile = &WingErr{0x1035, ErrCopyFile}
)
