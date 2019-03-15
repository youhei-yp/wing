// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package content

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql" // mysql driver
	"wing/logger"
	"wing/utils"
)

// SQLUtil : provides a named orm object to query or exec sql commend
type SQLUtil struct {
	db *sql.DB
}

// RowScanedFunc return row scaned result function
type RowScanedFunc func(...interface{})

// NewSQLUtil create a OrmUtil object and than using the given name ormer
// @params: option[0] tcp, udp protocols
//          option[1] database host ip address
//          option[2] database host port
func NewSQLUtil(user, password, dbname string, option ...string) (*SQLUtil, error) {
	dsn := ""
	if len(option) == 3 {
		protocol, host, port := option[0], option[1], option[2]
		dsn = fmt.Sprintf("%s:%s@%s(%s:%s)/%s?charset=utf8", user, password, protocol, host, port, dbname)
	} else {
		dsn = fmt.Sprintf("%s:%s@/%s?charset=utf8", user, password, dbname)
	}

	con, err := sql.Open("mysql", dsn)
	if err != nil {
		logger.E("Open mysql database err:", err)
		return nil, err
	}

	if err = con.Ping(); err != nil {
		logger.E("Ping test err:", err)
		return nil, err
	}
	logger.I("Connected database:", dsn)

	con.SetMaxIdleConns(100)
	con.SetMaxOpenConns(100)
	return &SQLUtil{db: con}, nil
}

// PrepareQuery execute query sql data and fill into the given container
// @notice : you must use ScanRow or ScanRows to scan query data
func (u *SQLUtil) PrepareQuery(sqlstr string, args ...interface{}) (*sql.Rows, error) {
	rows, err := u.db.Query(sqlstr, args)
	if err != nil {
		logger.E("Query sql["+sqlstr+"] err:", err)
		return nil, err
	}
	logger.I("Query:["+sqlstr+"] args:", args)
	return rows, nil
}

// ScanRow scan and get the top row columns
func (u *SQLUtil) ScanRow(rows *sql.Rows, outer ...interface{}) error {
	defer rows.Close()
	if !rows.Next() {
		logger.E("Query sql err: No row data!")
		return utils.ErrNoRowFound
	}
	rows.Columns()

	if err := rows.Scan(outer...); err != nil {
		logger.E("Scan sql err:", err)
		return err
	}
	logger.I("Scaned row result:", outer)
	return nil
}

// ScanRows scan and get all rows columns, each row returned by RowScanedFunc
func (u *SQLUtil) ScanRows(rows *sql.Rows, scanedFunc RowScanedFunc, temp ...interface{}) error {
	defer rows.Close()
	haveRows := false

	for rows.Next() {
		haveRows = true
		rows.Columns()
		if err := rows.Scan(temp...); err != nil {
			logger.E("Scan rows err:", err)
			return err
		}
		if scanedFunc != nil {
			logger.I("Scaned row result:", temp)
			scanedFunc(temp...)
		}
	}

	if !haveRows {
		logger.E("Query sql err: No row data!")
		return utils.ErrNoRowFound
	}
	return nil
}

// Exec handle insert|update|delete sql action
func (u *SQLUtil) Exec(sqlstr string, args ...interface{}) (sql.Result, error) {
	stmt, err := u.db.Prepare(sqlstr)
	if err != nil {
		logger.E("Prepare sql:["+sqlstr+"] err:", err)
		return nil, err
	}
	defer stmt.Close()

	result, err := stmt.Exec(args...)
	if err != nil {
		logger.E("Exec sql:["+sqlstr+"] err: ", err.Error())
		return nil, err
	}

	logger.I("Executed sql:[" + sqlstr + "]")
	return result, nil
}
