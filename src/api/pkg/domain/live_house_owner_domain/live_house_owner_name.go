package live_house_owner_domain

import (
	"errors"
	"unicode/utf8"
)

type LiveHouseOwnerName struct {
	value string
}

func NewLiveHouseOwnerName(name string) (LiveHouseOwnerName, error) {

	if (utf8.RuneCountInString(name) < 1) || (utf8.RuneCountInString(name) > 255) {
		return LiveHouseOwnerName{}, errors.New("name must be between 1 and 255 characters")
	}

	return LiveHouseOwnerName{
		value: name,
	}, nil
}

func (n LiveHouseOwnerName) String() string {
	return n.value
}
