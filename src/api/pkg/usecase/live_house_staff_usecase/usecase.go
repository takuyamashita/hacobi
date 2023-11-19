package live_house_staff_usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type LiveHouseStaffUsecaseIntf interface {
	RegisterAccount(name string, emailAddress string, password string, ctx context.Context) (string, error)
}

type LiveHouseStaffUsecase struct {
	uuidRepository                    usecase.UuidRepositoryIntf
	transactionRepository             usecase.TransationRepositoryIntf
	liveHouseStaffRepository          usecase.LiveHouseStaffRepositoryIntf
	liveHouseStaffEmailAddressChecker live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf
}

func NewLiveHouseStaffUsecase(
	uuidRepository usecase.UuidRepositoryIntf,
	transactionRepository usecase.TransationRepositoryIntf,
	liveHouseStaffRepository usecase.LiveHouseStaffRepositoryIntf,
	liveHouseStaffEmailAddressChecker live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf,
) LiveHouseStaffUsecaseIntf {
	return &LiveHouseStaffUsecase{
		transactionRepository:             transactionRepository,
		uuidRepository:                    uuidRepository,
		liveHouseStaffRepository:          liveHouseStaffRepository,
		liveHouseStaffEmailAddressChecker: liveHouseStaffEmailAddressChecker,
	}
}
