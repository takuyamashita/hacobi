package usecase

import (
	"context"
	"errors"
	"fmt"
	"io"
	"log"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/webauthn"
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/event"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
)

// 仮アカウントの登録と本登録メールの送信をおこないます
func RegisterProvisionalLiveHouseStaffAccount(emailAddress string, ctx context.Context, container container.Container) error {

	var (
		uuidRepo                  UuidRepositoryIntf
		txRepo                    TransationRepositoryIntf
		liveHouseStaffAccountRepo LiveHouseStaffAccountRepositoryIntf
		emailAddressChecker       live_house_staff_account_domain.LiveHouseStaffAccountEmailAddressCheckerIntf
		tokenGenerator            live_house_staff_account_domain.TokenGeneratorIntf
		eventPublisher            event.EventPublisherIntf[live_house_staff_account_domain.ProvisionalLiveHouseAccountCreated]
	)
	container.Make(&uuidRepo)
	container.Make(&txRepo)
	container.Make(&liveHouseStaffAccountRepo)
	container.Make(&emailAddressChecker)
	container.Make(&tokenGenerator)
	container.Make(&eventPublisher)

	commit, rollback, err := txRepo.Begin(ctx)
	defer rollback()
	if err != nil {
		return err
	}

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
		rollback()
		return err
	}

	event := live_house_staff_account_domain.ProvisionalLiveHouseAccountCreated{
		Token:        *account.ProvisionalToken(),
		EmailAddress: account.EmailAddress(),
	}

	eventPublisher.Publish(event)

	return commit()
}

// 本登録を開始します
// 本登録ではWebAuthnを使用したアカウント登録をおこなうため、そのために必要な情報を返します
func StartRegister(
	token string,
	ctx context.Context,
	container container.Container,
) (*protocol.PublicKeyCredentialCreationOptions, string, error) {

	var (
		publicKeyCredential       CredentialKeyIntf
		liveHouseStaffAccountRepo LiveHouseStaffAccountRepositoryIntf
	)
	container.Make(&publicKeyCredential)
	container.Make(&liveHouseStaffAccountRepo)

	tkn, err := live_house_staff_account_domain.NewTokenFromHexString(token)
	if err != nil {
		return nil, "", err
	}

	account, err := liveHouseStaffAccountRepo.FindByProvisionalRegistrationToken(tkn, ctx)
	if err != nil {
		return nil, "", err
	}
	if account == nil {
		return nil, "", errors.New("account not found")
	}

	// xxx: repositoryChallengeとdomainChallengeが存在しており混乱を招く
	repositoryChallenge, err := publicKeyCredential.CreateChallenge()
	if err != nil {
		return nil, "", err
	}

	domainChallenge, err := live_house_staff_account_domain.NewCredentialChallenge(
		repositoryChallenge.String(),
		time.Now(),
	)
	if err != nil {
		return nil, "", err
	}

	if err := account.AddCredentialChallenge(domainChallenge); err != nil {
		return nil, "", err
	}

	if err := liveHouseStaffAccountRepo.Save(account, ctx); err != nil {
		return nil, "", err
	}

	option := publicKeyCredential.CreateCredentialCreationOptions(repositoryChallenge, "localhost")

	return &option, account.Id().String(), nil

}

func FinishRegisterLiveHouseStaffAccount(
	reader io.Reader,
	accountId string,
	ctx context.Context,
	container container.Container,
) error {

	var (
		uuidRepo                  UuidRepositoryIntf
		txRepo                    TransationRepositoryIntf
		liveHouseStaffAccountRepo LiveHouseStaffAccountRepositoryIntf
		publicKeyCredential       CredentialKeyIntf
	)
	container.Make(&uuidRepo)
	container.Make(&txRepo)
	container.Make(&liveHouseStaffAccountRepo)
	container.Make(&publicKeyCredential)

	commit, rollback, err := txRepo.Begin(ctx)
	defer rollback()
	if err != nil {
		return err
	}

	parsedResponse, err := protocol.ParseCredentialCreationResponseBody(reader)
	if err != nil {

		// cast error to protocol.Error
		fmt.Println(err.(*protocol.Error).Details)

		log.Println(err)
		return err
	}

	liveHouseStaffAccountId := live_house_staff_account_domain.NewLiveHouseStaffAccountId(accountId)
	account, err := liveHouseStaffAccountRepo.FindById(liveHouseStaffAccountId, ctx)
	if err != nil {
		return err
	}

	for _, v := range account.CredentialChallenges() {
		log.Println(v.Challenge().String())
	}

	if err := parsedResponse.Verify(
		account.CredentialChallenges()[len(account.CredentialChallenges())-1].Challenge().String(),
		true,
		"localhost",
		[]string{"localhost"},
	); err != nil {
		return err
	}

	credential, err := webauthn.MakeNewCredential(parsedResponse)
	if err != nil {
		return err
	}

	log.Println(credential)

	return commit()
}
