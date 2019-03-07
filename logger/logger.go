// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package logger

import (
	"fmt"
	"github.com/astaxie/beego"
)

// Log levels to control the logging output.
const (
	Emergency = iota
	Alert
	Critical
	Error
	Warning
	Notice
	Informational
	Debug
)

func init() {
	beego.BeeLogger.SetLogger("console")
	beego.BeeLogger.SetLevel(Informational)
	beego.BeeLogger.EnableFuncCallDepth(true)
	beego.BeeLogger.SetLogFuncCallDepth(5) // default 5
}

// SetLogger provides a given filename to logs a messagge to file.
func SetLogger(filename string) {
	beego.BeeLogger.SetLogger("file", fmt.Sprintf(`{"filename":"%s"}`, filename))
}

// SetLevel Set log message level.
func SetLevel(level int) {
	beego.BeeLogger.SetLevel(level)
}

// Logs a message at emergency level.
func EM(v ...interface{}) {
	beego.Emergency(v)
}

// Logs a message at alert level.
func AL(v ...interface{}) {
	beego.Alert(v)
}

// Logs a message at critical level.
func CR(v ...interface{}) {
	beego.Critical(v)
}

// Logs a message at error level.
func E(v ...interface{}) {
	beego.Error(v)
}

// Logs a message at warning level.
func W(v ...interface{}) {
	beego.Warning(v)
}

// Logs a message at notice level.
func N(v ...interface{}) {
	beego.Notice(v)
}

// Logs a message at info level.
func I(v ...interface{}) {
	beego.Informational(v)
}

// Logs a message at debug level.
func D(v ...interface{}) {
	beego.Debug(v)
}
