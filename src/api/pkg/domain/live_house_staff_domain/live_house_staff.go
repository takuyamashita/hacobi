package live_house_staff_domain

type LiveHouseStaff struct {
	id       *LiveHouseStaffId
	name     LiveHouseStaffName
	email    LiveHouseStaffEmailAddress
	password LiveHouseStaffPassword
}

func NewliveHouseStaff(
	id *LiveHouseStaffId,
	name LiveHouseStaffName,
	email LiveHouseStaffEmailAddress,
	password LiveHouseStaffPassword,
) (LiveHouseStaff, error) {

	return LiveHouseStaff{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}, nil
}

func (owner LiveHouseStaff) Name() LiveHouseStaffName {
	return owner.name
}

func (owner LiveHouseStaff) EmailAddress() LiveHouseStaffEmailAddress {
	return owner.email
}

func (owner LiveHouseStaff) Password() LiveHouseStaffPassword {
	return owner.password
}
