package mail

import (
	"testing"
)

func TestSend(t *testing.T) {
	options := &Options{
		MailHost: "smtp.163.com",
		MailPort: 465,
		MailUser: "xx@163.com",
		MailPass: "xx", //密码或授权码
		MailTo:   "xx",
		Subject:  "subject",
		Body:     "this is body",
	}
	err := Send(options)
	if err != nil {
		t.Error("Mail Send error", err)
		return
	}
	t.Log("success")
}
