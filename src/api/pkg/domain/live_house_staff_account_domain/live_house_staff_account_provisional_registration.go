package live_house_staff_account_domain

type LiveHouseStaffAccountProvisionalRegistrationIntf interface {
	Token() Token
}

type LiveHouseStaffAccountProvisionalRegistrationImpl struct {
	token Token
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
