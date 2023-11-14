package live_house_owner_domain

import (
	"errors"
	"unicode/utf8"
)

type LiveHouseOwnerPassword struct {
	value string
}

func NewLiveHouseOwnerPassword(password string) (LiveHouseOwnerPassword, error) {

	if (utf8.RuneCountInString(password) < 8) || (utf8.RuneCountInString(password) > 255) {
		return LiveHouseOwnerPassword{}, errors.New("password must be between 8 and 255 characters")
	}

	return LiveHouseOwnerPassword{password}, nil
}

func (p LiveHouseOwnerPassword) String() string {
	return p.value
}
