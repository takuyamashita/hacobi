package usecase_test

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
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

// ~~~~~~~~~~~usecase.LiveHouseStaffAccountProvisionalRegistrationRepositoryIntf~~~~~~~~~~~ //
var _ repository.LiveHouseStaffAccountRepositoryIntf = (*LiveHouseStaffAccountRepositoryMock)(nil)

type LiveHouseStaffAccountRepositoryMock struct {
	Store *Store
}

func (repo LiveHouseStaffAccountRepositoryMock) Save(account live_house_staff_account_domain.LiveHouseStaffAccountIntf, ctx context.Context) error {

	repo.Store.LiveHouseStaffAccounts = append(repo.Store.LiveHouseStaffAccounts, account)

	return nil
}

func (repo LiveHouseStaffAccountRepositoryMock) FindById(
	id live_house_staff_account_domain.LiveHouseStaffAccountId,
	ctx context.Context,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	var account live_house_staff_account_domain.LiveHouseStaffAccountIntf
	for _, v := range repo.Store.LiveHouseStaffAccounts {
		if v.Id().String() == id.String() {
			account = v
		}
	}

	return account, nil
}

func (repo LiveHouseStaffAccountRepositoryMock) FindByEmail(emailAddress domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	var account live_house_staff_account_domain.LiveHouseStaffAccountIntf
	for _, v := range repo.Store.LiveHouseStaffAccounts {
		if v.EmailAddress().String() == emailAddress.String() {
			account = v
		}
	}

	return account, nil
}

func (repo LiveHouseStaffAccountRepositoryMock) FindByProvisionalRegistrationToken(token live_house_staff_account_domain.Token, ctx context.Context) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	var account live_house_staff_account_domain.LiveHouseStaffAccountIntf
	for _, v := range repo.Store.LiveHouseStaffAccounts {
		if v.ProvisionalToken().String() == token.String() {
			account = v
		}
	}

	return account, nil
}

// ~~~~~~~~~~~domain.MainIntf~~~~~~~~~~~ //
var _ domain.MailIntf = (*MailMock)(nil)

type MailMock struct {
	Store *Store
}

type SentMail struct {
	To      string
	Subject string
	Body    string
}

func (m MailMock) Send(to string, subject string, body string) error {

	m.Store.SentMails = append(m.Store.SentMails, SentMail{
		To:      to,
		Subject: subject,
		Body:    body,
	})

	return nil
}
