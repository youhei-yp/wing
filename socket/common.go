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
	State   int    `json:"state"`   // socket event status, 1 is success
	Message string `json:"message"` // socket event response data
}

//ValidateChecker validate parse struct datas
type ValidateChecker func(p interface{}) (bool, error)

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

// ParseJSONValidate parse json and validate input data
func ParseJSONValidate(data string, validateFunc ValidateChecker, outer interface{}) (bool, string) {
	if err := json.Unmarshal([]byte(data), outer); err != nil {
		logger.E("Unmarshal json data err:", err)
		return false, RespAck(ErrInvalidRequest)
	}

	if validateFunc != nil {
		if ok, err := validateFunc(outer); err != nil || !ok {
			logger.E("Invalid outer:", outer, "err:", err)
			return false, RespAck(ErrInvalidRequest)
		}
	}

	logger.I("Parsed request data:", outer)
	return true, ""
}

// RespAck marshal ack data to string
func RespAck(resp SocketResp) string {
	result, err := json.Marshal(resp)
	if err != nil {
		logger.E("Parse socket ack:", resp, "err:", err)
		return ""
	}
	return string(result)
}

// RespSuccess marsharl success ack to string
func RespSuccess() string {
	return RespAck(Success)
}

// RespMessage marsharl success ack witch given message to string
func RespMessage(msg string) string {
	resp := Success
	resp.Message = msg
	return RespAck(resp)
}

// RespStruct marsharl success ack witch given struct data
func RespStruct(data interface{}) string {
	msg, err := json.Marshal(data)
	if err != nil {
		logger.E("Parse socket ack struct:", data, "err:", err)
		return RespUnexpectedError()
	}
	return RespMessage(string(msg))
}

// RespNotFoundError marsharl not found error to string as socket ack
func RespNotFoundError() string {
	return RespAck(ErrNotFound)
}

// RespUnexpectedError marsharl unexpected error to string as socket ack
func RespUnexpectedError() string {
	return RespAck(ErrUnexpectedError)
}
