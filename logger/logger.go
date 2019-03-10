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
	"github.com/astaxie/beego/logs"
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

var logger_level int = Debug

func init() {
	beego.SetLevel(logger_level)
	beego.SetLogFuncCall(true)
	logs.SetLogger("console")
	logs.SetLogger("file", `{"filename":"./logs/server.log"}`)
	logs.SetLogFuncCallDepth(5)
}

// SetLogger provides a given filename to logs a messagge to file.
func SetLogger(filename string) {
	logs.SetLogger("file", fmt.Sprintf(`{"filename":"%s"}`, filename))
}

// SetLevel sets log message level.
func SetLevel(level int) {
	logger_level = level
	beego.SetLevel(level)
}

// Check the given level if enabled
func IsEnableLevel(level int) bool {
	return level >= logger_level
}

// Logs a message at emergency level.
func EM(v ...interface{}) {
	beego.Emergency(v...)
}

// Logs a message at alert level.
func AL(v ...interface{}) {
	beego.Alert(v...)
}

// Logs a message at critical level.
func CR(v ...interface{}) {
	beego.Critical(v...)
}

// Logs a message at error level.
func E(v ...interface{}) {
	beego.Error(v...)
}

// Logs a message at warning level.
func W(v ...interface{}) {
	beego.Warning(v...)
}

// Logs a message at notice level.
func N(v ...interface{}) {
	beego.Notice(v...)
}

// Logs a message at info level.
func I(v ...interface{}) {
	beego.Informational(v...)
}

// Logs a message at debug level.
func D(v ...interface{}) {
	beego.Debug(v...)
}
