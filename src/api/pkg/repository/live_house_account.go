package repository

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

var _ usecase.LiveHouseAccountRepositoryIntf = (*LiveHouseAccount)(nil)

type LiveHouseAccount struct {
	db *db.MySQL
}

func NewliveHouseAccount(db *db.MySQL) *LiveHouseAccount {
	return &LiveHouseAccount{
		db: db,
	}
}

func (repo LiveHouseAccount) Save(account live_house_account_domain.LiveHouseAccountIntf, ctx context.Context) error {

	//xxx: TODO transaction

	/*


		CREATE TABLE IF NOT EXISTS live_house_account_live_house_staff (
		    live_house_account_id VARCHAR(36) NOT NULL,
		    live_house_staff_id VARCHAR(36) NOT NULL,
		    role INT UNSIGNED NOT NULL DEFAULT 0,
		    created_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6),
		    updated_at DATETIME(6) NOT NULL DEFAULT CURRENT_TIMESTAMP(6) ON UPDATE CURRENT_TIMESTAMP(6),
		    INDEX live_house_account_id (live_house_account_id),
		    UNIQUE KEY live_house_account_id_live_house_staff_id (live_house_account_id, live_house_staff_id)
		) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci;

	*/

	tx, err := repo.db.BeginTx(ctx, nil)
	defer tx.Commit()
	if err != nil {
		return err
	}

	{

		_, err := tx.ExecContext(
			ctx,
			"INSERT INTO live_house_accounts (id) VALUES (?)", account.Id().String(),
		)

		if err != nil {
			tx.Rollback()
			return err
		}

		query := "INSERT INTO live_house_account_live_house_staff (live_house_account_id, live_house_staff_id, role) VALUES "
		args := []interface{}{}

		for _, staff := range account.Staffs() {

			query += "(?, ?, ?),"
			args = append(args, account.Id().String(), staff.Id().String(), staff.Role().Number())

		}

		query = query[:len(query)-1]

		_, err = repo.db.ExecContext(
			ctx,
			query,
			args...,
		)

		if err != nil {
			tx.Rollback()
			return err
		}

	}

	return nil
}
