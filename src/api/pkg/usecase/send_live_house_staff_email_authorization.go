package usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/event"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
)

func SendLiveHouseStaffEmailAuthorization(emailAddress string, ctx context.Context, container container.Container) error {

	var (
		liveHouseStaffAccountRepo LiveHouseStaffAccountRepositoryIntf
		tokenGenerator            live_house_staff_account_domain.TokenGeneratorIntf
		eventPublisher            event.EventPublisherIntf[live_house_staff_account_domain.AuthCreatedEvent]
	)
	container.Make(&liveHouseStaffAccountRepo)
	container.Make(&tokenGenerator)
	container.Make(&eventPublisher)

	token, err := tokenGenerator.Generate()
	if err != nil {
		return err
	}

	registration, err := live_house_staff_account_domain.NewLiveHouseStaffAccountProvisionalRegistration(token.String())
	if err != nil {
		return err
	}

	event := live_house_staff_account_domain.AuthCreatedEvent{
		Token: registration.Token(),
	}

	eventPublisher.Publish(event)

	return nil
}
