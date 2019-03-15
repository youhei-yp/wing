// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package socket

var (
	ErrNotFound        = SocketResp{1001, "Not found"}
	ErrUnexpectedError = SocketResp{1002, "Unexpected error"}
	ErrInvalidRequest  = SocketResp{1003, "Invalid request data"}
	Success            = SocketResp{0001, "OK"}
)
