package live_house_staff_usecase

import (
	"context"
	"errors"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type UuidRepository interface {
	Generate() (string, error)
}

type LiveHouseStaffRepository interface {
	Save(owner live_house_staff_domain.LiveHouseStaff, ctx context.Context) error
	FindByEmail(emailAddress live_house_staff_domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_domain.LiveHouseStaff, error)
}

type AccountUsecase struct {
	liveHouseStaffRepository LiveHouseStaffRepository
	uuidRepository           UuidRepository
	liveHouseStaff           live_house_staff_domain.LiveHouseStaff
}

func NewAccountUsecase(
	uuidRepository UuidRepository,
	liveHouseStaffRepository LiveHouseStaffRepository,
) AccountUsecase {
	return AccountUsecase{
		uuidRepository:           uuidRepository,
		liveHouseStaffRepository: liveHouseStaffRepository,
	}
}

func (usecase AccountUsecase) RegisterAccount(name string, emailAddress string, password string, ctx context.Context) (string, error) {

	liveHouseStaffEmailAddress, err := live_house_staff_domain.NewLiveHouseStaffEmailAddress(emailAddress)
	if err != nil {
		return "", err
	}

	isEmailAddressAlreadyRegistered, err := usecase.isEmailAddressAlreadyRegistered(liveHouseStaffEmailAddress, ctx)
	if err != nil {
		return "", err
	}
	if isEmailAddressAlreadyRegistered {
		return "", errors.New("email address is already registered")
	}

	id, err := usecase.uuidRepository.Generate()
	if err != nil {
		return "", err
	}

	liveHouseStaffId, err := live_house_staff_domain.NewLiveHouseStaffId(id)
	if err != nil {
		return "", err
	}

	liveHouseStaff, err := live_house_staff_domain.NewLiveHouseStaff(
		liveHouseStaffId,
		name,
		emailAddress,
		password,
	)
	if err != nil {
		return "", err
	}

	if err := usecase.liveHouseStaffRepository.Save(liveHouseStaff, ctx); err != nil {
		return "", err
	}

	return liveHouseStaff.Id().String(), nil
}

func (useCase AccountUsecase) isEmailAddressAlreadyRegistered(emailAddress live_house_staff_domain.LiveHouseStaffEmailAddress, ctx context.Context) (bool, error) {

	sameEmailAddressStaff, err := useCase.liveHouseStaffRepository.FindByEmail(emailAddress, ctx)
	if err != nil {
		return false, err
	}
	if sameEmailAddressStaff != nil {
		return true, nil
	}

	return false, nil
}
