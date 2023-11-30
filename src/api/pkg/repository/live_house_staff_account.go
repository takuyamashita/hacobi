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
		`
			INSERT INTO live_house_staff_accounts 
				(id, email, is_provisional)
			VALUES 
				(?, ?, ?) AS new
			ON DUPLICATE KEY UPDATE
				email = new.email,
				is_provisional = new.is_provisional
		`,
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
		`
			INSERT INTO live_house_staff_account_provisional_registrations
				(live_house_staff_account_id, token)
			VALUES
				(?, ?) AS new
			ON DUPLICATE KEY UPDATE
				token = new.token
		`,
		account.Id().String(),
		account.ProvisionalToken().String(),
	)

	if err != nil {
		tx.Rollback()
		return err
	}

	if account.CredentialChallengeLength() == 0 {

		_, err = tx.ExecContext(
			ctx,
			"DELETE FROM live_house_staff_account_credential_challenges WHERE live_house_staff_account_id = ?",
			account.Id().String(),
		)

		if err != nil {
			tx.Rollback()
			return err
		}

		return tx.Commit()
	} else {

	}

	stmt := "REPLACE INTO live_house_staff_account_credential_challenges (live_house_staff_account_id, challenge) VALUES "
	args := []interface{}{}
	for i, v := range account.CredentialChallenges() {
		args = append(args, account.Id().String(), v.Challenge().String())
		if i == 0 {
			stmt += "(?, ?)"
			continue
		}
		stmt += ", (?, ?)"
	}

	_, err = tx.ExecContext(
		ctx,
		stmt,
		args...,
	)
	if err != nil {
		tx.Rollback()
		return err
	}

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
			"SELECT token, created_at, live_house_staff_account_id FROM live_house_staff_account_provisional_registrations WHERE live_house_staff_account_id = ?",
			id,
		)

		var (
			token     string
			createdAt []uint8
			_id       string
		)
		if err := row.Scan(&token, &createdAt, &_id); err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
			return nil, nil
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

func (repo LiveHouseStaffAccountRepositoryImpl) FindById(
	id live_house_staff_account_domain.LiveHouseStaffAccountId,
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
		"SELECT id, email, is_provisional FROM live_house_staff_accounts WHERE id = ?",
		id.String(),
	)

	var (
		_id           string
		email         string
		isProvisional int
	)

	if err := row.Scan(&_id, &email, &isProvisional); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	credentialChallenges, err := repo.findChallenges(id, ctx)
	if err != nil {
		return nil, err
	}

	if isProvisional == 1 {
		row := repo.db.QueryRowContext(
			ctx,
			"SELECT token, created_at, live_house_staff_account_id FROM live_house_staff_account_provisional_registrations WHERE live_house_staff_account_id = ?",
			id.String(),
		)

		var (
			token     string
			createdAt []uint8
			_id       string
		)
		if err := row.Scan(&token, &createdAt, &_id); err != nil {
			if err != sql.ErrNoRows {
				return nil, err
			}
			return nil, nil
		}

		createdAtTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}

		account, err := live_house_staff_account_domain.NewLiveHouseStaffAccount(live_house_staff_account_domain.NewLiveHouseStaffAccountParams{
			Id:            id.String(),
			Email:         email,
			IsProvisional: int2Bool(isProvisional),
			ProvisionalRegistration: &live_house_staff_account_domain.ProvisionalRegistrationParam{
				Token:     token,
				CreatedAt: createdAtTime,
			},
		})
		if err != nil {
			return nil, err
		}

		for _, v := range credentialChallenges {
			account.AddCredentialChallenge(v)
		}

		return account, nil
	}

	account, err := live_house_staff_account_domain.NewLiveHouseStaffAccount(live_house_staff_account_domain.NewLiveHouseStaffAccountParams{
		Id:            id.String(),
		Email:         email,
		IsProvisional: int2Bool(isProvisional),
	})
	if err != nil {
		return nil, err
	}

	for _, v := range credentialChallenges {
		account.AddCredentialChallenge(v)
	}

	return account, nil
}

func (repo LiveHouseStaffAccountRepositoryImpl) FindByProvisionalRegistrationToken(
	token live_house_staff_account_domain.Token,
	ctx context.Context,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	row := repo.db.QueryRowContext(
		ctx,
		`
		SELECT 
			live_house_staff_accounts.id,
			live_house_staff_accounts.email,
			live_house_staff_accounts.is_provisional,
			live_house_staff_account_provisional_registrations.token,
			live_house_staff_account_provisional_registrations.created_at
		FROM live_house_staff_account_provisional_registrations
		INNER JOIN live_house_staff_accounts ON live_house_staff_accounts.id = live_house_staff_account_provisional_registrations.live_house_staff_account_id	
		WHERE live_house_staff_account_provisional_registrations.token = ?
		`,
		token.String(),
	)

	var (
		accountId     string
		email         string
		isProvisional int
		tokenStr      string
		createdAt     []uint8
	)
	if err := row.Scan(&accountId, &email, &isProvisional, &tokenStr, &createdAt); err != nil {
		if err == sql.ErrNoRows {
			return nil, nil
		}
		return nil, err
	}

	isProvisionalBool := func(i int) bool {
		if i == 1 {
			return true
		}
		return false
	}(isProvisional)
	createdAtTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
	if err != nil {
		return nil, err
	}

	return live_house_staff_account_domain.NewLiveHouseStaffAccount(live_house_staff_account_domain.NewLiveHouseStaffAccountParams{
		Id:            accountId,
		Email:         email,
		IsProvisional: isProvisionalBool,
		ProvisionalRegistration: &live_house_staff_account_domain.ProvisionalRegistrationParam{
			Token:     tokenStr,
			CreatedAt: createdAtTime,
		},
	})

}

func (repo LiveHouseStaffAccountRepositoryImpl) findChallenges(
	id live_house_staff_account_domain.LiveHouseStaffAccountId,
	ctx context.Context,
) ([]live_house_staff_account_domain.CredentialChallengeIntf, error) {

	rows, err := repo.db.QueryContext(
		ctx,
		"SELECT challenge, created_at FROM live_house_staff_account_credential_challenges WHERE live_house_staff_account_id = ? ORDER BY created_at ASC",
		id.String(),
	)
	if err != nil {
		return nil, err
	}

	challenges := []live_house_staff_account_domain.CredentialChallengeIntf{}
	for rows.Next() {

		var (
			challenge string
			createdAt []uint8
		)

		if err := rows.Scan(&challenge, &createdAt); err != nil {
			return nil, err
		}

		createdAtTime, err := time.Parse("2006-01-02 15:04:05", string(createdAt))
		if err != nil {
			return nil, err
		}

		c, err := live_house_staff_account_domain.NewCredentialChallenge(
			challenge,
			createdAtTime,
		)

		if err != nil {
			return nil, err
		}

		challenges = append(challenges, c)
	}

	return challenges, nil

}
