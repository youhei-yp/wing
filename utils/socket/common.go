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

import (
	"encoding/json"
	"github.com/asaskevich/govalidator"
	"wing/logger"
)

// SocketResp response data to client
type SocketResp struct {
	State   int    `json:"state"`
	Message string `json:"message"`
}

// ParseJSON parse json and validate input data
func ParseJSON(data string, outer interface{}) (bool, string) {
	if err := json.Unmarshal([]byte(data), outer); err != nil {
		logger.E("Unmarshal json data err:", err)
		return false, RespAck(ErrInvalidRequest)
	}

	if ok, err := govalidator.ValidateStruct(outer); err != nil || !ok {
		logger.E("Invalid outer:", outer, "err:", err)
		return false, RespAck(ErrInvalidRequest)
	}

	logger.I("Parsed request data:", outer)
	return true, ""
}

// RespAck marshal ack data to string
func RespAck(resp SocketResp) string {
	result, _ := json.Marshal(resp)
	return string(result)
}

// RespSuccess marsharl success ack to string
func RespSuccess() string {
	result, _ := json.Marshal(Success)
	return string(result)
}

// RespUnexpectedError marsharl unexpected error to string as socket ack
func RespUnexpectedError() string {
	result, _ := json.Marshal(Success)
	return string(result)
}
