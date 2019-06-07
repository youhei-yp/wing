// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
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
	"database/sql"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/youhei-yp/wing/invar"
	"reflect"
	"strings"
)

// WingProvider content provider to support database utils
type WingProvider struct {
	Conn *sql.DB
}

const (
	// limitPageItems limit to show lits items in one page
	limitPageItems = 50
)

var (
	// WingHelper content provider to hold database connections,
	// it will nil before mvc.OpenDatabase() called
	WingHelper *WingProvider
)

// OpenDatabase connect database and check ping result,
// the connections holded by mvc.WingHelper object
func OpenDatabase() error {
	dsn := fmt.Sprintf("%s:%s@/%s?charset=utf8",
		beego.AppConfig.String("dbuser"),
		beego.AppConfig.String("dbpwd"),
		beego.AppConfig.String("dbname"))

	// open and connect database
	con, err := sql.Open("mysql", dsn)
	if err != nil {
		return err
	}

	// check database validable
	if err = con.Ping(); err != nil {
		return err
	}

	con.SetMaxIdleConns(100)
	con.SetMaxOpenConns(100)
	WingHelper = &WingProvider{con}
	return nil
}

// Stub return content provider connection
func (w *WingProvider) Stub() *sql.DB {
	return w.Conn
}

// Query call sql.Query()
func (w *WingProvider) Query(query string, args ...interface{}) (*sql.Rows, error) {
	return w.Conn.Query(query, args...)
}

// Prepare call sql.Prepare()
func (w *WingProvider) Prepare(query string) (*sql.Stmt, error) {
	return w.Conn.Prepare(query)
}

// AppendLimit append page limitation end of sql string
func (w *WingProvider) AppendLimit(sql string, page int) string {
	return fmt.Sprintf("%s LIMIT %d, %d", sql, page*limitPageItems, limitPageItems)
}

// CheckAffected append page limitation end of sql string
func (w *WingProvider) Affected(result sql.Result) error {
	if row, err := result.RowsAffected(); err != nil || row == 0 {
		return invar.ErrNotChanged
	}
	return nil
}

// FormatSets format update sets for sql update
// [CODE:]
// sets := w.FormatSets(struct {
//     StringFiled string
//     IntFiled    int
//     I32Filed    int32
//     I6464Filed  int64
//     F32Filed    float32
//     F64Filed    float64
//     BoolFiled   bool
// }{"string filed value", 123, 32, 64, 32.123, 64.123, true})
// logger.I("sets:", sets)
// [CODE]
func (w *WingProvider) FormatSets(updates interface{}) string {
	sets := []string{}

	keys, values := reflect.TypeOf(updates), reflect.ValueOf(updates)
	for i := 0; i < keys.NumField(); i++ {
		name, value := keys.Field(i).Name, values.Field(i).Interface()
		switch value.(type) {
		case string:
			sets = append(sets, fmt.Sprintf(name+"='%s'", value.(string)))
		case int:
			sets = append(sets, fmt.Sprintf(name+"=%d", value.(int)))
		case int32:
			sets = append(sets, fmt.Sprintf(name+"=%d", value.(int32)))
		case int64:
			sets = append(sets, fmt.Sprintf(name+"=%d", value.(int64)))
		case float32:
			sets = append(sets, fmt.Sprintf(name+"=%v", value.(float32)))
		case float64:
			sets = append(sets, fmt.Sprintf(name+"=%v", value.(float64)))
		case bool:
			sets = append(sets, fmt.Sprintf(name+"=%v", value.(bool)))
		}
	}
	return strings.Join(sets, ", ")
}
