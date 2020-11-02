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
	"errors"
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"github.com/youhei-yp/wing/logger"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"
	"unicode"
)

// Try try-catch-finaly method
func Try(do func(), catcher func(error), finaly ...func()) {
	defer func() {
		if i := recover(); i != nil {
			execption := errors.New(fmt.Sprint(i))
			logger.E("catched exception:", i)
			catcher(execption)
			if len(finaly) > 0 {
				finaly[0]()
			}
		}
	}()
	do()
}

// Ternary ternary operation
func Ternary(condition bool, trueResult interface{}, falseResult interface{}) interface{} {
	if condition {
		return trueResult
	}
	return falseResult
}

// To2Digits fill zero if input digit not enough 2
func To2Digits(input interface{}) string {
	return fmt.Sprintf("%02d", input)
}

// To2Digits fill zero if input digit not enough 3
func To3Digits(input interface{}) string {
	return fmt.Sprintf("%03d", input)
}

// ToNDigits fill zero if input digit not enough N
func ToNDigits(input interface{}, n int) string {
	return fmt.Sprintf("%0"+strconv.Itoa(n)+"d", input)
}

// IgnoreSysSignalPIPE ignore system PIPE signal
func IgnoreSysSignalPIPE() {
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGPIPE)
	go func() {
		for {
			select {
			case sig := <-sc:
				if sig == syscall.SIGPIPE {
					logger.E("!! IGNORE BROKEN PIPE SIGNAL !!")
				}
			}
		}
	}()
}

// GetSortKey get first letter of Chinese Pinyin
func GetSortKey(str string) string {
	if str == "" { // check the input param
		return "*"
	}

	// get the first char and verify if it is a~Z char
	firstChar, sortKey := []rune(str)[0], ""
	isAZchar, err := regexp.Match("[a-zA-Z]", []byte(str))
	if err != nil {
		logger.E("Regexp match err:", err)
		return "*"
	}

	if isAZchar {
		sortKey = string(unicode.ToUpper(firstChar))
	} else {
		if unicode.Is(unicode.Han, firstChar) { // chinese
			str1 := pinyin.LazyConvert(string(firstChar), nil)
			s := []rune(str1[0])
			sortKey = string(unicode.ToUpper(s[0]))
		} else if unicode.IsNumber(firstChar) { // number
			sortKey = string(firstChar)
		} else { // other language
			sortKey = "#"
		}
	}
	return sortKey
}
