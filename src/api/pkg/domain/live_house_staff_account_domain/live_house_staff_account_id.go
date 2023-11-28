package live_house_staff_account_domain

type LiveHouseStaffAccountId struct {
	value string
}

func NewLiveHouseStaffAccountId(value string) LiveHouseStaffAccountId {
	return LiveHouseStaffAccountId{
		value: value,
	}
}

func (id LiveHouseStaffAccountId) String() string {
	return id.value
}
