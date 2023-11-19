package live_house_staff_domain

type LiveHouseStaffIntf interface {
	Id() LiveHouseStaffId
	Name() LiveHouseStaffName
	EmailAddress() LiveHouseStaffEmailAddress
	Password() LiveHouseStaffPassword
}

type liveHouseStaffImpl struct {
	id           LiveHouseStaffId
	name         LiveHouseStaffName
	emailAddress LiveHouseStaffEmailAddress
	password     LiveHouseStaffPassword
}

func NewLiveHouseStaff(
	id string,
	name string,
	emailAddress string,
	password string,
) (LiveHouseStaffIntf, error) {

	liveHouseStaffId, err := NewLiveHouseStaffId(id)
	if err != nil {
		return nil, err
	}

	liveHouseStaffName, err := newLiveHouseStaffName(name)
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
		name:         liveHouseStaffName,
		emailAddress: liveHouseStaffEmailAddress,
		password:     liveHouseStaffPassword,
	}, nil
}

func (staff liveHouseStaffImpl) Name() LiveHouseStaffName {
	return staff.name
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
