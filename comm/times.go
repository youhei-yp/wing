// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2019/05/22   yangping       New version
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

// GetDateString get short date string: 2018-11-11
func GetDateString(t time.Time) string {
	var buff bytes.Buffer
	buff.WriteString(strconv.Itoa(GetYear(t)))
	buff.WriteString("-")
	buff.WriteString(fill2Digits(GetMonth(t)))
	buff.WriteString("-")
	buff.WriteString(fill2Digits(GetDay(t)))
	return buff.String()
}

// GetLongDateString get long date string: 2018-11-11 12:12:12
func GetLongDateString(t time.Time) string {
	var buff bytes.Buffer
	buff.WriteString(strconv.Itoa(GetYear(t)))
	buff.WriteString("-")
	buff.WriteString(fill2Digits(GetMonth(t)))
	buff.WriteString("-")
	buff.WriteString(fill2Digits(GetDay(t)))
	buff.WriteString(" ")
	buff.WriteString(fill2Digits(GetHour(t)))
	buff.WriteString(":")
	buff.WriteString(fill2Digits(GetMinute(t)))
	buff.WriteString(":")
	buff.WriteString(fill2Digits(GetSecond(t)))
	return buff.String()
}

// GetShortDateString get long date string: 2018-11-11 12:12:12
func GetShortDateString(t time.Time) string {
	var buff bytes.Buffer
	buff.WriteString(fill2Digits(GetHour(t)))
	buff.WriteString(":")
	buff.WriteString(fill2Digits(GetMinute(t)))
	buff.WriteString(":")
	buff.WriteString(fill2Digits(GetSecond(t)))
	return buff.String()
}

// GetLongLongDateString get long date string: 2018-11-11 12:12:12,233
func GetLongLongDateString(t time.Time) string {
	var buff bytes.Buffer
	buff.WriteString(strconv.Itoa(GetYear(t)))
	buff.WriteString("-")
	buff.WriteString(fill2Digits(GetMonth(t)))
	buff.WriteString("-")
	buff.WriteString(fill2Digits(GetDay(t)))
	buff.WriteString(" ")
	buff.WriteString(fill2Digits(GetHour(t)))
	buff.WriteString(":")
	buff.WriteString(fill2Digits(GetMinute(t)))
	buff.WriteString(":")
	buff.WriteString(fill2Digits(GetSecond(t)))
	buff.WriteString(",")
	buff.WriteString(fill3Digits(GetMillionSecond(t)))
	return buff.String()
}

// GetTimestamp get current timestamp in milliseconds.
func GetTimestamp(t time.Time) int64 {
	return t.UnixNano() / 1e6
}

// CreateTime create time by millis.
func CreateTime(millis int64) time.Time {
	return time.Unix(millis, 0)
}

// GetNanosecond get current timestamp in Nanosecond.
func GetNanosecond(t time.Time) int64 {
	return t.UnixNano()
}

// GetYear get year of t.
func GetYear(t time.Time) int {
	return t.Year()
}

// GetYear get year of t.
func GetMonth(t time.Time) int {
	return int(t.Month())
}

// GetYear get day of t.
func GetDay(t time.Time) int {
	return t.Day()
}

// GetYear get hour of t.
func GetHour(t time.Time) int {
	return t.Hour()
}

// GetYear get minute of t.
func GetMinute(t time.Time) int {
	return t.Minute()
}

// GetYear get second of t.
func GetSecond(t time.Time) int {
	return t.Second()
}

// GetYear get million second of t.
func GetMillionSecond(t time.Time) int {
	return t.Nanosecond() / 1e6
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
	v := GetTimestamp(end)/1000 - GetTimestamp(start)/1000 // seconds
	h := v / 3600
	m := v % 3600 / 60
	s := v % 60
	return fill2Digits(int(h)) + ":" + fill2Digits(int(m)) + ":" + fill2Digits(int(s))
}

// GetLongHumanReadableDuration return readable time during start to end: 2d 6h 25m 48s
func GetLongHumanReadableDuration(start time.Time, end time.Time) string {
	v := int(GetTimestamp(end)/1000 - GetTimestamp(start)/1000) // seconds
	return strconv.Itoa(v/86400) + "d " + strconv.Itoa(v%86400/3600) + "h " + strconv.Itoa(v%3600/60) + "m " + strconv.Itoa(v%60) + "s"
}
