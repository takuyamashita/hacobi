package live_house_account_domain

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type LiveHouseAccountStaffs []LiveHouseAccountStaff

func newLiveHouseAccountStaffs() LiveHouseAccountStaffs {
	return LiveHouseAccountStaffs{}
}

type LiveHouseAccountStaff struct {
	id   live_house_staff_domain.LiveHouseStaffId
	role LiveHouseAccountStaffRole
}

func newLiveHouseAccountStaff(id live_house_staff_domain.LiveHouseStaffId, role liveHouseAccountStaffRoleValue) (LiveHouseAccountStaff, error) {

	liveHouseStaffRole, err := newLiveHouseAccountStaffRole(role)
	if err != nil {
		return LiveHouseAccountStaff{}, err
	}

	return LiveHouseAccountStaff{
		id:   id,
		role: liveHouseStaffRole,
	}, nil
}

func (s LiveHouseAccountStaff) Id() live_house_staff_domain.LiveHouseStaffId {
	return s.id
}

func (s LiveHouseAccountStaff) Role() LiveHouseAccountStaffRole {
	return s.role
}
