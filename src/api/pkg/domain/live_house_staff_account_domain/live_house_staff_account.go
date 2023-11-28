package live_house_staff_account_domain

import (
	"errors"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
)

type LiveHouseStaffAccountIntf interface {
	Id() LiveHouseStaffAccountId
	EmailAddress() domain.LiveHouseStaffEmailAddress
	IsProvisional() bool
	ProvisionalToken() *Token
}

type LiveHouseStaffAccountImpl struct {
	id    LiveHouseStaffAccountId
	email domain.LiveHouseStaffEmailAddress

	// 仮登録か否か
	isProvisional bool

	// 仮登録の場合のみ保持する
	provisionalRegistration LiveHouseStaffAccountProvisionalRegistrationIntf
}

type ProvisionalRegistrationParam struct {
	Token string
}

type NewLiveHouseStaffAccountParams struct {
	Id                      string
	Email                   string
	IsProvisional           bool
	ProvisionalRegistration *ProvisionalRegistrationParam
}

func NewLiveHouseStaffAccount(params NewLiveHouseStaffAccountParams) (account LiveHouseStaffAccountIntf, err error) {

	if params.IsProvisional == false && params.ProvisionalRegistration != nil {
		return nil, errors.New("仮登録でない場合は、ProvisionalRegistrationはnilである必要があります。")
	}
	if params.IsProvisional == true && params.ProvisionalRegistration == nil {
		return nil, errors.New("仮登録の場合は、ProvisionalRegistrationは必須です。")
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

	provisionalRegistration, err := NewLiveHouseStaffAccountProvisionalRegistration(params.ProvisionalRegistration.Token)
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
