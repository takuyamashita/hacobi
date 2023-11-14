package live_house_staff_usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type liveHouseStaffRepository interface {
	Save(owner live_house_staff_domain.LiveHouseStaff, ctx context.Context) (*live_house_staff_domain.LiveHouseStaffId, error)
}

type AccountUseCase struct {
	liveHouseStaffRepository liveHouseStaffRepository
	liveHouseStaff           live_house_staff_domain.LiveHouseStaff
}

func NewAccountUseCase(liveHouseStaffRepository liveHouseStaffRepository) AccountUseCase {
	return AccountUseCase{
		liveHouseStaffRepository: liveHouseStaffRepository,
	}
}

func (useCase AccountUseCase) RegisterAccount(name string, emailAddress string, password string, ctx context.Context) (*live_house_staff_domain.LiveHouseStaffId, error) {

	liveHouseStaffName, err := live_house_staff_domain.NewliveHouseStaffName(name)
	if err != nil {
		return nil, err
	}

	liveHouseStaffEmailAddress, err := live_house_staff_domain.NewLiveHouseStaffEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	liveHouseStaffPassword, err := live_house_staff_domain.NewliveHouseStaffPassword(password)
	if err != nil {
		return nil, err
	}

	liveHouseStaff, err := live_house_staff_domain.NewliveHouseStaff(
		nil,
		liveHouseStaffName,
		liveHouseStaffEmailAddress,
		liveHouseStaffPassword,
	)
	if err != nil {
		return nil, err
	}

	id, err := useCase.liveHouseStaffRepository.Save(liveHouseStaff, ctx)

	return id, err
}
