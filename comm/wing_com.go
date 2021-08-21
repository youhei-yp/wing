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
	"encoding/json"
	"errors"
	"fmt"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"github.com/huichen/sego"
	"github.com/mozillazg/go-pinyin"
	"github.com/youhei-yp/wing/logger"
	"os"
	"os/signal"
	"regexp"
	"strconv"
	"syscall"
	"unicode"
)

// Variables of Sego
var Segmenter sego.Segmenter

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

// ToMap transform given struct data to map data, the transform struct
// feilds must using json tag to mark the map key.
//	[CODE:]
//	type struct Sample {
//		Name string `json:"name"`
//	}
//	d := Sample{ Name : "name_value" }
//	md, _ := comm.ToMap(d)
//	// md data format is {
//	//     "name" : "name_value"
//	// }
//	[CODE]
func ToMap(input interface{}) (map[string]interface{}, error) {
	out := make(map[string]interface{})
	structbuf, err := json.Marshal(input)
	if err != nil {
		logger.E("Marshal input struct err:", err)
		return nil, err
	}

	// json buffer decode to map
	d := json.NewDecoder(bytes.NewReader(structbuf))
	d.UseNumber()
	if err = d.Decode(&out); err != nil {
		logger.E("Decode json data to map err:", err)
		return nil, err
	}

	return out, nil
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

// GetKeyWords get more search keywords
func GetKeyWords(str string) []string {
	segments := Segmenter.Segment([]byte(str))

	// use search mode 'true' to get more search keywords
	words := sego.SegmentsToSlice(segments, true)
	var keywords []string
	for _, v := range words {
		if v == " " {
			continue
		}
		keywords = append(keywords, v)
	}
	return keywords
}

// RemoveDuplicate remove duplicate data from array
func RemoveDuplicate(oldArr []string) []string {
	newArr := make([]string, 0)
	for i := 0; i < len(oldArr); i++ {
		repeat := false
		for j := i + 1; j < len(oldArr); j++ {
			if oldArr[i] == oldArr[j] {
				repeat = true
				break
			}
		}
		if !repeat {
			newArr = append(newArr, oldArr[i])
		}
	}
	return newArr
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

// AccessAllowOriginBy allow cross domain access for the given origins
func AccessAllowOriginBy(category int, origins string) {
	beego.InsertFilter("*", category, cors.Allow(&cors.Options{
		AllowAllOrigins: true,
		AllowOrigins:    []string{origins}, // use to set allow Origins
		AllowMethods:    []string{"GET", "POST", "OPTIONS"},
		AllowHeaders:    []string{"Origin", "Authorization", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
		ExposeHeaders:   []string{"Content-Length", "Access-Control-Allow-Origin", "Access-Control-Allow-Headers", "Content-Type"},
	}))
}

// AccessAllowOriginByLocal allow cross domain access for localhost,
// the port number must config in /conf/app.conf file like :
//	~~~~~~
//	; Server port of HTTP
//	httpport=3200
//	~~~~~~
func AccessAllowOriginByLocal(category int) {
	if beego.BConfig.Listen.HTTPPort > 0 {
		localhosturl := fmt.Sprintf("http://127.0.0.1:%v/", beego.BConfig.Listen.HTTPPort)
		AccessAllowOriginBy(category, localhosturl)
	}
}
