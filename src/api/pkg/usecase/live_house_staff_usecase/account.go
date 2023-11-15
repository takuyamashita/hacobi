package live_house_staff_usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type uuidRepository interface {
	Generate() (string, error)
}

type liveHouseStaffRepository interface {
	Save(owner live_house_staff_domain.LiveHouseStaff, ctx context.Context) error
}

type AccountUseCase struct {
	liveHouseStaffRepository liveHouseStaffRepository
	uuidRepository           uuidRepository
	liveHouseStaff           live_house_staff_domain.LiveHouseStaff
}

func NewAccountUseCase(
	uuidRepository uuidRepository,
	liveHouseStaffRepository liveHouseStaffRepository,
) AccountUseCase {
	return AccountUseCase{
		uuidRepository:           uuidRepository,
		liveHouseStaffRepository: liveHouseStaffRepository,
	}
}

func (useCase AccountUseCase) RegisterAccount(name string, emailAddress string, password string, ctx context.Context) (live_house_staff_domain.LiveHouseStaffId, error) {

	liveHouseStaffName, err := live_house_staff_domain.NewliveHouseStaffName(name)
	if err != nil {
		return live_house_staff_domain.LiveHouseStaffId{}, err
	}

	liveHouseStaffEmailAddress, err := live_house_staff_domain.NewLiveHouseStaffEmailAddress(emailAddress)
	if err != nil {
		return live_house_staff_domain.LiveHouseStaffId{}, err
	}

	liveHouseStaffPassword, err := live_house_staff_domain.NewliveHouseStaffPassword(password)
	if err != nil {
		return live_house_staff_domain.LiveHouseStaffId{}, err
	}

	id, err := useCase.uuidRepository.Generate()
	if err != nil {
		return live_house_staff_domain.LiveHouseStaffId{}, err
	}

	liveHouseStaffId, err := live_house_staff_domain.NewliveHouseStaffId(id)
	if err != nil {
		return live_house_staff_domain.LiveHouseStaffId{}, err
	}

	liveHouseStaff, err := live_house_staff_domain.NewliveHouseStaff(
		liveHouseStaffId,
		liveHouseStaffName,
		liveHouseStaffEmailAddress,
		liveHouseStaffPassword,
	)
	if err != nil {
		return live_house_staff_domain.LiveHouseStaffId{}, err
	}

	return liveHouseStaffId, useCase.liveHouseStaffRepository.Save(liveHouseStaff, ctx)
}
