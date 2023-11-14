package live_house_staff_domain

import (
	"errors"
	"unicode/utf8"
)

type LiveHouseStaffPassword struct {
	value string
}

func NewliveHouseStaffPassword(password string) (LiveHouseStaffPassword, error) {

	if (utf8.RuneCountInString(password) < 8) || (utf8.RuneCountInString(password) > 255) {
		return LiveHouseStaffPassword{}, errors.New("password must be between 8 and 255 characters")
	}

	return LiveHouseStaffPassword{password}, nil
}

func (p LiveHouseStaffPassword) String() string {
	return p.value
}
