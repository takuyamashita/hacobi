package live_house_staff_account_domain

import (
	"context"
	"log"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

type LiveHouseStaffAccountEmailAddressCheckerIntf interface {
	IsEmailAddressAlreadyRegistered(emailAddress string, ctx context.Context) (bool, error)
}

type liveHouseStaffAccountEmailAddressCheckerImpl struct {
	liveHouseStaffAccountRepository LiveHouseStaffAccountRepositoryIntf
}

func NewLiveHouseStaffAccountEmailAddressChecker(
	repository LiveHouseStaffAccountRepositoryIntf,
) LiveHouseStaffAccountEmailAddressCheckerIntf {
	return &liveHouseStaffAccountEmailAddressCheckerImpl{
		liveHouseStaffAccountRepository: repository,
	}
}

func (checker liveHouseStaffAccountEmailAddressCheckerImpl) IsEmailAddressAlreadyRegistered(emailAddress string, ctx context.Context) (bool, error) {

	liveHouseStaffEmailAddress, err := domain.NewLiveHouseStaffEmailAddress(emailAddress)

	sameEmailAddressStaff, err := checker.liveHouseStaffAccountRepository.FindByEmail(liveHouseStaffEmailAddress, ctx)
	log.Println(sameEmailAddressStaff)
	if err != nil {
		return false, err
	}
	if sameEmailAddressStaff != nil {
		return true, nil
	}

	return false, nil
}
