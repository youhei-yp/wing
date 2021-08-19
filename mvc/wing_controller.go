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
	"strings"
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

// ensureValidatorGenerated generat the validator instance if need
func ensureValidatorGenerated() {
	if Validator == nil {
		logger.D("Generat the singleton validator instance")
		Validator = validator.New()
	}
}

// RegisterValidators register struct field validators from given map
func RegisterValidators(valmap map[string]validator.Func) {
	for tag, valfunc := range valmap {
		RegisterFieldValidator(tag, valfunc)
	}
}

// RegisterFieldValidator register validators on struct field level
func RegisterFieldValidator(tag string, valfunc validator.Func) {
	ensureValidatorGenerated()
	if err := Validator.RegisterValidation(tag, valfunc); err != nil {
		logger.E("Register struct field validator:"+tag+", err:", err)
	} else {
		logger.D("struct field validator:", tag)
	}
}

// responCheckState check respon state and print out log, the datatype must
// range in ['json', 'jsonp', 'xml', 'yaml'], if outof range current controller
// just return blank string to close http connection.
func (c *WingController) responCheckState(datatype string, needCheck bool, state int, data ...interface{}) {
	if state != invar.StatusOK {
		if needCheck {
			c.ErrorState(state)
			return
		}

		errmsg := invar.StatusText(state)
		ctl, act := c.GetControllerAndAction()
		logger.E("Respone "+strings.ToUpper(datatype)+" error:", state, ">", ctl+"."+act, errmsg)
	} else {
		ctl, act := c.GetControllerAndAction()
		logger.I("Respone state:OK-"+strings.ToUpper(datatype)+" >", ctl+"."+act)
	}

	c.Ctx.Output.Status = state
	if data != nil && len(data) > 0 {
		c.Data[datatype] = data[0]
	}

	switch datatype {
	case "json":
		c.ServeJSON()
	case "jsonp":
		c.ServeJSONP()
	case "xml":
		c.ServeXML()
	case "yaml":
		c.ServeYAML()
	default:
		// just return blank string to close http connection
		logger.W("Invalid response data tyep:" + datatype)
		c.Ctx.ResponseWriter.Write([]byte(""))
	}
}

// ResponJSON sends a json response to client,
// it may not send data if the state is not status ok
func (c *WingController) ResponJSON(state int, data ...interface{}) {
	c.responCheckState("json", true, state, data...)
}

// ResponJSONUncheck sends a json response to client witchout uncheck state code.
func (c *WingController) ResponJSONUncheck(state int, dataORerr ...interface{}) {
	c.responCheckState("json", false, state, dataORerr...)
}

// ResponJSONP sends a jsonp response to client,
// it may not send data if the state is not status ok
func (c *WingController) ResponJSONP(state int, data ...interface{}) {
	c.responCheckState("jsonp", true, state, data...)
}

// ResponJSONPUncheck sends a jsonp response to client witchout uncheck state code.
func (c *WingController) ResponJSONPUncheck(state int, dataORerr ...interface{}) {
	c.responCheckState("jsonp", false, state, dataORerr...)
}

// ResponXML sends xml response to client,
// it may not send data if the state is not status ok
func (c *WingController) ResponXML(state int, data ...interface{}) {
	c.responCheckState("xml", true, state, data...)
}

// ResponXMLUncheck sends xml response to client witchout uncheck state code.
func (c *WingController) ResponXMLUncheck(state int, dataORerr ...interface{}) {
	c.responCheckState("xml", false, state, dataORerr...)
}

// ResponYAML sends yaml response to client,
// it may not send data if the state is not status ok
func (c *WingController) ResponYAML(state int, data ...interface{}) {
	c.responCheckState("yaml", true, state, data...)
}

// ResponYAML sends yaml response to client witchout uncheck state code.
func (c *WingController) ResponYAMLUncheck(state int, dataORerr ...interface{}) {
	c.responCheckState("yaml", false, state, dataORerr...)
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
//	@deprecated this method will delete at the next version, use ErrorValidate
//				instead or DoAfterValidated to auto print error log.
func (c *WingController) ErrorParams(ps interface{}) {
	logger.E("Invalid input params:", ps)
	c.ErrorState(invar.StatusErrParseParams)
}

// ErrorValidate response 400 invalid params error state to client, then print
// the params data and validate error
func (c *WingController) ErrorValidate(ps interface{}, err ...string) {
	logger.E("Invalid input params:", ps)
	c.ErrorState(invar.StatusErrParseParams, err...)
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
//		} /** , false /* not filter error message even code is 40x */ */)
//	}
//	[CODE]
func (c *WingController) DoAfterValidated(ps interface{}, nextFunc NextFunc, option ...interface{}) {
	if err := json.Unmarshal(c.Ctx.Input.RequestBody, ps); err != nil {
		c.ErrorUnmarshal(err.Error())
		return
	}

	ensureValidatorGenerated()
	if err := Validator.Struct(ps); err != nil {
		c.ErrorValidate(ps, err.Error())
		return
	}

	// parse uncheck option, default is false
	uncheck := (option != nil && len(option) > 0 && !option[0].(bool))
	logger.D("Using uncheck:", uncheck, "mode to filter error response message")

	// execute business function after validated
	status, resp := nextFunc()
	if resp != nil {
		c.responCheckState("json", !uncheck, status, resp)
	} else {
		c.responCheckState("json", !uncheck, status)
	}
}

// DoAfterValidatedXml do bussiness action after success validate the given xml data
// returned by GenInStruct function, see DoAfterValidated() to get more informations.
func (c *WingController) DoAfterValidatedXml(ps interface{}, nextFunc NextFunc, option ...interface{}) {
	if err := xml.Unmarshal(c.Ctx.Input.RequestBody, ps); err != nil {
		c.ErrorUnmarshal(err.Error())
		return
	}

	ensureValidatorGenerated()
	if err := Validator.Struct(ps); err != nil {
		c.ErrorValidate(ps, err.Error())
		return
	}

	// parse uncheck option, default is false
	uncheck := (option != nil && len(option) > 0 && !option[0].(bool))
	logger.D("Using uncheck:", uncheck, "mode to filter error response message")

	// execute business function after validated
	status, resp := nextFunc()
	if resp != nil {
		c.responCheckState("xml", !uncheck, status, resp)
	} else {
		c.responCheckState("xml", !uncheck, status)
	}
}
