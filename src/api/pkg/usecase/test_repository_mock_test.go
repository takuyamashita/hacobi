package usecase_test

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_email_authorization_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/repository"
)

// ~~~~~~~~~~~repository.LiveHouseStaffRepositoryIntf~~~~~~~~~~~ //
var _ repository.LiveHouseStaffRepositoryIntf = (*LiveHouseStaffRepositoryMock)(nil)

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
func (repo LiveHouseStaffRepositoryMock) FindByEmail(emailAddress domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_domain.LiveHouseStaffIntf, error) {

	var staff live_house_staff_domain.LiveHouseStaffIntf
	for _, v := range repo.Store.Staffs {
		if v.EmailAddress().String() == emailAddress.String() {
			staff = v
		}
	}

	return staff, nil
}
func (repo LiveHouseStaffRepositoryMock) FindById(id live_house_staff_domain.LiveHouseStaffId, ctx context.Context) (live_house_staff_domain.LiveHouseStaffIntf, error) {

	var staff live_house_staff_domain.LiveHouseStaffIntf
	for _, v := range repo.Store.Staffs {
		if v.Id().String() == id.String() {
			staff = v
		}
	}

	return staff, nil
}

// ~~~~~~~~~~~repository.TransationRepositoryIntf~~~~~~~~~~~ //
var _ repository.TransationRepositoryIntf = (*TransactionRepositoryMock)(nil)

type TransactionRepositoryMock struct {
}

func (repo TransactionRepositoryMock) Begin(ctx context.Context) (commit func() error, rollback func() error, err error) {

	commit = func() error { return nil }
	rollback = func() error { return nil }
	err = nil

	return
}

// ~~~~~~~~~~~usecase.LiveHouseAccountRepositoryIntf~~~~~~~~~~~ //
var _ repository.LiveHouseAccountRepositoryIntf = (*LiveHouseAccountRepositoryMock)(nil)

type LiveHouseAccountRepositoryMock struct {
	Store *Store
}

func (repo LiveHouseAccountRepositoryMock) Save(account live_house_account_domain.LiveHouseAccountIntf, ctx context.Context) error {
	repo.Store.Accounts = append(repo.Store.Accounts, account)
	return nil
}

// ~~~~~~~~~~~usecase.LiveHouseStaffEmailAuthorizationRepositoryIntf~~~~~~~~~~~ //
var _ repository.LiveHouseStaffEmailAuthorizationRepositoryIntf = (*LiveHouseStaffEmailAuthorizationRepositoryMock)(nil)

type LiveHouseStaffEmailAuthorizationRepositoryMock struct {
	Store *Store
}

func (repo LiveHouseStaffEmailAuthorizationRepositoryMock) Save(auth live_house_staff_email_authorization_domain.LiveHouseStaffEmailAuthorizationIntf, ctx context.Context) error {

	repo.Store.LiveHouseStaffEmailAuthorization = append(repo.Store.LiveHouseStaffEmailAuthorization, auth)

	return nil
}
