package live_house_staff_account_domain

import "github.com/takuyamashita/hacobi/src/api/pkg/domain"

type ProvisionalLiveHouseAccountCreated struct {
	Token        Token
	EmailAddress domain.LiveHouseStaffEmailAddress
}
