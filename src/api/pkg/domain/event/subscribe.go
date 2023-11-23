package event

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_email_authorization_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/mail"
)

func NewLiveHouseStaffEmailAuthorizationCreated(m domain.MailIntf) EventPublisherIntf[live_house_staff_email_authorization_domain.AuthCreatedEvent] {

	// When AuthCreatedEvent
	p := NewEventPublisher[live_house_staff_email_authorization_domain.AuthCreatedEvent]()

	// send email
	p.Subscribe(mail.NewLiveHouseStaffEmailAuthorization(m))

	return p
}
