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

// current logger output level
var currentLoggerLevel = Debug

func init() {
	beego.SetLevel(currentLoggerLevel)
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
	currentLoggerLevel = level
	beego.SetLevel(level)
}

// IsLevelEnabled check the given level if enabled
func IsLevelEnabled(level int) bool {
	return level >= currentLoggerLevel
}

// EM logs a message at emergency level.
func EM(v ...interface{}) {
	beego.Emergency(v...)
}

// AL logs a message at alert level.
func AL(v ...interface{}) {
	beego.Alert(v...)
}

// CR logs a message at critical level.
func CR(v ...interface{}) {
	beego.Critical(v...)
}

// E logs a message at error level.
func E(v ...interface{}) {
	beego.Error(v...)
}

// W logs a message at warning level.
func W(v ...interface{}) {
	beego.Warning(v...)
}

// N logs a message at notice level.
func N(v ...interface{}) {
	beego.Notice(v...)
}

// I logs a message at info level.
func I(v ...interface{}) {
	beego.Informational(v...)
}

// D logs a message at debug level.
func D(v ...interface{}) {
	beego.Debug(v...)
}
