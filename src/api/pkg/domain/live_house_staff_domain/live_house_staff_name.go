package live_house_staff_domain

import (
	"errors"
	"unicode/utf8"
)

type LiveHouseStaffName struct {
	value string
}

func NewliveHouseStaffName(name string) (LiveHouseStaffName, error) {

	if (utf8.RuneCountInString(name) < 1) || (utf8.RuneCountInString(name) > 255) {
		return LiveHouseStaffName{}, errors.New("name must be between 1 and 255 characters")
	}

	return LiveHouseStaffName{
		value: name,
	}, nil
}

func (n LiveHouseStaffName) String() string {
	return n.value
}
