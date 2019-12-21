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
	"net/http"
)

const (
	StatusOK                  = http.StatusOK
	StatusErrParseParams      = http.StatusBadRequest
	StatusErrInputParams      = http.StatusNotAcceptable
	StatusErrUnauthorized     = http.StatusUnauthorized
	StatusErrCaseException    = http.StatusNotFound
	StatusErrPermissionDenind = http.StatusForbidden
	StatusErrFuncDisabled     = http.StatusMethodNotAllowed
	StatusErrTimeout          = http.StatusRequestTimeout
	StatusErrDuplicate        = http.StatusConflict
	StatusErrInvalidState     = http.StatusPreconditionFailed
	StatusErrLocked           = http.StatusLocked
	StatusErrGone             = http.StatusGone
)

var statusText = map[int]string{
	StatusOK:                  "OK",
	StatusErrParseParams:      "Parse Input Params Error",
	StatusErrInputParams:      "Invalid Input Params",
	StatusErrUnauthorized:     "Unauthorized",
	StatusErrCaseException:    "Case Exception",
	StatusErrPermissionDenind: "Permission Denind",
	StatusErrFuncDisabled:     "Function Disabled",
	StatusErrTimeout:          "Request Timeout",
	StatusErrDuplicate:        "Duplicate Request",
	StatusErrInvalidState:     "Invalid State",
	StatusErrLocked:           "Resource Locked",
	StatusErrGone:             "Gone",
}

// StatusText returns a text for the HTTP status code. It returns the empty
// string if the code is unknown.
func StatusText(code int) string {
	codetext := statusText[code]
	if codetext == "" {
		return http.StatusText(code)
	}
	return codetext
}
