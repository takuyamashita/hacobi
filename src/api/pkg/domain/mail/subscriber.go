package mail

import (
	"log"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_email_authorization_domain"
)

type LiveHouseStaffEmailAuthorization struct {
	mail domain.MailIntf
}

func (e LiveHouseStaffEmailAuthorization) Handle(event live_house_staff_email_authorization_domain.AuthCreatedEvent) {

	log.Println("send email", event)
}
