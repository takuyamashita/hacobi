package repository

import (
	"context"
	"database/sql"

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

	tx, err := repo.db.BeginTx(ctx, &sql.TxOptions{})
	if err != nil {
		return err
	}

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO live_house_staff_accounts (id, email, is_provisional) VALUES (?, ?, ?)",
		account.Id().String(),
		account.EmailAddress().String(),
		account.IsProvisional(),
	)
	if err != nil {
		return err
	}

	if account.IsProvisional() {
		return tx.Commit()
	}

	_, err = tx.ExecContext(
		ctx,
		"INSERT INTO live_house_staff_account_provisional_registrations (live_house_staff_account_id, token) VALUES (?, ?)",
		account.Id().String(),
		account.ProvisionalToken().String(),
	)

	return tx.Commit()

}
