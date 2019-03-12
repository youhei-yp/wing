// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package utils

import (
	"errors"
)

var (
	// ErrInvalidData invalid data
	ErrInvalidData = errors.New("Invalid data")

	// ErrUnperparedState unperpared runtime state
	ErrUnperparedState = errors.New("unperpared runtime state")

	// ErrOrmNotUsing not using error
	ErrOrmNotUsing = errors.New("Orm not using")

	// ErrNoRowFound not found row error
	ErrNoRowFound = errors.New("No row found")

	// ErrSendFailed failed send sms or mail error
	ErrSendFailed = errors.New("Failed to send")

	// ErrAuthDenied permission denied
	ErrAuthDenied = errors.New("Permission denied")
)
