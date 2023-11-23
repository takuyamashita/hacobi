package infrastructure

import (
	"fmt"
	"net/smtp"
	"os"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

var _ domain.MailIntf = (*MailImpl)(nil)

type MailImpl struct{}

func NewMail() domain.MailIntf {
	return &MailImpl{}
}

func (m MailImpl) Send(to string, subject string, body string) error {

	from := "testfrom"
	smtpServer := fmt.Sprintf("%s:%d", "mail", 1025)
	auth := smtp.CRAMMD5Auth("username", "password")
	msg := []byte(fmt.Sprintf("To: %s\nSubject: %s\n\n%s", "", subject, body))

	if err := smtp.SendMail(smtpServer, auth, from, []string{to}, msg); err != nil {
		fmt.Fprintln(os.Stderr, err)
	}

	return nil
}
