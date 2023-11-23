package usecase

import (
	"context"
	"errors"

	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

func RegisterLiveHouseStaff(name string, emailAddress string, password string, ctx context.Context, container container.Container) (string, error) {

	var (
		uuidRepo             UuidRepositoryIntf
		txRepo               TransationRepositoryIntf
		liveHouseStaffRepo   LiveHouseStaffRepositoryIntf
		liveHouseAccountRepo LiveHouseAccountRepositoryIntf
		emailChecker         live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf
	)
	container.Make(&uuidRepo)
	container.Make(&txRepo)
	container.Make(&liveHouseStaffRepo)
	container.Make(&liveHouseAccountRepo)
	container.Make(&emailChecker)

	commit, rollback, err := txRepo.Begin(ctx)
	defer commit()
	if err != nil {
		return "", err
	}

	// ???????? lock ?????????
	isEmailAddressAlreadyRegistered, err := emailChecker.IsEmailAddressAlreadyRegistered(emailAddress, ctx)
	if err != nil {
		return "", err
	}
	if isEmailAddressAlreadyRegistered {
		return "", errors.New("email address is already registered")
	}

	id, err := uuidRepo.Generate()
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

	if err := liveHouseStaffRepo.Save(liveHouseStaff, ctx); err != nil {
		rollback()
		return "", err
	}

	return liveHouseStaff.Id().String(), nil
}
