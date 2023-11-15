package live_house_staff_usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type UuidRepository interface {
	Generate() (string, error)
}

type LiveHouseStaffRepository interface {
	Save(owner live_house_staff_domain.LiveHouseStaff, ctx context.Context) error
	FindByEmail(emailAddress live_house_staff_domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_domain.LiveHouseStaff, error)
}

type LiveHouseStaffUsecase struct {
	liveHouseStaffRepository LiveHouseStaffRepository
	uuidRepository           UuidRepository
	liveHouseStaff           live_house_staff_domain.LiveHouseStaff
}

func NewLiveHouseStaffUsecase(
	uuidRepository UuidRepository,
	liveHouseStaffRepository LiveHouseStaffRepository,
) LiveHouseStaffUsecase {
	return LiveHouseStaffUsecase{
		uuidRepository:           uuidRepository,
		liveHouseStaffRepository: liveHouseStaffRepository,
	}
}
