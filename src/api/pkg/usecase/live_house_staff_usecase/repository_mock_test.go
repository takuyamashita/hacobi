package live_house_staff_usecase_test

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type _compositIntf interface {
	usecase.LiveHouseStaffRepositoryIntf
	live_house_staff_domain.LiveHouseStaffRepositoryIntf
}

var _ _compositIntf = (*LiveHouseStaffRepositoryMock)(nil)

type LiveHouseStaffRepositoryMock struct {
	Store *Store
}

func NewliveHouseStaffMock(store *Store) *LiveHouseStaffRepositoryMock {
	return &LiveHouseStaffRepositoryMock{Store: store}
}

func (repo LiveHouseStaffRepositoryMock) Save(staff live_house_staff_domain.LiveHouseStaffIntf, ctx context.Context) error {

	repo.Store.Staffs = append(repo.Store.Staffs, staff)

	return nil
}

func (repo LiveHouseStaffRepositoryMock) FindByEmail(emailAddress live_house_staff_domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_domain.LiveHouseStaffIntf, error) {

	var staff live_house_staff_domain.LiveHouseStaffIntf
	for _, v := range repo.Store.Staffs {
		if v.EmailAddress().String() == emailAddress.String() {
			staff = v
		}
	}

	return staff, nil
}

func (repo LiveHouseStaffRepositoryMock) FindById(id live_house_staff_domain.LiveHouseStaffId) (live_house_staff_domain.LiveHouseStaffIntf, error) {

	var staff live_house_staff_domain.LiveHouseStaffIntf
	for _, v := range repo.Store.Staffs {
		if v.Id().String() == id.String() {
			staff = v
		}
	}

	return staff, nil
}
