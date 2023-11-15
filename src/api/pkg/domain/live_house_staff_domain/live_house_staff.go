package live_house_staff_domain

type LiveHouseStaff struct {
	id           LiveHouseStaffId
	name         LiveHouseStaffName
	emailAddress LiveHouseStaffEmailAddress
	password     LiveHouseStaffPassword
}

func NewliveHouseStaff(
	id LiveHouseStaffId,
	name LiveHouseStaffName,
	emailAddress LiveHouseStaffEmailAddress,
	password LiveHouseStaffPassword,
) (LiveHouseStaff, error) {

	return LiveHouseStaff{
		id:           id,
		name:         name,
		emailAddress: emailAddress,
		password:     password,
	}, nil
}

func (owner LiveHouseStaff) Name() LiveHouseStaffName {
	return owner.name
}

func (owner LiveHouseStaff) EmailAddress() LiveHouseStaffEmailAddress {
	return owner.emailAddress
}

func (owner LiveHouseStaff) Password() LiveHouseStaffPassword {
	return owner.password
}

func (owner LiveHouseStaff) Id() LiveHouseStaffId {
	return owner.id
}
