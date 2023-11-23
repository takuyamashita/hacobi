package live_house_staff_email_authorization_domain

import "github.com/takuyamashita/hacobi/src/api/pkg/domain"

type AuthCreatedEvent struct {
	To           Token
	EmailAddress domain.LiveHouseStaffEmailAddress
}
