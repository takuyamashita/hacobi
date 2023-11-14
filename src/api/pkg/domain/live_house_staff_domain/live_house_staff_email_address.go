package live_house_staff_domain

import (
	"errors"
	"regexp"
	"strings"
)

var (
	mailAddressRegexPatern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type LiveHouseStaffEmailAddress struct {
	localPart string
	domain    string
}

func NewLiveHouseStaffEmailAddress(address string) (LiveHouseStaffEmailAddress, error) {

	if !mailAddressRegexPatern.MatchString(address) {
		return LiveHouseStaffEmailAddress{}, errors.New("email is invalid")
	}

	addressParts := strings.Split(address, "@")
	if len(addressParts) != 2 {
		return LiveHouseStaffEmailAddress{}, errors.New("email is invalid")
	}

	return LiveHouseStaffEmailAddress{
		localPart: addressParts[0],
		domain:    addressParts[1],
	}, nil
}

func (a LiveHouseStaffEmailAddress) String() string {
	return a.localPart + "@" + a.domain
}
