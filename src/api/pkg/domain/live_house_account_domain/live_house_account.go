package live_house_account_domain

import (
	"errors"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type LiveHouseAccountIntf interface {
	AddStaff(id live_house_staff_domain.LiveHouseStaffId, role liveHouseAccountStaffRoleValue) error
	Id() LiveHouseAccountId
}

type liveHouseAccountImpl struct {
	id        LiveHouseAccountId
	staffs    LiveHouseAccountStaffs
	liveHouse *live_house_domain.LiveHouseId
}

type StaffParams struct {
	Id   live_house_staff_domain.LiveHouseStaffId
	Role liveHouseAccountStaffRoleValue
}

func NewLiveHouseAccount(id string, staffs []StaffParams, liveHouseId *live_house_domain.LiveHouseId) (LiveHouseAccountIntf, error) {

	if staffs == nil || len(staffs) == 0 {
		return nil, errors.New("staffs must have at least one staff")
	}

	liveHouseAccountId, err := NewLiveHouseAccountId(id)
	if err != nil {
		return nil, err
	}

	liveHouseAccountStaffs := newLiveHouseAccountStaffs()

	for _, staff := range staffs {
		s, err := newLiveHouseAccountStaff(staff.Id, 1)
		liveHouseAccountStaffs = append(liveHouseAccountStaffs, s)
		if err != nil {
			return nil, err
		}
	}

	liveHouseAccount := &liveHouseAccountImpl{
		id:        liveHouseAccountId,
		liveHouse: liveHouseId,
		staffs:    liveHouseAccountStaffs,
	}

	return liveHouseAccount, nil
}

func (a *liveHouseAccountImpl) AddStaff(id live_house_staff_domain.LiveHouseStaffId, role liveHouseAccountStaffRoleValue) error {

	staff, err := newLiveHouseAccountStaff(id, role)
	if err != nil {
		return err
	}

	a.staffs = append(a.staffs, staff)

	return nil
}

func (a liveHouseAccountImpl) Id() LiveHouseAccountId {
	return a.id
}
