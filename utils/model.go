// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package utils

import (
	"database/sql"
	"github.com/astaxie/beego/orm"
	"strings"
	"wing/logger"
)

// OrmUtil : provides a named orm object to query or exec sql commend
type OrmUtil struct {
	Ormer   orm.Ormer // ormer object
	ormName string    // ormer name
}

// NewOrmUtil create a OrmUtil object and than using the given name ormer
func NewOrmUtil(name string) (*OrmUtil, error) {
	ormer := orm.NewOrm()
	if err := ormer.Using(name); err != nil {
		logger.E("Using", name, "orm err:", err)
		return nil, err
	}

	u := &OrmUtil{ormName: name, Ormer: ormer}
	logger.D("Created a orm util as name:", name)
	return u, nil
}

// Query execute query sql data and fill into the given container
func (u *OrmUtil) Query(sqlstr string, container interface{}, args ...interface{}) error {
	if u.Ormer == nil {
		logger.E("Ormer", "["+u.ormName+"]", "not using!")
		return ErrOrmNotUsing
	}

	if err := u.Ormer.Raw(sqlstr, args).QueryRow(container); err != nil {
		if strings.Index(err.Error(), "no row found") != 0 {
			logger.D("No row found!")
			return ErrNoRowFound
		}
		logger.E("Query sql:["+sqlstr+"] err:", err)
		return err
	}
	logger.D("Executed query", "["+sqlstr+"]")
	return nil
}

// Exec handle insert|update|delete sql data
func (u *OrmUtil) Exec(sqlstr string, args ...interface{}) (sql.Result, error) {
	if u.Ormer == nil {
		logger.E("Ormer", "["+u.ormName+"]", "not using!")
		return nil, ErrOrmNotUsing
	}

	result, err := u.Ormer.Raw(sqlstr, args).Exec()
	if err != nil {
		logger.E("Exec sql:["+sqlstr+"] err:", err)
		return nil, err
	}
	logger.D("Executed sql", "["+sqlstr+"]")
	return result, nil
}
