package live_house_owner_domain

type LiveHouseOwner struct {
	id       *LiveHouseOwnerId
	name     LiveHouseOwnerName
	email    LiveHouseOwnerEmailAddress
	password LiveHouseOwnerPassword
}

func NewLiveHouseOwner(
	id *LiveHouseOwnerId,
	name LiveHouseOwnerName,
	email LiveHouseOwnerEmailAddress,
	password LiveHouseOwnerPassword,
) (*LiveHouseOwner, error) {

	return &LiveHouseOwner{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}, nil
}

func (owner LiveHouseOwner) Name() LiveHouseOwnerName {
	return owner.name
}

func (owner LiveHouseOwner) EmailAddress() LiveHouseOwnerEmailAddress {
	return owner.email
}

func (owner LiveHouseOwner) Password() LiveHouseOwnerPassword {
	return owner.password
}
