// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
// 00002       2019/06/30   zhaixing       Add function from godfs
// -------------------------------------------------------------------

package comm

import (
	"crypto/md5"
	"encoding/hex"
	"io"
)

func Try(f func(), catcher func(interface{})) {
	defer func() {
		if err := recover(); err != nil {
			catcher(err)
		}
	}()
	f()
}

// TValue ternary operation
func TValue(condition bool, trueValue interface{}, falseValue interface{}) interface{} {
	if condition {
		return trueValue
	}
	return falseValue
}

// Md5sum merge Md5 string to one Md5 string.
func Md5sum(input ...string) string {
	h := md5.New()
	if input != nil {
		for _, v := range input {
			io.WriteString(h, v)
		}
	}
	sliceCipherStr := h.Sum(nil)
	sMd5 := hex.EncodeToString(sliceCipherStr)
	return sMd5
}
