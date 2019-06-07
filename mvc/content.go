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

// AppendLike append like keyword end of sql string,
// DON'T call it after AppendLimit()
func (w *WingProvider) AppendLike(query, filed, keyword string, and ...bool) string {
	if len(and) > 0 && and[0] {
		return query + " AND " + filed + " LIKE '%%" + keyword + "%%'"
	}
	return query + " WHERE " + filed + " LIKE '%%" + keyword + "%%'"
}

// AppendLimit append page limitation end of sql string,
// DON'T call it before AppendLick()
func (w *WingProvider) AppendLimit(query string, page int) string {
	offset, items := page*limitPageItems, limitPageItems
	return query + " LIMIT " + fmt.Sprintf("%d, %d", offset, items)
}

// AppendLikeLimit append like keyword and limit end of sql string
func (w *WingProvider) AppendLikeLimit(query, filed, keyword string, page int, and ...bool) string {
	return w.AppendLimit(w.AppendLike(query, filed, keyword, and...), page)
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
//     EmptyString string
//     BlankString string
//     TrimString  string
//     IntFiled    int
//     I32Filed    int32
//     I64Filed    int64
//     F32Filed    float32
//     F64Filed    float64
//     BoolFiled   bool
// }{"string", "", " ", " trim ", 123, 32, 64, 32.123, 64.123, true})
//
// //sets: stringfiled='string', trimstring='trim', intfiled=123, i32filed=32, i64filed=64, f32filed=32.123, f64filed=64.123, boolfiled=true
// logger.I("sets:", sets)
// [CODE]
func (w *WingProvider) FormatSets(updates interface{}) string {
	sets := []string{}
	keys, values := reflect.TypeOf(updates), reflect.ValueOf(updates)
	for i := 0; i < keys.NumField(); i++ {
		name := strings.ToLower(keys.Field(i).Name)
		if name == "" {
			continue
		}

		value := values.Field(i).Interface()
		switch value.(type) {
		case string:
			trimvalue := strings.Trim(value.(string), " ")
			if trimvalue != "" {
				sets = append(sets, fmt.Sprintf(name+"='%s'", trimvalue))
			}
		case int, int8, int16, int32, int64, float32, float64, bool:
			sets = append(sets, fmt.Sprintf(name+"=%v", value))
		}
	}
	return strings.Join(sets, ", ")
}
