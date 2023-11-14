package live_house_staff_domain

type LiveHouseStaffId struct {
	value uint64
}

func NewliveHouseStaffId(id uint64) (LiveHouseStaffId, error) {
	return LiveHouseStaffId{id}, nil
}
