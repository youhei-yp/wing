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
	"gopkg.in/gomail.v2"
)

// MailAgent mail agent informations
type MailAgent struct {
	Acc  string `json:"acc"`  // username - mail address
	Pwd  string `json:"pwd"`  // account password
	Host string `json:"host"` // stmp/pop3 host
	Port int    `json:"port"` // stmp/pop3 port
}

// SendMail send email by mail account, it may set attachment from local file
func (a *MailAgent) SendMail(to []string, subject, body string, attach ...string) error {
	m := gomail.NewMessage()
	m.SetHeader("From", a.Acc)
	m.SetHeader("To", to[0])
	m.SetHeader("Subject", subject)
	m.SetBody("text/html", body)
	if len(attach) > 0 && attach[0] != "" {
		m.Attach(attach[0])
	}

	d := gomail.NewPlainDialer(a.Host, a.Port, a.Acc, a.Pwd)
	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
