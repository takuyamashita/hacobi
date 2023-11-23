package usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/event"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_email_authorization_domain"
)

func SendLiveHouseStaffEmailAuthorization(emailAddress string, ctx context.Context, container container.Container) error {

	var (
		liveHouseStaffEmailAuthorizationRepo LiveHouseStaffEmailAuthorizationRepositoryIntf
		tokenGenerator                       live_house_staff_email_authorization_domain.TokenGeneratorIntf
		eventPublisher                       event.EventPublisherIntf[live_house_staff_email_authorization_domain.AuthCreatedEvent]
	)
	container.Make(&liveHouseStaffEmailAuthorizationRepo)
	container.Make(&tokenGenerator)
	container.Make(&eventPublisher)

	token, err := tokenGenerator.Generate()
	if err != nil {
		return err
	}

	auth, err := live_house_staff_email_authorization_domain.NewLiveHouseStaffEmailAuthorization(emailAddress, token.String())
	if err != nil {
		return err
	}

	liveHouseStaffEmailAuthorizationRepo.Save(auth, ctx)

	event := live_house_staff_email_authorization_domain.AuthCreatedEvent{
		Token:        auth.Token(),
		EmailAddress: auth.EmailAddress(),
	}

	eventPublisher.Publish(event)

	return nil
}
