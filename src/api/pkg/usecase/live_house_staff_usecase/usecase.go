package live_house_staff_usecase

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type LiveHouseStaffUsecase struct {
	uuidRepository                    UuidRepository
	liveHouseStaffRepository          LiveHouseStaffRepository
	liveHouseStaffEmailAddressChecker live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf
}

func NewLiveHouseStaffUsecase(
	uuidRepository UuidRepository,
	liveHouseStaffRepository LiveHouseStaffRepository,
	liveHouseStaffEmailAddressChecker live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf,
) LiveHouseStaffUsecase {
	return LiveHouseStaffUsecase{
		uuidRepository:                    uuidRepository,
		liveHouseStaffRepository:          liveHouseStaffRepository,
		liveHouseStaffEmailAddressChecker: liveHouseStaffEmailAddressChecker,
	}
}
