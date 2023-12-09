package live_house_staff_account_domain

import (
	"errors"
	"unicode/utf8"
)

type displayName string

func NewDisplayName(value string) (displayName, error) {

	if value == "" {
		return "", errors.New("displayName must not be empty")
	}

	if utf8.RuneCountInString(value) > 15 {
		return "", errors.New("displayName must be less than 15 characters")
	}

	return displayName(value), nil
}

type LiveHouseStaffProfileIntf interface {
	DisplayName() displayName
}

type LiveHouseStaffProfileImpl struct {
	displayName displayName
}

func NewLiveHouseStaffProfile(displayName string) (LiveHouseStaffProfileIntf, error) {

	dipName, err := NewDisplayName(displayName)
	if err != nil {
		return nil, err
	}

	return LiveHouseStaffProfileImpl{
		displayName: dipName,
	}, nil
}

func (p LiveHouseStaffProfileImpl) DisplayName() displayName {
	return p.displayName
}
