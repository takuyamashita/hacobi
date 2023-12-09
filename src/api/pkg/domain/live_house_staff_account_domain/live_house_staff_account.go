package live_house_staff_account_domain

import (
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
	CredentialKeys() []domain.PublicKeyId
	SetProfile(profile LiveHouseStaffProfileIntf) error
	Profile() LiveHouseStaffProfileIntf
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

	credentialKeys []domain.PublicKeyId
	profile        LiveHouseStaffProfileIntf
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

func (account LiveHouseStaffAccountImpl) CredentialKeys() []domain.PublicKeyId {
	return account.credentialKeys
}

func (account LiveHouseStaffAccountImpl) Profile() LiveHouseStaffProfileIntf {
	return account.profile
}

func (account *LiveHouseStaffAccountImpl) SetProfile(profile LiveHouseStaffProfileIntf) error {

	account.profile = profile

	return nil
}
