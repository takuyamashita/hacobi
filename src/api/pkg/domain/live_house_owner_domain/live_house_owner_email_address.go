package live_house_owner_domain

import (
	"errors"
	"regexp"
	"strings"
)

var (
	mailAddressRegexPatern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type LiveHouseOwnerEmailAddress struct {
	localPart string
	domain    string
}

func NewLiveHouseOwnerEmailAddress(address string) (LiveHouseOwnerEmailAddress, error) {

	if !mailAddressRegexPatern.MatchString(address) {
		return LiveHouseOwnerEmailAddress{}, errors.New("email is invalid")
	}

	addressParts := strings.Split(address, "@")
	if len(addressParts) != 2 {
		return LiveHouseOwnerEmailAddress{}, errors.New("email is invalid")
	}

	return LiveHouseOwnerEmailAddress{
		localPart: addressParts[0],
		domain:    addressParts[1],
	}, nil
}

func (a LiveHouseOwnerEmailAddress) String() string {
	return a.localPart + "@" + a.domain
}
