// Copyright (c) 2018-2019 WING All Rights Reserved.
//
// Author : yangping
// Email  : youhei_yp@163.com
//
// Prismy.No | Date       | Modified by. | Description
// -------------------------------------------------------------------
// 00001       2020/10/30   yangping       New version
// -------------------------------------------------------------------

package comm

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"errors"
	"fmt"
	"github.com/satori/go.uuid"
	"github.com/youhei-yp/wing/logger"
	"github.com/youhei-yp/wing/secure"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

const (
	timeLayout      = "2006-01-02T15:04:05Z"
	sortQueryFormat = "AccessKeyId=%s" +
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
	requestURLFormat          string
}

// SmsContent sms template
type SmsContent struct {
	SignName        string
	TemplateCode    string
	TemplatePrefix  string
	TemplateSuffix  string
	TemplateContent string
}

// getResult get method response
type respResult struct {
	Message   string `json:"Message"`
	RequestID string `json:"RequestId"`
	// BizID     string `describtion:"BizId"`
	Code string `describtion:"Code"`
}

// encodeURL replace encode string to use in web transation
func (s *SmsSender) encodeURL(src string) string {
	ue := url.QueryEscape(src)
	ue = strings.Replace(ue, "+", "%%20", -1)
	ue = strings.Replace(ue, "*", "%2A", -1)
	ue = strings.Replace(ue, "%%7E", "~", -1)
	ue = strings.Replace(ue, "/", "%%2F", -1)
	return ue
}

// requestRemoteSend executes http get method to request remote send
func (s *SmsSender) requestRemoteSend(requesturl string) ([]byte, error) {
	u, err := url.Parse(requesturl)
	if err != nil {
		logger.E("Parse request url:", requesturl, "err:", err)
		return nil, err
	}
	u.RawQuery = u.Query().Encode()

	rs := []rune(u.String())
	length := len(rs)
	geturl := string(rs[0:length])

	logger.I("Execute http get, url:", geturl)
	res, err := http.Get(geturl)
	if err != nil {
		logger.E("Execute get method err:", err)
		return nil, err
	}

	// read executed response data
	rst, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		logger.E("Read get method response err:", err)
		return nil, err
	}
	return rst, nil
}

// getQueryString parse sms request url query string
func (s *SmsSender) getQueryString(phones, signname, tplcode, content string) string {
	signnonce := uuid.NewV4()
	timestamp := url.QueryEscape(time.Now().UTC().Format(timeLayout))
	return fmt.Sprintf(sortQueryFormat,
		s.accessKeyID, // access key id of aiyun
		phones,        // target phone numbers to send to
		signname,      // signature name
		signnonce,     // signature nonce
		tplcode,       // sms template code
		content,       // sms content
		timestamp,     // send timestamp
	)
}

// Send sends
func (s *SmsSender) Send(phones, signname, tplcode, content string) error {
	queryString := s.getQueryString(phones, signname, tplcode, content)

	key := []byte(s.accessSecret)
	signstr := fmt.Sprintf("GET&%%2F&%s", s.encodeURL(queryString))

	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(signstr))

	signture := s.encodeURL(secure.EncodeBase64(string(mac.Sum(nil))))
	requesturl := fmt.Sprintf(s.requestURLFormat, signture, queryString)
	logger.I("Send sms, request url:", requesturl)

	resp, err := s.requestRemoteSend(requesturl)
	if err != nil {
		logger.E("Failed request cloud server to send sms")
		return err
	}

	result := &respResult{}
	if err = json.Unmarshal(resp, result); err != nil {
		logger.E("Failed unmarshal send result:", result)
		return err
	}
	logger.I("Cloud server handled resp:", result.Message, result.RequestID, "reslut.Code", result.Code)

	// check send result status
	if result.Message != "OK" {
		logger.E("Failed send sms:", content)
		return errors.New("Failed to send")
	}
	return nil
}

// SendCode send verify sms witch code
func (s *SmsSender) SendCode(sms SmsContent, phones string, code string) error {
	tplcode := sms.TemplateCode
	signName := url.QueryEscape(sms.SignName)
	content := url.QueryEscape(sms.TemplatePrefix + code + sms.TemplateSuffix) // "{\"code\":\"888123\"}"
	if err := s.Send(phones, signName, tplcode, content); err != nil {
		logger.E("Failed send verify sms to:", phones)
		return err
	}
	logger.I("Send verify sms code:", code, "to:", phones)
	return nil
}

// NewSmsSender create a sms sender for given cloud service
func NewSmsSender(secret, keyid, requrl string) *SmsSender {
	sender := &SmsSender{
		accessSecret: secret, accessKeyID: keyid, requestURLFormat: requrl,
	}
	return sender
}
