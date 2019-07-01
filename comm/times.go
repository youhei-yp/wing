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
	"bytes"
	"strconv"
	"time"
)

const (
	// DateLayout standery date layout format at day minimum
	DateLayout = "2006-01-02"

	// TimeLayout standery time layout format at second minimum
	TimeLayout = "2006-01-02 15:04:05"

	// HourLayout standery time layout format at second minimum
	HourLayout = "15:04:05"

	// MSLayout standery time layout format at million second minimum
	MSLayout = "2006-01-02 15:04:05.000"
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
	deslayout := time.Unix(des, 0).Format(DateLayout)
	return IsToday(deslayout)
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

// TodayUnix return today unix time at 0:00:00
func TodayUnix() int64 {
	now := time.Now().Format(DateLayout)
	st, _ := time.Parse(DateLayout, now)
	return st.Unix()
}

// fill3Digits add zero for number > 10
func fill2Digits(input int) string {
	if input < 10 {
		return "0" + strconv.Itoa(input)
	}
	return strconv.Itoa(input)
}

// fill3Digits add zero for number > 10 or 100
func fill3Digits(input int) string {
	if input < 10 {
		return "00" + strconv.Itoa(input)
	}
	if input < 100 {
		return "0" + strconv.Itoa(input)
	}
	return strconv.Itoa(input)
}

// GetHumanReadableDuration return readable time during start to end: 12:12:12
func GetHumanReadableDuration(start time.Time, end time.Time) string {
	v := end.Unix() - start.Unix() // seconds
	h := v / 3600
	m := v % 3600 / 60
	s := v % 60
	return fill2Digits(int(h)) + ":" + fill2Digits(int(m)) + ":" + fill2Digits(int(s))
}

// GetLongHumanReadableDuration return readable time during start to end: 2d 6h 25m 48s
func GetLongHumanReadableDuration(start time.Time, end time.Time) string {
	v := int(end.Unix() - start.Unix()) // seconds
	return strconv.Itoa(v/86400) + "d " + strconv.Itoa(v%86400/3600) + "h " + strconv.Itoa(v%3600/60) + "m " + strconv.Itoa(v%60) + "s"
}
