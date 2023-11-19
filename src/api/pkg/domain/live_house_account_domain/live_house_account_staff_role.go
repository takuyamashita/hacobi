package live_house_account_domain

import "errors"

const (
	roleMaster liveHouseAccountStaffRoleValue = 1 << iota
	roleMmember
)

type liveHouseAccountStaffRoleValue uint

func GetRoleMaster() liveHouseAccountStaffRoleValue {
	return roleMaster
}

func GetRoleMember() liveHouseAccountStaffRoleValue {
	return roleMmember
}

func GetRoleFromValue(value uint) (liveHouseAccountStaffRoleValue, error) {
	switch value {
	case uint(roleMaster):
		return GetRoleMaster(), nil
	case uint(roleMmember):
		return GetRoleMember(), nil
	default:
		return 0, errors.New("invalid role value")
	}
}

type LiveHouseAccountStaffRole struct {
	value liveHouseAccountStaffRoleValue
}

func newLiveHouseAccountStaffRole(role liveHouseAccountStaffRoleValue) (LiveHouseAccountStaffRole, error) {
	return LiveHouseAccountStaffRole{value: role}, nil
}
