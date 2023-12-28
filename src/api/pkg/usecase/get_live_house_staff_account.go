package usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
)

func GetLiveHouseStaffAccount(
	accountID string,
	ctx context.Context,
	container container.Container,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	var (
		liveHouseStaffAccountRepo LiveHouseStaffAccountRepositoryIntf
	)

	container.Make(&liveHouseStaffAccountRepo)

	id := live_house_staff_account_domain.NewLiveHouseStaffAccountId(accountID)

	account, err := liveHouseStaffAccountRepo.FindById(id, ctx)

	if err != nil {
		return nil, err
	}

	return account, nil

}
