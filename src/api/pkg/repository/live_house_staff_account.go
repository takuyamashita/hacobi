package repository

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type LiveHouseStaffAccountRepositoryIntf interface {
	usecase.LiveHouseStaffAccountRepositoryIntf
}

type LiveHouseStaffAccountRepositoryImpl struct {
	db *db.MySQL
}

func NewLiveHouseStaffAccountRepository(db *db.MySQL) LiveHouseStaffAccountRepositoryIntf {
	return &LiveHouseStaffAccountRepositoryImpl{
		db: db,
	}
}

func (repo LiveHouseStaffAccountRepositoryImpl) Save(
	account live_house_staff_account_domain.LiveHouseStaffAccountIntf,
	ctx context.Context,
) error {

	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO live_house_staff_email_authorizations (email, token) VALUES (?, ?)",
		account.EmailAddress().String(),
	)
	if err != nil {
		return err
	}

	return nil
}
