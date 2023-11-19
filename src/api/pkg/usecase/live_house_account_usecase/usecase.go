package live_house_account_usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type LiveHouseAccountUsecaseIntf interface {
	RegisterLiveHouseAccount(userId string, ctx context.Context) (string, error)
}

type LiveHouseAccountUsecase struct {
	uuidRepository             usecase.UuidRepositoryIntf
	liveHouseStaffRepository   usecase.LiveHouseStaffRepositoryIntf
	liveHouseAccountRepository usecase.LiveHouseAccountRepositoryIntf
	transationRepository       usecase.TransationRepositoryIntf
}

func NewLiveHouseAccountUsecase(
	uuidRepository usecase.UuidRepositoryIntf,
	liveHouseStaffRepository usecase.LiveHouseStaffRepositoryIntf,
	liveHouseAccountRepository usecase.LiveHouseAccountRepositoryIntf,
	transationRepository usecase.TransationRepositoryIntf,
) LiveHouseAccountUsecaseIntf {
	return &LiveHouseAccountUsecase{
		uuidRepository:             uuidRepository,
		liveHouseStaffRepository:   liveHouseStaffRepository,
		liveHouseAccountRepository: liveHouseAccountRepository,
		transationRepository:       transationRepository,
	}
}
