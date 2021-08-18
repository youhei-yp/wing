// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package mvc

import (
	"encoding/json"
	"encoding/xml"
	"github.com/astaxie/beego"
	"github.com/go-playground/validator/v10"
	"github.com/youhei-yp/wing/invar"
	"github.com/youhei-yp/wing/logger"
)

// WingController the base bee controller to support common utils
type WingController struct {
	beego.Controller
}

// GenInStruct generate the input param object for validate.
type GenInStruct func() interface{}

// NextFunc do action after input params validated.
type NextFunc func() (int, interface{})

var (
	// Validator use for verify the input params on struct level
	Validator *validator.Validate
)

// ensureValidatorIns generat the validator instance if need
func ensureValidatorGenerated() {
	if Validator == nil {
		Validator = validator.New()
	}
}

// RegisterValidators register struct field validators from given map
func RegisterValidators(valmap map[string]validator.Func) {
	for tag, valfunc := range valmap {
		RegisterFieldValidator(tag, valfunc)
	}
}

// RegisterValidators register validators on struct field level
func RegisterFieldValidator(tag string, valfunc validator.Func) {
	ensureValidatorGenerated()
	if err := Validator.RegisterValidation(tag, valfunc); err != nil {
		logger.E("Register struct field validator:"+tag+", err:", err)
	}
}

// printLogWithError printf error log
func (c *WingController) printLogWithError(tag, msg string, err ...string) {
	if err != nil && len(err) > 0 {
		logger.E(tag+":", msg, ", err:", err[0])
	} else {
		logger.E(tag+":", msg)
	}
}

// ResponJSON sends a json response to client,
// it may not send data if the state is not status ok
func (c *WingController) ResponJSON(state int, data ...interface{}) {
	if state != invar.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-JSON >", ctl+"."+act)
	c.Ctx.Output.Status = state
	if data != nil && len(data) > 0 {
		c.Data["json"] = data[0]
	}
	c.ServeJSON()
}

// ResponJSONP sends a jsonp response to client,
// it may not send data if the state is not status ok
func (c *WingController) ResponJSONP(state int, data ...interface{}) {
	if state != invar.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-JSONP >", ctl+"."+act)
	c.Ctx.Output.Status = state
	if data != nil && len(data) > 0 {
		c.Data["jsonp"] = data[0]
	}
	c.ServeJSONP()
}

// ResponXML sends xml response to client,
// it may not send data if the state is not status ok
func (c *WingController) ResponXML(state int, data ...interface{}) {
	if state != invar.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-XML >", ctl+"."+act)
	c.Ctx.Output.Status = state
	if data != nil && len(data) > 0 {
		c.Data["xml"] = data[0]
	}
	c.ServeXML()
}

// ResponYAML sends yaml response to client,
// it may not send data if the state is not status ok
func (c *WingController) ResponYAML(state int, data ...interface{}) {
	if state != invar.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-YAML >", ctl+"."+act)
	c.Ctx.Output.Status = state
	if data != nil && len(data) > 0 {
		c.Data["yaml"] = data[0]
	}
	c.ServeYAML()
}

// ResponData sends YAML, XML OR JSON, depending on the value of the Accept header,
// it may not send data if the state is not status ok
func (c *WingController) ResponData(state int, data ...map[interface{}]interface{}) {
	if state != invar.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-DATA >", ctl+"."+act)
	c.Ctx.Output.Status = state
	if data != nil && len(data) > 0 {
		c.Data = data[0]
	}
	c.ServeFormatted()
}

// ResponOK sends a empty success response to client
func (c *WingController) ResponOK() {
	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK >", ctl+"."+act)

	w := c.Ctx.ResponseWriter
	w.WriteHeader(invar.StatusOK)
	// FIXME here maybe not set content type when response error
	// w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(""))
}

// ErrorState response error state to client
func (c *WingController) ErrorState(state int, err ...string) {
	ctl, act := c.GetControllerAndAction()
	errmsg := invar.StatusText(state)
	if err != nil && len(err) > 0 {
		errmsg += ", " + err[0]
	}
	logger.E("Respone error:", state, ">", ctl+"."+act, errmsg)

	w := c.Ctx.ResponseWriter
	w.WriteHeader(state)
	// FIXME here maybe not set content type when response error
	// w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(""))
}

// ErrorUnmarshal response 400 unmarshal params error state to client
func (c *WingController) ErrorUnmarshal(err ...string) {
	c.ErrorState(invar.StatusErrParseParams, err...)
}

// ErrorParams response 400 invalid params error state to client
func (c *WingController) ErrorParams(ps interface{}) {
	c.ErrorState(invar.StatusErrParseParams, ps.(string))
}

// ErrorUnauthed response 401 unauthenticated error state to client
func (c *WingController) ErrorUnauthed(err ...string) {
	c.ErrorState(invar.StatusErrUnauthorized, err...)
}

// ErrorDenind response 403 permission denind error state to client
func (c *WingController) ErrorDenind(err ...string) {
	c.ErrorState(invar.StatusErrPermissionDenind, err...)
}

// ErrorException response 404 not found error state to client
func (c *WingController) ErrorException(err ...string) {
	c.ErrorState(invar.StatusErrCaseException, err...)
}

// ErrorDisabled response 405 function disabled error state to client
func (c *WingController) ErrorDisabled(err ...string) {
	c.ErrorState(invar.StatusErrFuncDisabled, err...)
}

// ErrorInput response 406 invalid inputs error state to client
func (c *WingController) ErrorInput(err ...string) {
	c.ErrorState(invar.StatusErrInputParams, err...)
}

// ErrorDuplicate response 409 duplicate error state to client
func (c *WingController) ErrorDuplicate(err ...string) {
	c.ErrorState(invar.StatusErrDuplicate, err...)
}

// ErrorGone response 410 gone error state to client
func (c *WingController) ErrorGone(err ...string) {
	c.ErrorState(invar.StatusErrGone, err...)
}

// ClientFrom return client ip from who requested
func (c *WingController) ClientFrom() string {
	return c.Ctx.Request.RemoteAddr
}

// BindValue bind value with key from url, the dest container must pointer
func (c *WingController) BindValue(key string, dest interface{}) error {
	if err := c.Ctx.Input.Bind(dest, key); err != nil {
		logger.E("Parse", key, "from url, err:", err)
		return invar.ErrInvalidData
	}
	return nil
}

// DoAfterValidated do bussiness action after success validate the json data returned
// by GenInStruct function, and you must register the field level validator for returned
// data's struct, then use it in struct describetion.
//	[CODE:]
//	types.go
//	~~~~~~~~~~~~~~~
//	type struct Accout {
//		Acc string `json:"acc" validate:"required,IsVaildUuid"`
//		PWD string `json:"pwd" validate:"required_without"`
//		Num int    `json:"num"`
//	}
//
//	// define custom validator on struct field level
//	func isVaildUuid(fl validator.FieldLevel) bool {
//		m, _ := regexp.Compile("^[0-9a-zA-Z]*$")
//		str := fl.Field().String()
//		return m.MatchString(str)
//	}
//
//	func init() {
//		mvc.RegisterFieldValidator("IsVaildUuid", isVaildUuid)
//	}
//
//	controller.go
//	~~~~~~~~~~~~~~~
//	func (c *AccController) AccLogin() {
//		ps := &types.Accout{}
//		c.DoAfterValidated(ps, func() (int, interface{}) {
//			// do same business function
//			// directe use c and ps param in this methed.
//			// ...
//			return http.StatusOK, "Done business"
//		})
//	}
//	[CODE]
func (c *WingController) DoAfterValidated(ps interface{}, nextFunc NextFunc) {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, ps); err != nil {
		c.ErrorUnmarshal(err.Error())
		return
	}

	ensureValidatorGenerated()
	if err := Validator.Struct(ps); err != nil {
		c.ErrorParams(err.Error())
		return
	}

	// execute business function after validated
	status, resp := nextFunc()
	if resp != nil {
		c.ResponJSON(status, resp)
	} else {
		c.ResponJSON(status)
	}
}

// DoAfterValidatedXml do bussiness action after success validate the given xml data
// returned by GenInStruct function, see DoAfterValidated() to get more informations.
func (c *WingController) DoAfterValidatedXml(ps interface{}, nextFunc NextFunc) {
	if err := xml.Unmarshal(c.Ctx.Input.RequestBody, ps); err != nil {
		c.ErrorUnmarshal(err.Error())
		return
	}

	ensureValidatorGenerated()
	if err := Validator.Struct(ps); err != nil {
		c.ErrorParams(err.Error())
		return
	}

	// execute business function after validated
	status, resp := nextFunc()
	if resp != nil {
		c.ResponXML(status, resp)
	} else {
		c.ResponXML(status)
	}
}
