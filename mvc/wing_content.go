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

// ScanCallback use for scan query result from rows
type ScanCallback func(rows *sql.Rows) error

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
// the connections holded by mvc.WingHelper object,
// the charset maybe 'utf8' or 'utf8mb4' same as database set.
func OpenDatabase(charset string) error {
	dsn := fmt.Sprintf("%s:%s@/%s?charset=%s",
		beego.AppConfig.String("dbuser"), beego.AppConfig.String("dbpwd"),
		beego.AppConfig.String("dbname"), charset)

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

// IsEmpty call sql.Query() to check target data if exist
func (w *WingProvider) IsEmpty(query string, args ...interface{}) (bool, error) {
	rows, err := w.Conn.Query(query, args...)
	if err != nil {
		return false, err
	}
	defer rows.Close()
	return !rows.Next(), nil
}

// QueryOne call sql.Query() to query one record
func (w *WingProvider) QueryOne(query string, cb ScanCallback, args ...interface{}) error {
	rows, err := w.Conn.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	if !rows.Next() {
		return invar.ErrNotFound
	}
	rows.Columns()
	return cb(rows)
}

// QueryArray call sql.Query() to query multi records
func (w *WingProvider) QueryArray(query string, cb ScanCallback, args ...interface{}) error {
	rows, err := w.Conn.Query(query, args...)
	if err != nil {
		return err
	}
	defer rows.Close()

	for rows.Next() {
		rows.Columns()
		if err := cb(rows); err != nil {
			return err
		}
	}
	return nil
}

// Insert call sql.Prepare() and stmt.Exec() to insert a new record
func (w *WingProvider) Insert(query string, args ...interface{}) (int64, error) {
	stmt, err := w.Conn.Prepare(query)
	if err != nil {
		return -1, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return -1, err
	}
	return result.LastInsertId()
}

// Execute call sql.Prepare() and stmt.Exec() to update or delete records
func (w *WingProvider) Execute(query string, args ...interface{}) error {
	stmt, err := w.Conn.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		return err
	}
	return w.Affected(result)
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
		case bool:
			sets = append(sets, fmt.Sprintf(name+"=%v", value))
		case invar.Bool:
			boolvalue := value.(invar.Bool)
			if boolvalue != invar.BNone {
				truevalue := (boolvalue == invar.BTrue)
				sets = append(sets, fmt.Sprintf(name+"=%v", truevalue))
			}
		case string:
			trimvalue := strings.Trim(value.(string), " ")
			if trimvalue != "" { // filter empty string fields
				sets = append(sets, fmt.Sprintf(name+"='%s'", trimvalue))
			}
		case int, int8, int16, int32, int64, float32, float64,
			invar.Status, invar.Box, invar.Role, invar.Limit, invar.Lang, invar.Kind:
			if fmt.Sprintf("%v", value) != "0" { // filter 0 fields
				sets = append(sets, fmt.Sprintf(name+"=%v", value))
			}
		}
	}
	return strings.Join(sets, ", ")
}

// Atomicity call sql.Begin() , tx.Rollback() and tx.Commit() to start one transaction
// All operations in a transaction are either completed or not completed. They will not end in an intermediate link.
// If an error occurs during the execution of the transaction, it will be rolled back to the state before the transaction starts
// [CODE:]
// 		args := make(map[string][]interface{})
//		args[query] = []interface{}{...arg}
// [CODE]
func (w *WingProvider) Atomicity(args map[string][]interface{}) error {
	atomic, err := w.Conn.Begin()
	if err != nil {
		return err
	}
	defer atomic.Rollback()
	for query, arg := range args {
		result, err := atomic.Exec(query, arg...)
		if err != nil {
			return err
		}
		if err = w.Affected(result); err != nil {
			return err
		}
	}
	return atomic.Commit()
}
