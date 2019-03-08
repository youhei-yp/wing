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
	ErrOrmNotUsing = errors.New("Orm not using")
	ErrNoRowFound  = errors.New("No row found")
	ErrSendFailed  = errors.New("Failed to send")
)
