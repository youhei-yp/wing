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
	"fmt"
	"time"
)

// Day, week, month, quarter, year duration on nanosecond
const (
	Day     = time.Hour * 24
	Week    = Day * 7
	Month   = Day * 30
	Quarter = Month * 3
	Year    = Day * 365
)

// Day, week, month, quarter, year duration on millisecond
const (
	DayMs     = Day / time.Millisecond
	WeekMs    = Week / time.Millisecond
	MonthMs   = Month / time.Millisecond
	QuarterMs = Quarter / time.Millisecond
	YearMs    = Year / time.Millisecond
)

// Day, week, month, quarter, year duration on second
const (
	DaySeconds     = Day / time.Second
	WeekSeconds    = Week / time.Second
	MonthSeconds   = Month / time.Second
	QuarterSeconds = Quarter / time.Second
	YearSeconds    = Year / time.Second
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

// Today return today unix time
func Today() int64 {
	return time.Now().Unix()
}

// Yesterday return yesterday unix time base on current
func Yesterday() int64 {
	return time.Now().AddDate(0, 0, -1).Unix()
}

// Tommorrow return tommorrow unix time srart from current
func Tommorrow() int64 {
	return time.Now().AddDate(0, 0, 1).Unix()
}

// NextWeek return next week unix time start from current
func NextWeek() int64 {
	return time.Now().AddDate(0, 0, 1).Unix()
}

// NextMonth return next month unix time start from current
func NextMonth() int64 {
	return time.Now().AddDate(0, 1, 0).Unix()
}

// NextQuarter return next quarter unix time start from current
func NextQuarter() int64 {
	return time.Now().AddDate(0, 3, 0).Unix()
}

// NextYear return next year unix time start from current
func NextYear() int64 {
	return time.Now().AddDate(1, 0, 0).Unix()
}

// NextTime return next unix time start from current
func NextTime(duration time.Duration) int64 {
	return time.Now().Add(duration).Unix()
}

// TodayUnix return today unix time at 0:00:00
func TodayUnix() int64 {
	now := time.Now().Format(DateLayout)
	st, _ := time.Parse(DateLayout, now)
	return st.Unix()
}

// YesterdayUnix return yesterday unix time at 0:00:00
func YesterdayUnix() int64 {
	return NextUnix(-Day)
}

// TommorrowUnix return tommorrow unix time at 0:00:00
func TommorrowUnix() int64 {
	return NextUnix(Day)
}

// WeekUnix return next week unix time at 0:00:00
func WeekUnix() int64 {
	return NextUnix(Week)
}

// MonthUnix return next week unix time at 0:00:00
func MonthUnix() int64 {
	return NextUnix2(0, 1, 0)
}

// QuarterUnix return next quarter unix time at 0:00:00
func QuarterUnix() int64 {
	return NextUnix2(0, 3, 0)
}

// YearUnix return next year unix time at 0:00:00
func YearUnix() int64 {
	return NextUnix2(1, 0, 0)
}

// NextUnix return next 0:00:00 unix time at day after given duration
func NextUnix(duration time.Duration) int64 {
	nt := time.Now().Add(duration).Format(DateLayout)
	st, _ := time.Parse(DateLayout, nt)
	return st.Unix()
}

// NextUnix2 return next 0:00:00 unix time at day after given years, months and days
func NextUnix2(years, months, days int) int64 {
	nt := time.Now().AddDate(years, months, days).Format(DateLayout)
	st, _ := time.Parse(DateLayout, nt)
	return st.Unix()
}

// YearDiff return diff years, months, days
func YearDiff(start, end time.Time) (int, int, int) {
	v := end.Unix() - start.Unix()
	y, m, d := int64(YearSeconds), int64(MonthSeconds), int64(DaySeconds)
	return int(v / y), int(v % y / m), int(v % m / d)
}

// DayDiff return diff days, hours, minutes, seconds
func DayDiff(start, end time.Time) (int, int, int, int) {
	v := int(end.Unix() - start.Unix())
	var d, h, m int = int(DaySeconds), 3600, 60
	return (v / d), (v % d / h), (v % h / m), (v % m)
}

// HourDiff return diff hours, minutes, seconds
func HourDiff(start, end time.Time) (int, int, int) {
	v, h, m := int(end.Unix()-start.Unix()), 3600, 60
	return (v / h), (v % h / m), (v % m)
}

// DurHours return readable time during start to end: 06:25:48,
// you can se the format string, but it must contain 3 %0xd to parse numbers
func DurHours(start, end time.Time, format ...string) string {
	h, m, s := HourDiff(start, end)
	if len(format) > 0 && format[0] != "" {
		return fmt.Sprintf(format[0], h, m, s)
	}
	return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

// DurDays return readable time during start to end: 2d 6h 25m 48s,
// you can set the format string, but it must contain 4 %0xd to parse numbers
func DurDays(start, end time.Time, format ...string) string {
	d, h, m, s := DayDiff(start, end)
	if len(format) > 0 && format[0] != "" {
		return fmt.Sprintf(format[0], d, h, m, s)
	}
	return fmt.Sprintf("%dd %dh %dm %ds", d, h, m, s)
}
