package live_house_staff_domain

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

type LiveHouseStaffEmailAddressCheckerIntf interface {
	IsEmailAddressAlreadyRegistered(emailAddress string, ctx context.Context) (bool, error)
}

type liveHouseStaffEmailAddressCheckerImpl struct {
	liveHouseStaffRepository LiveHouseStaffRepositoryIntf
}

func NewLiveHouseStaffEmailAddressChecker(
	repository LiveHouseStaffRepositoryIntf,
) LiveHouseStaffEmailAddressCheckerIntf {
	return &liveHouseStaffEmailAddressCheckerImpl{
		liveHouseStaffRepository: repository,
	}
}

func (checker liveHouseStaffEmailAddressCheckerImpl) IsEmailAddressAlreadyRegistered(emailAddress string, ctx context.Context) (bool, error) {

	liveHouseStaffEmailAddress, err := domain.NewLiveHouseStaffEmailAddress(emailAddress)

	sameEmailAddressStaff, err := checker.liveHouseStaffRepository.FindByEmail(liveHouseStaffEmailAddress, ctx)
	if err != nil {
		return false, err
	}
	if sameEmailAddressStaff != nil {
		return true, nil
	}

	return false, nil
}
