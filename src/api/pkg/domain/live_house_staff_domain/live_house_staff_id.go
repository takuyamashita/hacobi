package live_house_staff_domain

type LiveHouseStaffId struct {
	value string
}

func NewLiveHouseStaffId(id string) (LiveHouseStaffId, error) {
	return LiveHouseStaffId{id}, nil
}

func (id LiveHouseStaffId) String() string {
	return id.value
}
