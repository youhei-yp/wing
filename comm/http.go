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
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/youhei-yp/wing/invar"
	"github.com/youhei-yp/wing/logger"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

const (
	// ContentTypeJson json content type
	ContentTypeJson = "application/json;charset=UTF-8"

	// ContentTypeForm form content type
	ContentTypeForm = "application/x-www-form-urlencoded"
)

// EncodeUrl encode url params
func EncodeUrl(rawurl string) string {
	enurl, err := url.Parse(rawurl)
	if err != nil {
		logger.E("Encode urlm err:", err)
		return rawurl
	}
	enurl.RawQuery = enurl.Query().Encode()
	return enurl.String()
}

// HttpGet handle http get method
func HttpGet(tagurl string, params ...interface{}) ([]byte, error) {
	if len(params) > 0 {
		tagurl = fmt.Sprintf(tagurl, params...)
	}

	rawurl := EncodeUrl(tagurl)
	resp, err := http.Get(rawurl)
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
	logger.I("Handled http get:", tagurl)
	return body, nil
}

// HttpPost handle http post method, you can set content type as
// comm.ContentTypeJson or comm.ContentTypeForm, or other you need set.
// [CODE:]
//     // set post data as json string
//     data := struct {"key": "Value", "id": "123"}
//     resp, err := comm.HttpPost(tagurl, data)
//
//     // set post data as form string
//     data := "key=Value&id=123"
//     resp, err := comm.HttpPost(tagurl, data, comm.ContentTypeForm)
// [CODE]
func HttpPost(tagurl string, postdata interface{}, contentType ...string) ([]byte, error) {
	ct := ContentTypeJson
	if len(contentType) > 0 {
		ct = contentType[0]
	}

	switch ct {
	case ContentTypeJson:
		return httpPostJson(tagurl, postdata)
	case ContentTypeForm:
		return httpPostForm(tagurl, postdata.(url.Values))
	}
	return nil, invar.ErrInvalidParams
}

// HttpGetString call HttpGet and trim " char both begin and end
func HttpGetString(tagurl string, params ...interface{}) (string, error) {
	resp, err := HttpGet(tagurl, params...)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(resp), "\""), nil
}

// HttpPostString call HttpPost and trim " char both begin and end.
func HttpPostString(tagurl string, postdata interface{}, contentType ...string) (string, error) {
	resp, err := HttpPost(tagurl, postdata, contentType...)
	if err != nil {
		return "", err
	}
	return strings.Trim(string(resp), "\""), nil
}

// HttpGetStruct handle http get method and unmarshal data to struct object
func HttpGetStruct(tagurl string, out interface{}, params ...interface{}) error {
	body, err := HttpGet(tagurl, params...)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, out); err != nil {
		logger.E("Unmarshal bady to struct err:", err)
		return err
	}
	return nil
}

// HttpPostStruct handle http post method and unmarshal data to struct object
func HttpPostStruct(tagurl string, postdata, out interface{}, contentType ...string) error {
	body, err := HttpPost(tagurl, postdata, contentType...)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, out); err != nil {
		logger.E("Unmarshal bady to struct err:", err)
		return err
	}
	return nil
}

// HttpClientGet handle get by http.Client, you can set useTLS to enable TLS or not
func HttpClientGet(tagurl string, useTLS bool, params ...interface{}) ([]byte, error) {
	if len(params) > 0 {
		tagurl = fmt.Sprintf(tagurl, params...)
	}

	rawurl := EncodeUrl(tagurl)
	req, err := http.NewRequest("GET", rawurl, http.NoBody)
	if err != nil {
		logger.E("Create http request err:", err)
		return nil, err
	}
	return httpClientDo(req, useTLS)
}

// HttpClientGet handle post by http.Client, you can set useTLS to enable TLS or not
func HttpClientPost(tagurl string, useTLS bool, postdata ...interface{}) ([]byte, error) {
	var body io.Reader
	if len(postdata) > 0 {
		params, err := json.Marshal(postdata)
		if err != nil {
			logger.E("Marshal post data err:", err)
			return nil, err
		}
		body = bytes.NewReader(params)
	} else {
		body = http.NoBody
	}

	req, err := http.NewRequest("POST", tagurl, body)
	if err != nil {
		logger.E("Create http request err:", err)
		return nil, err
	}
	req.Header.Set("Content-Type", ContentTypeJson)
	return httpClientDo(req, useTLS)
}

// HttpClientGetStruct handle http get method and unmarshal data to struct object
func HttpClientGetStruct(tagurl string, useTLS bool, out interface{}, params ...interface{}) error {
	body, err := HttpClientGet(tagurl, useTLS, params...)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, out); err != nil {
		logger.E("Unmarshal bady to struct err:", err)
		return err
	}
	return nil
}

// HttpClientPostStruct handle http post method and unmarshal data to struct object
func HttpClientPostStruct(tagurl string, useTLS bool, out interface{}, postdata ...interface{}) error {
	body, err := HttpClientPost(tagurl, useTLS, postdata...)
	if err != nil {
		return err
	}

	if err := json.Unmarshal(body, out); err != nil {
		logger.E("Unmarshal bady to struct err:", err)
		return err
	}
	return nil
}

// httpPostJson http post method, you can set post data as json struct.
func httpPostJson(tagurl string, postdata interface{}) ([]byte, error) {
	params, err := json.Marshal(postdata)
	if err != nil {
		logger.E("Marshal post data err:", err)
		return nil, err
	}

	resp, err := http.Post(tagurl, ContentTypeJson, bytes.NewReader(params))
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
	logger.I("Handled http post:", tagurl, "params:", postdata)
	return body, nil
}

// httpPostForm http post method, you can set post data as url.Values.
func httpPostForm(tagurl string, postdata url.Values) ([]byte, error) {
	resp, err := http.PostForm(tagurl, postdata)
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
	logger.I("Handled http post:", tagurl, "params:", postdata)
	return body, nil
}

// httpClientDo handle http client DO method, and return response.
func httpClientDo(req *http.Request, useTLS bool) ([]byte, error) {
	client := &http.Client{
		Transport: &http.Transport{
			TLSClientConfig: &tls.Config{
				InsecureSkipVerify: !useTLS,
			},
		},
	}

	resp, err := client.Do(req)
	if err != nil {
		logger.E("Execute client DO err:", err)
		return nil, err
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.E("Read client DO response err:", err)
		return nil, err
	}
	logger.D("Handled http client do")
	return body, nil
}
