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
	CredentialChallenges() []CredentialChallengeIntf
	CredentialLength() int
	AddCredentialChallenge(challenge CredentialChallengeIntf) error
}

type LiveHouseStaffAccountImpl struct {
	id    LiveHouseStaffAccountId
	email domain.LiveHouseStaffEmailAddress

	// 仮登録か否か
	isProvisional bool

	// 仮登録の場合のみ保持する
	provisionalRegistration LiveHouseStaffAccountProvisionalRegistrationIntf

	// PublickKeyを登録する際に必要なChallenge
	credentialChallenges []CredentialChallengeIntf
}

type ProvisionalRegistrationParam struct {
	Token     string
	CreatedAt time.Time
}

type NewLiveHouseStaffAccountParams struct {
	Id                      string
	Email                   string
	IsProvisional           bool
	ProvisionalRegistration *ProvisionalRegistrationParam
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

	if params.ProvisionalRegistration == nil {
		return &LiveHouseStaffAccountImpl{
			id:                      id,
			email:                   email,
			isProvisional:           params.IsProvisional,
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

func (account *LiveHouseStaffAccountImpl) AddCredentialChallenge(challenge CredentialChallengeIntf) error {

	if !account.isProvisional {
		return errors.New("仮登録でない場合は、CredentialChallengeを追加できません。")
	}

	if account.provisionalRegistration == nil {
		return errors.New("仮登録の場合は、ProvisionalRegistrationはnilではありません。")
	}

	account.credentialChallenges = append(account.credentialChallenges, challenge)

	return nil
}

func (account LiveHouseStaffAccountImpl) CredentialChallenges() []CredentialChallengeIntf {
	return account.credentialChallenges
}

func (account LiveHouseStaffAccountImpl) CredentialLength() int {

	if account.credentialChallenges == nil {
		return 0
	}

	return len(account.credentialChallenges)
}
