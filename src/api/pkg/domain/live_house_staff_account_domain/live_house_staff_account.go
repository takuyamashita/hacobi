package live_house_staff_account_domain

import "github.com/takuyamashita/hacobi/src/api/pkg/domain"

type LiveHouseStaffAccountIntf interface {
	EmailAddress() domain.LiveHouseStaffEmailAddress
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

func NewLiveHouseStaffAccount(
	id LiveHouseStaffAccountId,
	email domain.LiveHouseStaffEmailAddress,
	isProvisional bool,
	provisionalRegistration LiveHouseStaffAccountProvisionalRegistrationIntf,
) LiveHouseStaffAccountIntf {
	return &LiveHouseStaffAccountImpl{
		id:                      id,
		email:                   email,
		isProvisional:           isProvisional,
		provisionalRegistration: provisionalRegistration,
	}
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
