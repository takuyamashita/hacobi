package live_house_owner_domain

type LiveHouseOwnerId struct {
	value uint64
}

func NewLiveHouseOwnerId(id uint64) (LiveHouseOwnerId, error) {
	return LiveHouseOwnerId{id}, nil
}
