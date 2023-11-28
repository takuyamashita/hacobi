package live_house_staff_account_domain

import "time"

type LiveHouseStaffAccountProvisionalRegistrationIntf interface {
	Token() Token
	CreatedAt() time.Time
}

type LiveHouseStaffAccountProvisionalRegistrationImpl struct {
	token     Token
	createdAt time.Time
}

func NewLiveHouseStaffAccountProvisionalRegistration(token string) (LiveHouseStaffAccountProvisionalRegistrationIntf, error) {

	tkn, err := newTokenFromHexString(token)
	if err != nil {
		return nil, err
	}

	return &LiveHouseStaffAccountProvisionalRegistrationImpl{
		token: tkn,
	}, nil
}

func (auth LiveHouseStaffAccountProvisionalRegistrationImpl) Token() Token {
	return auth.token
}

func (auth LiveHouseStaffAccountProvisionalRegistrationImpl) CreatedAt() time.Time {
	return auth.createdAt
}
