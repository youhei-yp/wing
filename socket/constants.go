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
	ErrInvalidAccount  = SocketResp{1004, "Invalid account or password"}
	ErrInvalidCode     = SocketResp{1005, "Invalid OAuth code"}
	ErrDuplicateEmail  = SocketResp{1006, "Duplicate email to register"}
	ErrDuplicatePhone  = SocketResp{1007, "Duplicate phone to register"}
	ErrExpiredCode     = SocketResp{1008, "OAuth code expired"}
	ErrSendFailed      = SocketResp{1009, "Can not send OAuth code"}
	ErrOverloadLimit   = SocketResp{1010, "Overload the max limitations"}
	Success            = SocketResp{0001, "OK"}
)
