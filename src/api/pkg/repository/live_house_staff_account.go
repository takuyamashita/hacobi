package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

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

func int2Bool(i int) bool {
	return i == 1
}

type liveHouseStaffAccountSelectStmtTmpl string
type liveHouseStaffAccountSelectStmtTmplArgs func() []interface{}

const liveHouseStaffAccountSelectStmt liveHouseStaffAccountSelectStmtTmpl = `
		SELECT
			a.id,
			a.email,
			a.is_provisional,
			r.token,
			r.created_at,
			c.challenge,
			c.created_at
		FROM 
			live_house_staff_accounts
			AS
				a
		LEFT JOIN 
			live_house_staff_account_provisional_registrations 
			AS
				r
			ON
				a.id = r.live_house_staff_account_id
		LEFT JOIN
			live_house_staff_account_credential_challenges
			AS
				c
			ON
				a.id = c.live_house_staff_account_id	
		WHERE %s.%s = ?
`

func (stmt liveHouseStaffAccountSelectStmtTmpl) bind(
	tableName string,
	columnName string,
	value string,
) (liveHouseStaffAccountSelectStmtTmpl, liveHouseStaffAccountSelectStmtTmplArgs) {

	s := fmt.Sprintf(string(stmt), tableName, columnName)

	return liveHouseStaffAccountSelectStmtTmpl(s), func() []interface{} {
		return []interface{}{value}
	}
}

func (repo LiveHouseStaffAccountRepositoryImpl) findBy(
	stmt liveHouseStaffAccountSelectStmtTmpl,
	args liveHouseStaffAccountSelectStmtTmplArgs,
	ctx context.Context,
) (account live_house_staff_account_domain.LiveHouseStaffAccountIntf, err error) {

	var (
		accountId          string
		email              string
		isProvisional      int
		token              sql.NullString
		tokenCreatedAt     sql.NullTime
		challenge          sql.NullString
		challengeCreatedAt sql.NullTime
	)

	row := repo.db.QueryRowContext(
		ctx,
		string(stmt),
		args()...,
	)

	if err := row.Scan(
		&accountId,
		&email,
		&isProvisional,
		&token,
		&tokenCreatedAt,
		&challenge,
		&challengeCreatedAt,
	); err != nil {

		if errors.Is(err, sql.ErrNoRows) {
			return nil, nil
		}

		return nil, err
	}

	var (
		provisionalRegistrationParam *live_house_staff_account_domain.ProvisionalRegistrationParam
		credentialChallengeParam     *live_house_staff_account_domain.CredentialChallengeParam
	)

	if token.Valid {

		provisionalRegistrationParam = &live_house_staff_account_domain.ProvisionalRegistrationParam{
			Token:     token.String,
			CreatedAt: tokenCreatedAt.Time,
		}
	}

	if challenge.Valid {

		credentialChallengeParam = &live_house_staff_account_domain.CredentialChallengeParam{
			Challenge: challenge.String,
			CreatedAt: challengeCreatedAt.Time,
		}
	}

	var credentialKeys []domain.PublicKeyId = make([]domain.PublicKeyId, 0)
	rows, err := repo.db.QueryContext(
		ctx,
		`
			SELECT
				public_key_id
			FROM
				live_house_staff_account_credential_relations
			WHERE
				live_house_staff_account_id = ?
		`,
		accountId,
	)
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var credentialKey string
		if err := rows.Scan(&credentialKey); err != nil {
			return nil, err
		}
		credentialKeyId, err := domain.NewPublicKeyId(credentialKey)
		if err != nil {
			return nil, err
		}
		credentialKeys = append(credentialKeys, credentialKeyId)
	}

	account, err = live_house_staff_account_domain.NewLiveHouseStaffAccount(live_house_staff_account_domain.NewLiveHouseStaffAccountParams{
		Id:                        accountId,
		Email:                     email,
		IsProvisional:             int2Bool(isProvisional),
		ProvisionalRegistration:   provisionalRegistrationParam,
		CredentialChallengeParams: credentialChallengeParam,
		CredentialKeys:            credentialKeys,
	})
	if err != nil {
		return nil, err
	}
	return account, nil
}

func (repo LiveHouseStaffAccountRepositoryImpl) FindByEmail(
	emailAddress domain.LiveHouseStaffEmailAddress,
	ctx context.Context,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	stmt, args := liveHouseStaffAccountSelectStmt.bind("a", "email", emailAddress.String())

	return repo.findBy(stmt, args, ctx)
}

func (repo LiveHouseStaffAccountRepositoryImpl) FindById(
	id live_house_staff_account_domain.LiveHouseStaffAccountId,
	ctx context.Context,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	stmt, args := liveHouseStaffAccountSelectStmt.bind("a", "id", id.String())

	return repo.findBy(stmt, args, ctx)
}

func (repo LiveHouseStaffAccountRepositoryImpl) FindByProvisionalRegistrationToken(
	token live_house_staff_account_domain.Token,
	ctx context.Context,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	stmt, args := liveHouseStaffAccountSelectStmt.bind("r", "token", token.String())

	return repo.findBy(stmt, args, ctx)
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

	if account.CredentialChallenge() == nil {

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

	_, err = tx.ExecContext(
		ctx,
		`
			INSERT INTO live_house_staff_account_credential_challenges
				(live_house_staff_account_id, challenge, created_at)
			VALUES
				(?, ?, ?) AS new
			ON DUPLICATE KEY UPDATE
				challenge = new.challenge,
				created_at = new.created_at
		`,
		account.Id().String(),
		account.CredentialChallenge().Challenge().String(),
		account.CredentialChallenge().CreatedAt(),
	)
	if err != nil {
		tx.Rollback()
		return err
	}

	return tx.Commit()

}
