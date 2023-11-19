package live_house_account_domain

type LiveHouseAccountStaffRole struct {
	value int
}

func newLiveHouseAccountStaffRole(role int) (LiveHouseAccountStaffRole, error) {
	return LiveHouseAccountStaffRole{value: role}, nil
}
