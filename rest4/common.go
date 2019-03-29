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

import (
	"github.com/asaskevich/govalidator"
	"github.com/astaxie/beego"
	"wing/logger"
)

// Rest4Controller basic rest4 controller
type Rest4Controller struct {
	beego.Controller
	CtlName string // controller name for logger
}

// Rest4Resp response data to client
type Rest4Resp struct {
	State  int    `json:"state"`  // as http status, 200 is success
	Code   int    `json:"code"`   // internal handled status code
	UUID   string `json:"uuid"`   // account uuid
	Token  string `json:"token"`  // request token to call rest interface
	Latest int64  `json:"latest"` // account latest update time
	More   string `json:"more"`   // extension data
}

// Prepare get controller name
func (r *Rest4Controller) Prepare() {
	r.CtlName, _ = r.GetControllerAndAction()
}

// ParseFrom parse form data from request body
func (r *Rest4Controller) ParseFrom(form interface{}) bool {
	if err := r.ParseForm(form); err != nil {
		logger.E(r.CtlName, "Parse form, err:", err)
		r.RespErr(ErrInvalidRequest)
		return false
	}

	if ok, err := govalidator.ValidateStruct(form); err != nil || !ok {
		logger.E(r.CtlName, "Invalid form:", form, "err:", err)
		r.RespErr(ErrInvalidRequest)
		return false
	}

	logger.I(r.CtlName, "Parsed request form:", form)
	return true
}

// Obatin response handled success with uuid
func (r *Rest4Controller) Obatin(uuid, token, more string, latest int64) Rest4Resp {
	resp := Success
	resp.UUID, resp.Token, resp.More, resp.Latest = uuid, token, more, latest
	logger.I(r.CtlName, "Obatin resp:", resp)
	return resp
}

// RespData response handled data to client
func (r *Rest4Controller) RespData(resp Rest4Resp) {
	logger.I(r.CtlName, "Response result:", resp)
	r.Data["json"] = resp
	r.ServeJSON()
}

// RespErr response error data to client
func (r *Rest4Controller) RespErr(resp Rest4Resp) {
	logger.E(r.CtlName, "Response err:", resp)
	r.Data["json"] = resp
	r.ServeJSON()
}

// RespSuccess response common success
func (r *Rest4Controller) RespSuccess() {
	r.RespData(Success)
}

// RespUUID response success with uuid
func (r *Rest4Controller) RespUUID(uuid string) {
	resp := Success
	resp.UUID = uuid
	r.RespData(resp)
}

// RespToken response success with token
func (r *Rest4Controller) RespToken(token, pwd string) {
	resp := Success
	resp.Token = token
	resp.More = pwd
	r.RespData(resp)
}

// RespUnexpectedError response unexpected error
func (r *Rest4Controller) RespUnexpectedError() {
	r.RespErr(ErrUnexpectedError)
}
