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
	"crypto/hmac"
	"crypto/sha1"
	"encoding/json"
	"fmt"
	"github.com/satori/go.uuid"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
	"wing/logger"
	"wing/utils"
	"wing/utils/crypto"
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

// GetResult get method response
type GetResult struct {
	Message   string `json:"Message"`
	RequestId string `json:"RequestId"`
	BizId     string "BizId"
	Code      string "Code"
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
	signnonce, _ := uuid.NewV4()
	timestamp := url.QueryEscape(time.Now().UTC().Format(time_layout))
	return fmt.Sprintf(sort_query_format,
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
	signstr := fmt.Sprintf("GET&%%2F&%s", s.encodeUrl(queryString))

	mac := hmac.New(sha1.New, key)
	mac.Write([]byte(signstr))

	signture := s.encodeUrl(crypto.ToBase64String(mac.Sum(nil)))
	requesturl := fmt.Sprintf(s.requestUrlFormat, signture, queryString)
	logger.I("Send sms, request url:", requesturl)

	resp, err := s.execHttpGet(requesturl)
	if err != nil {
		logger.E("Failed request cloud server to send sms")
		return err
	}

	result := &GetResult{}
	if err = json.Unmarshal(resp, result); err != nil {
		logger.E("Failed unmarshal send result:", result)
		return err
	}
	logger.I("Cloud server handled resp:", result.Message, result.RequestId)

	// check send result status
	if result.Message != "OK" {
		logger.E("Failed send sms:", content)
		return utils.ErrSendFailed
	}
	return nil
}

// NewSmsSender create a sms sender for given cloud service
func NewSmsSender(secret, keyid, requrl string) *SmsSender {
	sender := &SmsSender{
		accessSecret: secret, accessKeyID: keyid, requestUrlFormat: requrl,
	}
	return sender
}
