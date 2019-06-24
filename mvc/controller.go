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
	"net/http"
)

// WingController the base bee controller to support common utils
type WingController struct {
	beego.Controller
}

// ResponJSON sends a json response to client
func (c *WingController) ResponJSON(state int, data ...interface{}) {
	if state != http.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:", state, ">", ctl+"."+act)
	c.Ctx.Output.Status = state
	if len(data) > 0 {
		c.Data["json"] = data[0]
	}
	c.ServeJSON()
}

// ResponJSONP sends a jsonp response to client
func (c *WingController) ResponJSONP(state int, data ...interface{}) {
	if state != http.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:", state, ">", ctl+"."+act)
	c.Ctx.Output.Status = state
	if len(data) > 0 {
		c.Data["jsonp"] = data[0]
	}
	c.ServeJSONP()
}

// ResponXML sends xml response to client
func (c *WingController) ResponXML(state int, data ...interface{}) {
	if state != http.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:", state, ">", ctl+"."+act)
	c.Ctx.Output.Status = state
	if len(data) > 0 {
		c.Data["xml"] = data[0]
	}
	c.ServeXML()
}

// ResponYAML sends yaml response to client
func (c *WingController) ResponYAML(state int, data ...interface{}) {
	if state != http.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:", state, ">", ctl+"."+act)
	c.Ctx.Output.Status = state
	if len(data) > 0 {
		c.Data["yaml"] = data[0]
	}
	c.ServeYAML()
}

// ResponData sends YAML, XML OR JSON, depending on the value of the Accept header
func (c *WingController) ResponData(state int, data ...map[interface{}]interface{}) {
	if state != http.StatusOK {
		c.ErrorState(state)
		return
	}

	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:", state, ">", ctl+"."+act)
	c.Ctx.Output.Status = state
	if len(data) > 0 {
		c.Data = data[0]
	}
	c.ServeFormatted()
}

// ErrorState response error state to client
func (c *WingController) ErrorState(state int) {
	ctl, act := c.GetControllerAndAction()
	logger.E("Respone error:", state, ">", ctl+"."+act)

	w := c.Ctx.ResponseWriter
	w.WriteHeader(state)
	// FIXME here maybe not set content type when response error
	// w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(""))
}

// ErrorUnmarshal response unmarshal error state to client
func (c *WingController) ErrorUnmarshal(tag, err string) {
	logger.E(tag+":", "unmarshal params, err:", err)
	c.ErrorState(http.StatusBadRequest)
}

// ErrorParams response invalid params error state to client
func (c *WingController) ErrorParams(tag string, ps interface{}) {
	logger.E(tag+":", "invalid input params:", ps)
	c.ErrorState(http.StatusBadRequest)
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
