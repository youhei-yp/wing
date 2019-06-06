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
	return w.Conn.Query(query, args)
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
