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
	"time"
)

// Standery time layout format at second minimum
const TimeLayout string = "2006-01-02 15:04:05"

// IsSameDay equal given days string based on TimeLayout
func IsSameDay(src string, des string) bool {
	st, _ := time.Parse(TimeLayout, src)
	dt, _ := time.Parse(TimeLayout, des)
	return st.Day() == dt.Day()
}

// IsToday check the given day string if today
func IsToday(des string) bool {
	dt, _ := time.Parse(TimeLayout, des)
	return time.Now().Day() == dt.Day()
}
