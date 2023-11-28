package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type LiveHouseStaffAccountRepositoryIntf interface {
	usecase.LiveHouseStaffAccountRepositoryIntf
	live_house_staff_account_domain.LiveHouseStaffAccountRepositoryIntf
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

	if !account.IsProvisional() {
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

func (repo LiveHouseStaffAccountRepositoryImpl) FindByEmail(
	emailAddress domain.LiveHouseStaffEmailAddress,
	ctx context.Context,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	int2Bool := func(i int) bool {
		if i == 1 {
			return true
		}
		return false
	}

	row := repo.db.QueryRowContext(
		ctx,
		"SELECT id, email, is_provisional FROM live_house_staff_accounts WHERE email = ?",
		emailAddress.String(),
	)

	var (
		id            string
		email         string
		isProvisional int
	)

	if err := row.Scan(&id, &email, &isProvisional); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	if isProvisional == 1 {
		row := repo.db.QueryRowContext(
			ctx,
			"SELECT token, created_at FROM live_house_staff_account_provisional_registrations WHERE live_house_staff_account_id = ?",
			id,
		)

		var (
			token     string
			createdAt []uint8
		)
		if err := row.Scan(&token, &createdAt); err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
		}

		createdAtTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}

		return live_house_staff_account_domain.NewLiveHouseStaffAccount(live_house_staff_account_domain.NewLiveHouseStaffAccountParams{
			Id:            id,
			Email:         email,
			IsProvisional: int2Bool(isProvisional),
			ProvisionalRegistration: &live_house_staff_account_domain.ProvisionalRegistrationParam{
				Token:     token,
				CreatedAt: createdAtTime,
			},
		})
	}

	return live_house_staff_account_domain.NewLiveHouseStaffAccount(live_house_staff_account_domain.NewLiveHouseStaffAccountParams{
		Id:            id,
		Email:         email,
		IsProvisional: int2Bool(isProvisional),
	})

}
