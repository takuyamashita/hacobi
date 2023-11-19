package live_house_staff_usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type LiveHouseStaffUsecaseIntf interface {
	RegisterAccount(name string, emailAddress string, password string, ctx context.Context) (string, error)
}

type LiveHouseStaffUsecase struct {
	uuidRepository                    UuidRepositoryIntf
	liveHouseStaffRepository          LiveHouseStaffRepositoryIntf
	liveHouseStaffEmailAddressChecker live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf
}

func NewLiveHouseStaffUsecase(
	uuidRepository UuidRepositoryIntf,
	liveHouseStaffRepository LiveHouseStaffRepositoryIntf,
	liveHouseStaffEmailAddressChecker live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf,
) LiveHouseStaffUsecaseIntf {
	return &LiveHouseStaffUsecase{
		uuidRepository:                    uuidRepository,
		liveHouseStaffRepository:          liveHouseStaffRepository,
		liveHouseStaffEmailAddressChecker: liveHouseStaffEmailAddressChecker,
	}
}
