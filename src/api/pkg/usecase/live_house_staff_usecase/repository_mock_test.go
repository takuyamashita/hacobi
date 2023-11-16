package live_house_staff_usecase

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type _compositIntf interface {
	LiveHouseStaffRepository
	live_house_staff_domain.LiveHouseStaffRepository
}

var _ _compositIntf = (*LiveHouseStaffRepositoryMock)(nil)

type Store struct {
	_compositIntf
	Users map[string]live_house_staff_domain.LiveHouseStaffIntf
}

type LiveHouseStaffRepositoryMock struct {
	Store Store
}

func NewliveHouseStaffMock(store Store) *LiveHouseStaffRepositoryMock {
	return &LiveHouseStaffRepositoryMock{Store: store}
}

func (repo LiveHouseStaffRepositoryMock) Save(owner live_house_staff_domain.LiveHouseStaffIntf, ctx context.Context) error {

	return nil
}

func (repo LiveHouseStaffRepositoryMock) FindByEmail(emailAddress live_house_staff_domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_domain.LiveHouseStaffIntf, error) {

	var user live_house_staff_domain.LiveHouseStaffIntf
	for _, v := range repo.Store.Users {
		if v.EmailAddress().String() == emailAddress.String() {
			user = v
		}
	}

	return user, nil
}
