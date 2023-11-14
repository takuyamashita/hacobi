package livehouseowner

import (
	"context"

	livehouseownerdomain "github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_owner"
)

type LiveHouseOwnerRepository interface {
	Save(owner livehouseownerdomain.LiveHouseOwner, ctx context.Context) (*livehouseownerdomain.LiveHouseOwnerId, error)
}

type AccountUseCase struct {
	liveHouseOwnerRepository LiveHouseOwnerRepository
	liveHouseOwner           livehouseownerdomain.LiveHouseOwner
}

func NewAccountUseCase(liveHouseOwnerRepository LiveHouseOwnerRepository) AccountUseCase {
	return AccountUseCase{
		liveHouseOwnerRepository: liveHouseOwnerRepository,
	}
}

func (useCase AccountUseCase) RegisterAccount(name string, emailAddress string, password string, ctx context.Context) (*livehouseownerdomain.LiveHouseOwnerId, error) {

	liveHouseOwnerName, err := livehouseownerdomain.NewLiveHouseOwnerName(name)
	if err != nil {
		return nil, err
	}

	liveHouseOwnerEmailAddress, err := livehouseownerdomain.NewLiveHouseOwnerEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	liveHouseOwnerPassword, err := livehouseownerdomain.NewLiveHouseOwnerPassword(password)
	if err != nil {
		return nil, err
	}

	liveHouseOwner, err := livehouseownerdomain.NewLiveHouseOwner(
		nil,
		liveHouseOwnerName,
		liveHouseOwnerEmailAddress,
		liveHouseOwnerPassword,
	)
	if err != nil {
		return nil, err
	}

	id, err := useCase.liveHouseOwnerRepository.Save(liveHouseOwner, ctx)

	return id, err
}
