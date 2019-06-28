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
	"gopkg.in/gomail.v2"
)

/*
 * Useage
 *
 * [CODE:]
 * mailTempleteSubject = "You hava a mail！"
 * mailTemplateFormat  = `
 *	<p>
 *		<span font-weight:bold; style="font-size:16px; color:#363636">Dear</span><br><br>
 *		<span style="font-size:14px; color:#484848">Your account %s have not activate, please click the follow link to activate it.</span><br>
 *		<span style="font-size:14px; color:#484848">%s</span><br>
 *	</p>
 *	<p align="right">
 *		<span style="font-size:12px; color:#484848">From %s</span><br>
 *		<span style="font-size:10px; color:#636363">%s</span>
 *	</p>
 *	`
 *
 * mailagent = &comm.MailAgent{
 *     Acc: account, Pwd: password, Host: smtphost, Port: smtpport,
 * }
 * subject := mailTempleteSubject
 * message := fmt.Sprintf(mailTemplateFormat, account, link, who,
 *     time.Now().Format(templateTimeFormat))
 * // send mail with attachment
 * // return ma.SendMail(to, subject, message, fileName)
 * return ma.SendMail(to, subject, message)
 * [CODE]
 */

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
