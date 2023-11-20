package live_house_staff_domain

type LiveHouseStaffIntf interface {
	Id() LiveHouseStaffId
	DisplayName() LiveHouseStaffDisplayName
	EmailAddress() LiveHouseStaffEmailAddress
	Password() LiveHouseStaffPassword
}

type liveHouseStaffImpl struct {
	id           LiveHouseStaffId
	displayName  LiveHouseStaffDisplayName
	emailAddress LiveHouseStaffEmailAddress
	password     LiveHouseStaffPassword
}

func NewLiveHouseStaff(
	id string,
	displayName string,
	emailAddress string,
	password string,
) (LiveHouseStaffIntf, error) {

	liveHouseStaffId, err := NewLiveHouseStaffId(id)
	if err != nil {
		return nil, err
	}

	liveHouseStaffDisplayName, err := newLiveHouseStaffDisplayName(displayName)
	if err != nil {
		return nil, err
	}

	liveHouseStaffEmailAddress, err := newLiveHouseStaffEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	liveHouseStaffPassword, err := newLiveHouseStaffPassword(password)
	if err != nil {
		return nil, err
	}

	return liveHouseStaffImpl{
		id:           liveHouseStaffId,
		displayName:  liveHouseStaffDisplayName,
		emailAddress: liveHouseStaffEmailAddress,
		password:     liveHouseStaffPassword,
	}, nil
}

func (staff liveHouseStaffImpl) DisplayName() LiveHouseStaffDisplayName {
	return staff.displayName
}

func (staff liveHouseStaffImpl) EmailAddress() LiveHouseStaffEmailAddress {
	return staff.emailAddress
}

func (staff liveHouseStaffImpl) Password() LiveHouseStaffPassword {
	return staff.password
}

func (staff liveHouseStaffImpl) Id() LiveHouseStaffId {
	return staff.id
}
