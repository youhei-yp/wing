// Copyright (c) 2018-2019 Xunmo All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2018/12/01   youhei         New version
// -------------------------------------------------------------------

package captcha

import (
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wing/logger"
)

const (
	time_layout       = "2006-01-02T15:04:05Z"
	sort_query_format = "AccessKeyId=%s" +
		"&Action=SendSms" +
		"&Format=JSON" +
		"&OutId=123" +
		"&PhoneNumbers=%s" +
		"&RegionId=cn-hangzhou" +
		"&SignName=%s" +
		"&SignatureMethod=HMAC-SHA1" +
		"&SignatureNonce=%s" +
		"&SignatureVersion=1.0" +
		"&TemplateCode=%s" +
		"&TemplateParam=%s" +
		"&Timestamp=%s" +
		"&Version=2017-05-25"
)

// SmsSender sender, including smtp authtication and user info
type SmsSender struct {
	accessSecret, accessKeyID string
	requestUrlFormat          string
}

// encodeUrl replace encode string to use in web transation
func (s *SmsSender) encodeUrl(src string) string {
	ue := url.QueryEscape(src)
	ue = strings.Replace(ue, "+", "%%20", -1)
	ue = strings.Replace(ue, "*", "%2A", -1)
	ue = strings.Replace(ue, "%%7E", "~", -1)
	ue = strings.Replace(ue, "/", "%%2F", -1)
	return ue
}

// execHttpGet executes http get method
func (s *SmsSender) execHttpGet(requesturl string) ([]byte, error) {
	u, err := url.Parse(requesturl)
	if err != nil {
		logger.E("Parse request url:", requesturl, "err:", err)
		return nil, err
	}
	u.RawQuery = u.Query().Encode()

	rs := []rune(u.String())
	length := len(rs)
	geturl := string(rs[0 : length-3])

	logger.D("Execute http get, url:", geturl)
	res, err := http.Get(geturl)
	if err != nil {
		logger.E("Execute get method err:", err)
		return nil, err
	}
	defer res.Body.Close()

	// read executed response data
	rst, err := ioutil.ReadAll(res.Body)
	if err != nil {
		logger.E("Read get method response err:", err)
		return nil, err
	}
	return rst, nil
}

// getQueryString parse sms request url query string
func (s *SmsSender) getQueryString(accessID, phones, sign, tplcode, content string) string {
	nonce, _ := uuid.NewV4()
	timestamp := url.QueryEscape(time.Now().UTC().Format(time_layout))
	return fmt.Sprintf(sort_query_format,
		accessID,  // access key id of aiyun
		phones,    // target phone numbers to send to
		sign,      // signature name
		nonce,     // signature nonce
		tplcode,   // sms template code
		content,   // sms content
		timestamp, // send timestamp
	)
}

func (s *SmsSender) Send(accessID, phones, sign, tplcode, content string) error {
	// type SMSResponse struct {
	// 	Message   string `json:"Message"`
	// 	RequestId string `json:"RequestId"`
	// 	BizId     string "BizId"
	// 	Code      string "Code"
	// }

	// tplcode := sms.TemplateCode
	// signName := url.QueryEscape(sms.SignName)
	// content := url.QueryEscape(fmt.Spintf(sms.TemplateFormat, code)) // "{\"code\":\"888123\"}"

	// queryString := getQueryString(access_key_id, phones, signName, tplcode, content)

	// key := []byte(access_secret)
	// sign := "GET&%%2F&" + encodeUrl(queryString)

	// mac := hmac.New(sha1.New, key)
	// mac.Write([]byte(sign))

	// signture := encodeUrl(base64.StdEncoding.EncodeToString(mac.Sum(nil)))
	// requesturl := fmt.Sprintf(request_url_format, signture, queryString)
	// logger.D("Send sms, request url:", requesturl)

	// resp, err := execHttpGet(requesturl)

	// fmt.Printf("result:%s", resp)

	// result := &SMSResponse{}
	// json.Unmarshal(resp, result)

	// logger.D("jsonresult:", result)
	// if result.Message != "OK" {
	// 	logger.E("send err:", result.Message)
	// 	return false
	// }
	return nil
}

// NewSmsSender create a sms sender for given cloud service
func NewSmsSender(secret, keyid, requrl string) *SmsSender {
	sender := &SmsSender{
		accessSecret: secret, accessKeyID: keyid, requestUrlFormat: requrl,
	}
	return sender
}
