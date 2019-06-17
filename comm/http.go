// Copyright (c) 2018-2019 Dunyu All Rights Reserved.
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
	"encoding/json"
	"fmt"
	"github.com/youhei-yp/wing/logger"
	"io/ioutil"
	"net/http"
)

const (
	// ContentTypeJson json content type
	ContentTypeJson = "application/json;charset=UTF-8"
)

// HttpGet handle http get method
func HttpGet(url string, params ...interface{}) ([]byte, error) {
	if len(params) > 0 {
		url = fmt.Sprintf(url, params...)
	}

	resp, err := http.Get(url)
	if err != nil {
		logger.E("Handle http get err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.E("Read get response err:", err)
		return nil, err
	}
	logger.I("Handled http get:", url)
	return body, nil
}

// HttpPost handle http post method
func HttpPost(url string, postdata interface{}) ([]byte, error) {
	params, err := json.Marshal(postdata)
	if err != nil {
		logger.E("Marshal post data err:", err)
		return nil, err
	}

	resp, err := http.Post(url, ContentTypeJson, bytes.NewReader(params))
	if err != nil {
		logger.E("Handle http post err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.E("Read post response err:", err)
		return nil, err
	}
	logger.I("Handled http post:", url, "params:", postdata)
	return body, nil
}

// HttpGetStruct handle http get method and unmarshal data to struct object
func HttpGetStruct(url string, out interface{}, params ...interface{}) error {
	body, err := HttpGet(url, params...)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, out); err != nil {
		logger.E("Unmarshal dady to struct err:", err)
		return err
	}
	return nil
}

// HttpPostStruct handle http post method and unmarshal data to struct object
func HttpPostStruct(url string, postdata, out interface{}) error {
	body, err := HttpPost(url, postdata)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, out); err != nil {
		logger.E("Unmarshal dady to struct err:", err)
		return err
	}
	return nil
}
