package infrastructure

import "github.com/takuyamashita/hacobi/src/api/pkg/domain"

var _ domain.MailIntf = (*MailImpl)(nil)

type MailImpl struct{}

func (m MailImpl) Send(to string, subject string, body string) error {
	return nil
}
