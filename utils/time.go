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

const (
	// DateLayout standery date layout format at day minimum
	DateLayout = "2006-01-02"

	// TimeLayout standery time layout format at second minimum
	TimeLayout = "2006-01-02 15:04:05"
)

// IsToday check the given day string if today
func IsToday(des string) bool {
	now := time.Now().Format(DateLayout)
	st, _ := time.Parse(DateLayout, now)
	dt, _ := time.Parse(DateLayout, des)
	return st.Unix() == dt.Unix()
}

// IsTodayUnix check the given time string if today
func IsTodayUnix(des int64) bool {
	now := time.Now().Format(DateLayout)
	st, _ := time.Parse(DateLayout, now)
	return st.Unix() == des
}

// IsSameDay equal given days string based on TimeLayout
func IsSameDay(src string, des string) bool {
	st, _ := time.Parse(DateLayout, src)
	dt, _ := time.Parse(DateLayout, des)
	return st.Unix() == dt.Unix()
}

// IsSameTime equal given time string based on TimeLayout
func IsSameTime(src string, des string) bool {
	st, _ := time.Parse(TimeLayout, src)
	dt, _ := time.Parse(TimeLayout, des)
	return st.Unix() == dt.Unix()
}

// TodayUnix return today date unix
func TodayUnix() int64 {
	now := time.Now().Format(DateLayout)
	st, _ := time.Parse(DateLayout, now)
	return st.Unix()
}
