package live_house_staff_account_domain

import (
	"errors"
	"time"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

type LiveHouseStaffAccountIntf interface {
	Id() LiveHouseStaffAccountId
	EmailAddress() domain.LiveHouseStaffEmailAddress
	IsProvisional() bool
	ProvisionalToken() *Token
	CredentialChallenge() CredentialChallengeIntf
	SetCredentialChallenge(challenge CredentialChallengeIntf) error
	AddCredentialKey(id domain.PublicKeyId) error
}

type LiveHouseStaffAccountImpl struct {
	id    LiveHouseStaffAccountId
	email domain.LiveHouseStaffEmailAddress

	// 仮登録か否か
	isProvisional bool

	// 仮登録の場合のみ保持する
	provisionalRegistration LiveHouseStaffAccountProvisionalRegistrationIntf

	// PublickKeyを登録する際に必要なChallenge
	credentialChallenge CredentialChallengeIntf

	// publicKeys
	credentialKeys []domain.PublicKeyId
}

type ProvisionalRegistrationParam struct {
	Token     string
	CreatedAt time.Time
}

type CredentialChallengeParam struct {
	Challenge string
	CreatedAt time.Time
}

type NewLiveHouseStaffAccountParams struct {
	Id                        string
	Email                     string
	IsProvisional             bool
	ProvisionalRegistration   *ProvisionalRegistrationParam
	CredentialChallengeParams *CredentialChallengeParam
}

func NewLiveHouseStaffAccount(params NewLiveHouseStaffAccountParams) (account LiveHouseStaffAccountIntf, err error) {

	if !params.IsProvisional && params.ProvisionalRegistration != nil {
		return nil, errors.New("仮登録でない場合は、ProvisionalRegistrationはnilである必要があります。")
	}
	if params.IsProvisional {
		if params.ProvisionalRegistration == nil {
			return nil, errors.New("仮登録の場合は、ProvisionalRegistrationは必須です。")
		}
		if params.ProvisionalRegistration.Token == "" {
			return nil, errors.New("仮登録の場合は、ProvisionalRegistration.Tokenは必須です。")
		}
		if params.ProvisionalRegistration.CreatedAt.IsZero() {
			return nil, errors.New("仮登録の場合は、ProvisionalRegistration.CreatedAtは必須です。")
		}
	}

	id := NewLiveHouseStaffAccountId(params.Id)

	email, err := domain.NewLiveHouseStaffEmailAddress(params.Email)
	if err != nil {
		return nil, err
	}

	var challenge CredentialChallengeIntf

	if params.CredentialChallengeParams != nil {
		challenge, err = NewCredentialChallenge(params.CredentialChallengeParams.Challenge, params.CredentialChallengeParams.CreatedAt)
		if err != nil {
			return nil, err
		}
	}

	if params.ProvisionalRegistration == nil {
		return &LiveHouseStaffAccountImpl{
			id:                      id,
			email:                   email,
			isProvisional:           params.IsProvisional,
			credentialChallenge:     challenge,
			provisionalRegistration: nil,
		}, nil
	}

	provisionalRegistration, err := NewLiveHouseStaffAccountProvisionalRegistration(
		params.ProvisionalRegistration.Token,
		params.ProvisionalRegistration.CreatedAt,
	)
	if err != nil {
		return nil, err
	}

	return &LiveHouseStaffAccountImpl{
		id:                      id,
		email:                   email,
		isProvisional:           params.IsProvisional,
		provisionalRegistration: provisionalRegistration,
		credentialChallenge:     challenge,
	}, nil
}

func (account LiveHouseStaffAccountImpl) Id() LiveHouseStaffAccountId {
	return account.id
}

func (account LiveHouseStaffAccountImpl) EmailAddress() domain.LiveHouseStaffEmailAddress {
	return account.email
}

func (account LiveHouseStaffAccountImpl) ProvisionalToken() *Token {

	if !account.isProvisional {
		return nil
	}

	if account.provisionalRegistration == nil {
		return nil
	}

	t := account.provisionalRegistration.Token()

	return &t
}

func (account LiveHouseStaffAccountImpl) IsProvisional() bool {
	return account.isProvisional
}

func (account *LiveHouseStaffAccountImpl) SetCredentialChallenge(challenge CredentialChallengeIntf) error {

	account.credentialChallenge = challenge

	return nil
}

func (account LiveHouseStaffAccountImpl) CredentialChallenge() CredentialChallengeIntf {
	return account.credentialChallenge
}

func (account *LiveHouseStaffAccountImpl) AddCredentialKey(id domain.PublicKeyId) error {

	account.credentialKeys = append(account.credentialKeys, id)

	return nil
}
