package live_house_staff_usecase

import (
	"context"
	"errors"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type uuidRepository interface {
	Generate() (string, error)
}

type LiveHouseStaffRepository interface {
	Save(owner live_house_staff_domain.LiveHouseStaff, ctx context.Context) error
	FindByEmail(emailAddress live_house_staff_domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_domain.LiveHouseStaff, error)
}

type AccountUseCase struct {
	liveHouseStaffRepository LiveHouseStaffRepository
	uuidRepository           uuidRepository
	liveHouseStaff           live_house_staff_domain.LiveHouseStaff
}

func NewAccountUseCase(
	uuidRepository uuidRepository,
	liveHouseStaffRepository LiveHouseStaffRepository,
) AccountUseCase {
	return AccountUseCase{
		uuidRepository:           uuidRepository,
		liveHouseStaffRepository: liveHouseStaffRepository,
	}
}

func (useCase AccountUseCase) RegisterAccount(name string, emailAddress string, password string, ctx context.Context) (string, error) {

	liveHouseStaffEmailAddress, err := live_house_staff_domain.NewLiveHouseStaffEmailAddress(emailAddress)
	if err != nil {
		return "", err
	}

	sameEmailAddressStaff, err := useCase.liveHouseStaffRepository.FindByEmail(liveHouseStaffEmailAddress, ctx)
	if err != nil {
		return "", err
	}
	if sameEmailAddressStaff != nil {
		return "", errors.New("email address is already registered")
	}

	id, err := useCase.uuidRepository.Generate()
	if err != nil {
		return "", err
	}

	liveHouseStaffId, err := live_house_staff_domain.NewLiveHouseStaffId(id)
	if err != nil {
		return "", err
	}

	liveHouseStaff, err := live_house_staff_domain.NewliveHouseStaff(
		liveHouseStaffId,
		name,
		emailAddress,
		password,
	)
	if err != nil {
		return "", err
	}

	if err := useCase.liveHouseStaffRepository.Save(liveHouseStaff, ctx); err != nil {
		return "", err
	}

	return liveHouseStaff.Id().String(), nil
}
