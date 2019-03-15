// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package rest4

var (
	ErrNotFound        = Rest4Resp{404, 1001, "", "", 0, "Not found"}
	ErrUnexpectedError = Rest4Resp{400, 1002, "", "", 0, "Unexpected error"}
	ErrInvalidRequest  = Rest4Resp{400, 1003, "", "", 0, "Invalid request data"}
	ErrInvalidAccount  = Rest4Resp{400, 1004, "", "", 0, "Invalid account or password"}
	ErrInvalidCode     = Rest4Resp{400, 1005, "", "", 0, "Invalid OAuth code"}
	ErrDuplicateEmail  = Rest4Resp{400, 1006, "", "", 0, "Duplicate email to register"}
	ErrDuplicatePhone  = Rest4Resp{400, 1007, "", "", 0, "Duplicate phone to register"}
	ErrExpiredCode     = Rest4Resp{400, 1008, "", "", 0, "OAuth code expired"}
	ErrSendFailed      = Rest4Resp{400, 1009, "", "", 0, "Can not send OAuth code"}
	ErrOverloadLimit   = Rest4Resp{400, 1010, "", "", 0, "Overload the max limitations"}
	Success            = Rest4Resp{200, 0001, "", "", 0, "OK"}
)
