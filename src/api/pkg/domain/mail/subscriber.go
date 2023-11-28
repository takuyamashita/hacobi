package mail

import (
	"fmt"
	"os"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
)

type LiveHouseStaffAccountProvisionalRegistration struct {
	mail domain.MailIntf
}

func NewLiveHouseStaffAccountProvisionalRegistration(m domain.MailIntf) LiveHouseStaffAccountProvisionalRegistration {
	return LiveHouseStaffAccountProvisionalRegistration{
		mail: m,
	}
}

func (e LiveHouseStaffAccountProvisionalRegistration) Handle(event live_house_staff_account_domain.AuthCreatedEvent) {

	e.mail.Send(
		event.EmailAddress.String(),
		"認証メール",
		fmt.Sprintf("%s://%s/livehouse/register/staff/%s", os.Getenv("URL_PROTOCOL"), os.Getenv("FQDN"), event.Token.String()),
	)
}
