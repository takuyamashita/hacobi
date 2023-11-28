package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/event"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
)

func RegisterProvisionalLiveHouseAccount(emailAddress string, ctx context.Context, container container.Container) error {

	var (
		uuidRepo                  UuidRepositoryIntf
		liveHouseStaffAccountRepo LiveHouseStaffAccountRepositoryIntf
		emailAddressChecker       live_house_staff_account_domain.LiveHouseStaffAccountEmailAddressCheckerIntf
		tokenGenerator            live_house_staff_account_domain.TokenGeneratorIntf
		eventPublisher            event.EventPublisherIntf[live_house_staff_account_domain.ProvisionalLiveHouseAccountCreated]
	)
	container.Make(&uuidRepo)
	container.Make(&liveHouseStaffAccountRepo)
	container.Make(&emailAddressChecker)
	container.Make(&tokenGenerator)
	container.Make(&eventPublisher)

	isAlreadyRegistered, err := emailAddressChecker.IsEmailAddressAlreadyRegistered(emailAddress, ctx)
	if err != nil {
		return err
	}
	if isAlreadyRegistered {
		return errors.New("email address is already registered")
	}

	token, err := tokenGenerator.Generate()
	if err != nil {
		return err
	}

	id, err := uuidRepo.Generate()
	if err != nil {
		return err
	}

	account, err := live_house_staff_account_domain.NewLiveHouseStaffAccount(live_house_staff_account_domain.NewLiveHouseStaffAccountParams{
		Id:            id,
		Email:         emailAddress,
		IsProvisional: true,
		ProvisionalRegistration: &live_house_staff_account_domain.ProvisionalRegistrationParam{
			Token:     token.String(),
			CreatedAt: time.Now(),
		},
	})

	if err := liveHouseStaffAccountRepo.Save(account, ctx); err != nil {
		return err
	}

	event := live_house_staff_account_domain.ProvisionalLiveHouseAccountCreated{
		Token:        *account.ProvisionalToken(),
		EmailAddress: account.EmailAddress(),
	}

	eventPublisher.Publish(event)

	return nil
}
