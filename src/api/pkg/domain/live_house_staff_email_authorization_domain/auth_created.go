package live_house_staff_email_authorization_domain

import "github.com/takuyamashita/hacobi/src/api/pkg/domain"

type AuthCreatedEvent struct {
	Token        Token
	EmailAddress domain.LiveHouseStaffEmailAddress
}
