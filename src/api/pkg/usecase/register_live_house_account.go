package usecase

import (
	"context"
	"errors"

	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

func RegisterLiveHouseAccount(userId string, ctx context.Context, container container.Container) (string, error) {

	var (
		uuidRepo             UuidRepositoryIntf
		txRepo               TransationRepositoryIntf
		liveHouseAccountRepo LiveHouseAccountRepositoryIntf
		liveHouseStaffRepo   LiveHouseStaffRepositoryIntf
	)
	container.Make(&uuidRepo)
	container.Make(&txRepo)
	container.Make(&liveHouseAccountRepo)
	container.Make(&liveHouseStaffRepo)

	accountId, err := uuidRepo.Generate()
	if err != nil {
		return "", err
	}

	liveHouseStaffId, err := live_house_staff_domain.NewLiveHouseStaffId(userId)
	if err != nil {
		return "", err
	}

	commit, rollback, err := txRepo.Begin(ctx)
	defer commit()
	if err != nil {
		return "", err
	}

	{

		// ???????? lock ?????????
		liveHouseStaff, err := liveHouseStaffRepo.FindById(liveHouseStaffId, ctx)
		if err != nil {
			return "", err
		}
		if liveHouseStaff == nil {
			return "", errors.New("live house staff not found")
		}

		account, err := live_house_account_domain.NewLiveHouseAccount(
			accountId,
			[]live_house_account_domain.StaffParams{
				{
					Id:   liveHouseStaffId,
					Role: live_house_account_domain.GetRoleMaster(),
				},
			},
			nil,
		)
		if err != nil {
			return "", err
		}

		err = liveHouseAccountRepo.Save(account, ctx)
		if err != nil {
			rollback()
			return "", err
		}

		return account.Id().String(), nil
	}
}
