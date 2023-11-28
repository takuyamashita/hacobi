package event

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/mail"
)

func NewLiveHouseStaffAccountProvisionalRegistrationCreated(m domain.MailIntf) EventPublisherIntf[live_house_staff_account_domain.ProvisionalLiveHouseAccountCreated] {

	// When ProvisionalLiveHouseAccountCreated
	p := NewEventPublisher[live_house_staff_account_domain.ProvisionalLiveHouseAccountCreated]()

	// send email
	p.Subscribe(mail.NewLiveHouseStaffAccountProvisionalRegistration(m))

	return p
}
