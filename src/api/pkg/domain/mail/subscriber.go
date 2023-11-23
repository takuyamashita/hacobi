package mail

import (
	"fmt"
	"os"

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

	e.mail.Send(
		event.EmailAddress.String(),
		"認証メール",
		fmt.Sprintf("%s/live_house_staff_register/?token=%s", os.Getenv("DOMAIN"), event.Token.String()),
	)
}
