package livehouseowner

import (
	"errors"
	"regexp"
	"strings"
	"unicode/utf8"
)

var (
	mailAddressRegexPatern = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")
)

type LiveHouseOwner interface {
	Name() livehouseOwnerName
	EmailAddress() liveHouseOwnerEmailAddress
	Password() liveHouseOwnerPassword
}

type liveHouseOwner struct {
	id       *LiveHouseOwnerId
	name     livehouseOwnerName
	email    liveHouseOwnerEmailAddress
	password liveHouseOwnerPassword
}

func NewLiveHouseOwner(
	id *LiveHouseOwnerId,
	name livehouseOwnerName,
	email liveHouseOwnerEmailAddress,
	password liveHouseOwnerPassword,
) (LiveHouseOwner, error) {

	return &liveHouseOwner{
		id:       id,
		name:     name,
		email:    email,
		password: password,
	}, nil
}

func (owner liveHouseOwner) Name() livehouseOwnerName {
	return owner.name
}

func (owner liveHouseOwner) EmailAddress() liveHouseOwnerEmailAddress {
	return owner.email
}

func (owner liveHouseOwner) Password() liveHouseOwnerPassword {
	return owner.password
}

type livehouseOwnerName string

func NewLiveHouseOwnerName(name string) (livehouseOwnerName, error) {

	if (utf8.RuneCountInString(name) < 1) || (utf8.RuneCountInString(name) > 255) {
		return "", errors.New("name must be between 1 and 255 characters")
	}

	return livehouseOwnerName(name), nil
}

func (n livehouseOwnerName) String() string {
	return string(n)
}

type LiveHouseOwnerId struct {
	value uint64
}

func NewLiveHouseOwnerId(id uint64) (LiveHouseOwnerId, error) {
	return LiveHouseOwnerId{id}, nil
}

type liveHouseOwnerEmailAddress struct {
	localPart string
	domain    string
}

func NewLiveHouseOwnerEmailAddress(address string) (liveHouseOwnerEmailAddress, error) {

	if !mailAddressRegexPatern.MatchString(address) {
		return liveHouseOwnerEmailAddress{}, errors.New("email is invalid")
	}

	addressParts := strings.Split(address, "@")
	if len(addressParts) != 2 {
		return liveHouseOwnerEmailAddress{}, errors.New("email is invalid")
	}

	return liveHouseOwnerEmailAddress{
		localPart: addressParts[0],
		domain:    addressParts[1],
	}, nil
}

func (a liveHouseOwnerEmailAddress) String() string {
	return a.localPart + "@" + a.domain
}

type liveHouseOwnerPassword string

func NewLiveHouseOwnerPassword(password string) (liveHouseOwnerPassword, error) {

	if (utf8.RuneCountInString(password) < 8) || (utf8.RuneCountInString(password) > 255) {
		return "", errors.New("password must be between 8 and 255 characters")
	}

	return liveHouseOwnerPassword(password), nil
}

func (p liveHouseOwnerPassword) String() string {
	return string(p)
}
