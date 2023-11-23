package mail

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_email_authorization_domain"
)

type LiveHouseStaffEmailAuthorization struct {
	mail domain.MailIntf
}

func NewLiveHouseStaffEmailAuthorization(m domain.MailIntf) LiveHouseStaffEmailAuthorization {
	return LiveHouseStaffEmailAuthorization{
		mail: m,
	}
}

func (e LiveHouseStaffEmailAuthorization) Handle(event live_house_staff_email_authorization_domain.AuthCreatedEvent) {

	e.mail.Send(event.EmailAddress.String(), "認証メール", "認証メールの本文")
}
