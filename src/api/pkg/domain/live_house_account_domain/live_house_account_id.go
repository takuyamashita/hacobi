package live_house_account_domain

type LiveHouseAccountId struct {
	value string
}

func NewLiveHouseAccountId(value string) (LiveHouseAccountId, error) {
	return LiveHouseAccountId{value: value}, nil
}

func (id LiveHouseAccountId) String() string {
	return id.value
}
