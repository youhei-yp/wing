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
	"github.com/astaxie/beego"
	"github.com/youhei-yp/wing/invar"
	"github.com/youhei-yp/wing/logger"
)

// WingController the base bee controller to support common utils
type WingController struct {
	beego.Controller
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
