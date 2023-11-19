package live_house_account_usecase

import "github.com/takuyamashita/hacobi/src/api/pkg/usecase"

type LiveHouseAccountUsecaseIntf interface {
}

type LiveHouseAccountUsecase struct {
	uuidRepository           usecase.UuidRepositoryIntf
	liveHouseStaffRepository usecase.LiveHouseStaffRepositoryIntf
}
