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

// ResponJSON sends a json response to client
func (c *WingController) ResponJSON(data interface{}) {
	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-JSON >", ctl+"."+act)
	c.Ctx.Output.Status = invar.StatusOK
	c.Data["json"] = data
	c.ServeJSON()
}

// ResponJSONP sends a jsonp response to client
func (c *WingController) ResponJSONP(data interface{}) {
	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-JSONP >", ctl+"."+act)
	c.Ctx.Output.Status = invar.StatusOK
	c.Data["jsonp"] = data
	c.ServeJSONP()
}

// ResponXML sends xml response to client
func (c *WingController) ResponXML(data interface{}) {
	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-XML >", ctl+"."+act)
	c.Ctx.Output.Status = invar.StatusOK
	c.Data["xml"] = data
	c.ServeXML()
}

// ResponYAML sends yaml response to client
func (c *WingController) ResponYAML(data ...interface{}) {
	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-YAML >", ctl+"."+act)
	c.Ctx.Output.Status = invar.StatusOK
	c.Data["yaml"] = data
	c.ServeYAML()
}

// ResponData sends YAML, XML OR JSON, depending on the value of the Accept header
func (c *WingController) ResponData(data map[interface{}]interface{}) {
	ctl, act := c.GetControllerAndAction()
	logger.I("Respone state:OK-DATA >", ctl+"."+act)
	c.Ctx.Output.Status = invar.StatusOK
	c.Data = data
	c.ServeFormatted()
}

// ResponOK sends a json response to client
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
	if len(err) > 0 {
		errmsg += ", " + err[0]
	}
	logger.E("Respone error:", state, ">", ctl+"."+act, errmsg)

	w := c.Ctx.ResponseWriter
	w.WriteHeader(state)
	// FIXME here maybe not set content type when response error
	// w.Header().Set("Content-Type", "application/json")
	w.Write([]byte(""))
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
