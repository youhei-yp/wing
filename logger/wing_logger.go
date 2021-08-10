// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// -------------------------------------------------------------------

package logger

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/logs"
)

const (
	logConfigLevel   = "logger::level"   // configs key of logger level
	logConfigMaxDays = "logger::maxdays" // configs key of logger max days

	// LevelDebug debug level of logger
	LevelDebug = "debug"

	// LevelInfo info level of logger
	LevelInfo = "info"

	// LevelWarn warn level of logger
	LevelWarn = "warn"

	// LevelError error level of logger
	LevelError = "error"
)

// init initialize app logger
func init() {
	config := getLoggerConfigs()
	beego.SetLogger(logs.AdapterFile, config)
	beego.SetLogFuncCall(true)
	logs.SetLogFuncCallDepth(5)
	logs.Async(3) // allow asynchronous

	// set application logger level
	switch beego.AppConfig.String(logConfigLevel) {
	case LevelDebug:
		beego.SetLevel(beego.LevelDebug)
	case LevelInfo:
		beego.SetLevel(beego.LevelInformational)
	case LevelWarn:
		beego.SetLevel(beego.LevelWarning)
	case LevelError:
		beego.SetLevel(beego.LevelError)
	}
}

// getLoggerConfigs get logger configs
func getLoggerConfigs() string {
	app := beego.BConfig.AppName
	if app == "" || app == "beego" {
		app = "wing"
	}

	maxdays := beego.AppConfig.String(logConfigMaxDays)
	if maxdays == "" {
		maxdays = "7"
	}
	return "{\"filename\":\"logs/" + app + ".log\", \"daily\":true, \"maxdays\":" + maxdays + "}"
}

// SetOutputLogger close console logger on prod mode and only remain file logger.
func SetOutputLogger() {
	if beego.BConfig.RunMode != "dev" && GetLevel() != LevelDebug {
		beego.BeeLogger.DelLogger(logs.AdapterConsole)
	}
}

// GetLevel return current logger output level
func GetLevel() string {
	switch beego.BeeLogger.GetLevel() {
	case beego.LevelDebug:
		return LevelDebug
	case beego.LevelInformational:
		return LevelInfo
	case beego.LevelWarning:
		return LevelWarn
	case beego.LevelError:
		return LevelError
	default:
		return ""
	}
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
