package live_house_owner_usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_owner_domain"
)

type LiveHouseOwnerRepository interface {
	Save(owner live_house_owner_domain.LiveHouseOwner, ctx context.Context) (*live_house_owner_domain.LiveHouseOwnerId, error)
}

type AccountUseCase struct {
	liveHouseOwnerRepository LiveHouseOwnerRepository
	liveHouseOwner           live_house_owner_domain.LiveHouseOwner
}

func NewAccountUseCase(liveHouseOwnerRepository LiveHouseOwnerRepository) AccountUseCase {
	return AccountUseCase{
		liveHouseOwnerRepository: liveHouseOwnerRepository,
	}
}

func (useCase AccountUseCase) RegisterAccount(name string, emailAddress string, password string, ctx context.Context) (*live_house_owner_domain.LiveHouseOwnerId, error) {

	liveHouseOwnerName, err := live_house_owner_domain.NewLiveHouseOwnerName(name)
	if err != nil {
		return nil, err
	}

	liveHouseOwnerEmailAddress, err := live_house_owner_domain.NewLiveHouseOwnerEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	liveHouseOwnerPassword, err := live_house_owner_domain.NewLiveHouseOwnerPassword(password)
	if err != nil {
		return nil, err
	}

	liveHouseOwner, err := live_house_owner_domain.NewLiveHouseOwner(
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
