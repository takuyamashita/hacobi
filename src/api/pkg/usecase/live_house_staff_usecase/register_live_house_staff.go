package live_house_staff_usecase

import (
	"context"
	"errors"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

func (usecase LiveHouseStaffUsecase) RegisterAccount(name string, emailAddress string, password string, ctx context.Context) (string, error) {

	//xxx: transactionを貼る

	isEmailAddressAlreadyRegistered, err := usecase.liveHouseStaffEmailAddressChecker.IsEmailAddressAlreadyRegistered(emailAddress, ctx)
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

	liveHouseStaff, err := live_house_staff_domain.NewLiveHouseStaff(
		id,
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
