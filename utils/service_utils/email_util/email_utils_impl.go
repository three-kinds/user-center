package email_util

import (
	"bytes"
	_ "embed"
	"fmt"
	"github.com/three-kinds/user-center/initializers"
	"github.com/three-kinds/user-center/utils/service_utils/se"
	"gopkg.in/gomail.v2"
	"html/template"
)

type lEmail struct {
	ToList     []string
	CcList     []string
	Subject    string
	Content    string
	AttachList []string
}

//go:embed templates/reset_password_email.html
var resetPasswordEmailTemplate string

type EmailUtilImpl struct {
}

func (u *EmailUtilImpl) sendEmail(email *lEmail) error {
	config := initializers.Config

	m := gomail.NewMessage()
	m.SetAddressHeader("From", config.EmailUsername, config.EmailFrom)
	m.SetHeader("To", email.ToList...)
	m.SetHeader("Cc", email.CcList...)
	m.SetHeader("Subject", email.Subject)
	m.SetBody("text/html", email.Content)
	for _, attach := range email.AttachList {
		m.Attach(attach)
	}

	d := gomail.NewDialer(config.EmailHost, config.EmailPort, config.EmailUsername, config.EmailPassword)
	d.SSL = config.EmailUseSSL
	err := d.DialAndSend(m)
	if err != nil {
		return se.ServerKnownError(fmt.Sprintf("send email failed: %s", err.Error()))
	}
	return nil

}

func (u *EmailUtilImpl) SendResetPasswordEmail(email string, codeKey string) error {
	tmpl := template.Must(template.New("resetPasswordEmailTemplate").Parse(resetPasswordEmailTemplate))
	content := bytes.Buffer{}

	err := tmpl.Execute(&content, codeKey)
	if err != nil {
		return se.ServerKnownError(fmt.Sprintf("resetPasswordEmailTemplate render failed: %s", err))
	}

	e := &lEmail{}
	e.ToList = []string{email}
	e.Subject = "重置密码"
	e.Content = content.String()

	go func() {
		_ = u.sendEmail(e)
	}()

	return nil
}

func NewEmailUtilImpl() *EmailUtilImpl {
	return &EmailUtilImpl{}
}
