package usecase

import (
	"context"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
)

func StartLiveHouseStaffAccountLogin(
	emailAddress string,
	ctx context.Context,
	container container.Container,
) (*protocol.PublicKeyCredentialRequestOptions, error) {

	var (
		liveHouseStaffAccountRepo LiveHouseStaffAccountRepositoryIntf
		publicKeyCredential       CredentialKeyIntf
		accountCredentialRepo     AccountCredentialRepositoryIntf
	)
	container.Make(&liveHouseStaffAccountRepo)
	container.Make(&publicKeyCredential)
	container.Make(&accountCredentialRepo)

	email, err := domain.NewLiveHouseStaffEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	account, err := liveHouseStaffAccountRepo.FindByEmail(email, ctx)
	if err != nil {
		return nil, err
	}

	challenge, err := publicKeyCredential.CreateChallenge()
	if err != nil {
		return nil, err
	}
	c, err := live_house_staff_account_domain.NewCredentialChallenge(challenge.String(), time.Now())

	account.SetCredentialChallenge(c)

	if err := liveHouseStaffAccountRepo.Save(account, ctx); err != nil {
		return nil, err
	}

	credentials, err := accountCredentialRepo.FindByIds(account.CredentialKeys(), ctx)
	if err != nil {
		return nil, err
	}

	options := publicKeyCredential.CreateCredentialAssertionOptions(challenge, "localhost", credentials)

	return &options, nil
}

/*
func LoginLiveHouseStaffAccount(
	emailAddress string,
	reader io.Reader,
	ctx context.Context,
	container container.Container,
) error {

	var (
		liveHouseStaffAccountRepo LiveHouseStaffAccountRepositoryIntf
		publicKeyCredential       CredentialKeyIntf
		accountCredentialRepo     AccountCredentialRepositoryIntf
	)
	container.Make(&liveHouseStaffAccountRepo)
	container.Make(&publicKeyCredential)
	container.Make(&accountCredentialRepo)

	type Body struct {
		Email                       string
		CredentialAssertionResponse protocol.CredentialAssertionResponse
	}

	var body Body

	if err := json.NewDecoder(reader).Decode(&body); err != nil {
		return nil
	}

	body.CredentialAssertionResponse.Parse()

	parsedResponse, err := protocol.ParseCredentialRequestResponseBody(reader)
	if err != nil {
		return err
	}

	account, err := liveHouseStaffAccountRepo.FindByEmailAddress(emailAddress, ctx)
	if err != nil {
		return err
	}

	credential, err := accountCredentialRepo.FindByPublicKeyId(domain.PublicKeyId(parsedResponse.RawID), ctx)
	if err != nil {
		return err
	}

	if err := parsedResponse.Verify(
		credential.AuthenticatorData,
		credential.PublicKey,
		credential.Signature,
		"localhost",
		[]string{"http://localhost"},
	); err != nil {
		return err
	}

	return nil
}
*/
