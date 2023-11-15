package live_house_staff_domain

type LiveHouseStaff interface {
	Id() LiveHouseStaffId
	Name() LiveHouseStaffName
	EmailAddress() LiveHouseStaffEmailAddress
	Password() LiveHouseStaffPassword
}

type liveHouseStaff struct {
	id           LiveHouseStaffId
	name         LiveHouseStaffName
	emailAddress LiveHouseStaffEmailAddress
	password     LiveHouseStaffPassword
}

func NewLiveHouseStaff(
	id LiveHouseStaffId,
	name string,
	emailAddress string,
	password string,
) (LiveHouseStaff, error) {

	liveHouseStaffName, err := NewliveHouseStaffName(name)
	if err != nil {
		return nil, err
	}

	liveHouseStaffEmailAddress, err := NewLiveHouseStaffEmailAddress(emailAddress)
	if err != nil {
		return nil, err
	}

	liveHouseStaffPassword, err := NewliveHouseStaffPassword(password)
	if err != nil {
		return nil, err
	}

	return liveHouseStaff{
		id:           id,
		name:         liveHouseStaffName,
		emailAddress: liveHouseStaffEmailAddress,
		password:     liveHouseStaffPassword,
	}, nil
}

func (staff liveHouseStaff) Name() LiveHouseStaffName {
	return staff.name
}

func (staff liveHouseStaff) EmailAddress() LiveHouseStaffEmailAddress {
	return staff.emailAddress
}

func (staff liveHouseStaff) Password() LiveHouseStaffPassword {
	return staff.password
}

func (staff liveHouseStaff) Id() LiveHouseStaffId {
	return staff.id
}
