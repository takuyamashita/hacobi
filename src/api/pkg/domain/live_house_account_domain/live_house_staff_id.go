package live_house_account_domain

type LiveHouseStaffId struct {
	value uint64
}

func NewliveHouseStaffId(id uint64) (LiveHouseStaffId, error) {
	return LiveHouseStaffId{id}, nil
}
