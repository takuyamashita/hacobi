package live_house_staff_domain

import (
	"errors"
	"unicode/utf8"
)

type LiveHouseStaffDisplayName struct {
	value string
}

func newLiveHouseStaffDisplayName(name string) (LiveHouseStaffDisplayName, error) {

	if (utf8.RuneCountInString(name) < 1) || (utf8.RuneCountInString(name) > 255) {
		return LiveHouseStaffDisplayName{}, errors.New("name must be between 1 and 255 characters")
	}

	return LiveHouseStaffDisplayName{
		value: name,
	}, nil
}

func (n LiveHouseStaffDisplayName) String() string {
	return n.value
}
