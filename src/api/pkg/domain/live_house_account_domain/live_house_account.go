package live_house_account_domain

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type liveHouseAccount struct {
	id        *LiveHouseAccountId
	name      LiveHouseAccountName
	staffs    []live_house_staff_domain.LiveHouseStaffId
	liveHouse *live_house_domain.LiveHouseId
}
