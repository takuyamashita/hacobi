package live_house_account_usecase

import (
	"context"
	"errors"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

func (u LiveHouseAccountUsecase) RegisterLiveHouseAccount(userId string, ctx context.Context) (string, error) {

	accountId, err := u.uuidRepository.Generate()
	if err != nil {
		return "", err
	}

	liveHouseStaffId, err := live_house_staff_domain.NewLiveHouseStaffId(userId)
	if err != nil {
		return "", err
	}

	commit, rollback, err := u.transationRepository.Begin(ctx)
	defer commit()
	if err != nil {
		return "", err
	}

	{

		// ???????? lock ?????????
		liveHouseStaff, err := u.liveHouseStaffRepository.FindById(liveHouseStaffId, ctx)
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

		err = u.liveHouseAccountRepository.Save(account, ctx)
		if err != nil {
			rollback()
			return "", err
		}

		return account.Id().String(), nil
	}
}
