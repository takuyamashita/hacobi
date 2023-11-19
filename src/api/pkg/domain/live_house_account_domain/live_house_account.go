package live_house_account_domain

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type liveHouseAccount struct {
	id        LiveHouseAccountId
	staffs    LiveHouseAccountStaffs
	liveHouse *live_house_domain.LiveHouseId
}

type StaffParams struct {
	id   live_house_staff_domain.LiveHouseStaffId
	role int
}

func NewLiveHouseAccount(id string, staffs []StaffParams, liveHouseId live_house_domain.LiveHouseId) (*liveHouseAccount, error) {

	liveHouseAccountId, err := NewLiveHouseAccountId(id)
	if err != nil {
		return nil, err
	}

	liveHouseAccountStaffs := newLiveHouseAccountStaffs()

	for _, staff := range staffs {
		s, err := newLiveHouseAccountStaff(staff.id, 1)
		liveHouseAccountStaffs = append(liveHouseAccountStaffs, s)
		if err != nil {
			return nil, err
		}
	}

	liveHouseAccount := &liveHouseAccount{
		id:        liveHouseAccountId,
		liveHouse: &liveHouseId,
		staffs:    liveHouseAccountStaffs,
	}

	return liveHouseAccount, nil
}

func (a *liveHouseAccount) AddStaff(id live_house_staff_domain.LiveHouseStaffId, role int) error {
	staff, err := newLiveHouseAccountStaff(id, role)
	if err != nil {
		return err
	}

	a.staffs = append(a.staffs, staff)

	return nil
}
