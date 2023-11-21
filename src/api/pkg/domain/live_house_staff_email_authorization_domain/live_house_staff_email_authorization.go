package live_house_staff_email_authorization_domain

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

type LiveHouseStaffEmailAuthorizationIntf interface {
}

type liveHouseStaffEmailAuthorizationImpl struct {
	emailAddress domain.LiveHouseStaffEmailAddress
	token        Token
}

func NewLiveHouseStaffEmailAuthorization(emailAddress string, token string) (LiveHouseStaffEmailAuthorizationIntf, error) {

	liveHouseStaffEmailAddress, err := domain.NewLiveHouseStaffEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	return &liveHouseStaffEmailAuthorizationImpl{
		emailAddress: liveHouseStaffEmailAddress,
	}, nil
}
