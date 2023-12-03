package usecase

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"io"
	"time"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/account_credential_domain"
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

	options := publicKeyCredential.CreateCredentialAssertionOptions(challenge, "localhost", nil)

	return &options, nil
}

func LoginLiveHouseStaffAccount(
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
		return err
	}

	parsedResponse, err := body.CredentialAssertionResponse.Parse()
	if err != nil {
		return err
	}

	email, err := domain.NewLiveHouseStaffEmailAddress(body.Email)
	if err != nil {
		return err
	}

	account, err := liveHouseStaffAccountRepo.FindByEmail(email, ctx)
	if err != nil {
		return err
	}

	credentials, err := accountCredentialRepo.FindByIds(account.CredentialKeys(), ctx)
	if err != nil {
		return err
	}

	var credential account_credential_domain.AccountCredentialIntf
	for _, c := range credentials {
		if bytes.Equal(c.PublicKeyId(), parsedResponse.RawID) {
			credential = c
			break
		}
	}
	if credential == nil {
		return errors.New("credential not found")
	}

	err = parsedResponse.Verify(
		account.CredentialChallenge().Challenge().String(),
		"localhost", []string{"http://localhost"},
		"",
		true,
		credential.PublicKey(),
	)

	if err != nil {
		return err
	}

	credential.SetAuthenticatorCount(parsedResponse.Response.AuthenticatorData.Counter)
	credential.SetUserPresent(parsedResponse.Response.AuthenticatorData.Flags.HasUserPresent())
	credential.SetUserVerified(parsedResponse.Response.AuthenticatorData.Flags.HasUserVerified())
	credential.SetBackupEligible(parsedResponse.Response.AuthenticatorData.Flags.HasAttestedCredentialData())
	credential.SetBackupState(parsedResponse.Response.AuthenticatorData.Flags.HasBackupState())

	if err := accountCredentialRepo.Save(credential, ctx); err != nil {
		return err
	}

	return err
}
