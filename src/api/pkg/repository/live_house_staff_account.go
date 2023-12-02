package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"log"

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

func (repo LiveHouseStaffAccountRepositoryImpl) FindByEmail(
	emailAddress domain.LiveHouseStaffEmailAddress,
	ctx context.Context,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	params := newFindByEmailParams(emailAddress)

	return repo.findBy(params, ctx)
}

func (repo LiveHouseStaffAccountRepositoryImpl) FindById(
	id live_house_staff_account_domain.LiveHouseStaffAccountId,
	ctx context.Context,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	params := newFindByIdParams(id)

	return repo.findBy(params, ctx)
}

func (repo LiveHouseStaffAccountRepositoryImpl) FindByProvisionalRegistrationToken(
	token live_house_staff_account_domain.Token,
	ctx context.Context,
) (live_house_staff_account_domain.LiveHouseStaffAccountIntf, error) {

	params := newFindByProvisionalRegistrationTokenParams(token)

	return repo.findBy(params, ctx)
}

type findByType string

func (t findByType) stmt() string {

	switch t {
	case findByTypeId:
		return fmt.Sprintf(selectStmtTmpl, "a", "id", "?")
	case findByTypeEmail:
		return fmt.Sprintf(selectStmtTmpl, "a", "email", "?")
	case findByTypeToken:
		return fmt.Sprintf(selectStmtTmpl, "r", "token", "?")
	default:
		return ""
	}
}

const (
	findByTypeEmail findByType = "email"
	findByTypeId    findByType = "id"
	findByTypeToken findByType = "token"

	selectStmtTmpl = `
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
		WHERE %s.%s = %s
	`
)

type findByParams struct {
	findByType  findByType
	id          live_house_staff_account_domain.LiveHouseStaffAccountId
	email       domain.LiveHouseStaffEmailAddress
	registToken live_house_staff_account_domain.Token
}

func newFindByIdParams(id live_house_staff_account_domain.LiveHouseStaffAccountId) findByParams {
	return findByParams{
		findByType: findByTypeId,
		id:         id,
	}
}

func newFindByEmailParams(email domain.LiveHouseStaffEmailAddress) findByParams {
	return findByParams{
		findByType: findByTypeEmail,
		email:      email,
	}
}

func newFindByProvisionalRegistrationTokenParams(token live_house_staff_account_domain.Token) findByParams {
	return findByParams{
		findByType:  findByTypeToken,
		registToken: token,
	}
}

func (p findByParams) args() []interface{} {

	switch p.findByType {
	case findByTypeId:
		return []interface{}{p.id.String()}
	case findByTypeEmail:
		return []interface{}{p.email.String()}
	case findByTypeToken:
		return []interface{}{p.registToken.String()}
	default:
		return []interface{}{}
	}
}

func int2Bool(i int) bool {
	if i == 1 {
		return true
	}
	return false
}

func (repo LiveHouseStaffAccountRepositoryImpl) findBy(params findByParams, ctx context.Context) (account live_house_staff_account_domain.LiveHouseStaffAccountIntf, err error) {

	log.Println(params.findByType.stmt())

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
		params.findByType.stmt(),
		params.args()...,
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

	account, err = live_house_staff_account_domain.NewLiveHouseStaffAccount(live_house_staff_account_domain.NewLiveHouseStaffAccountParams{
		Id:                        accountId,
		Email:                     email,
		IsProvisional:             int2Bool(isProvisional),
		ProvisionalRegistration:   provisionalRegistrationParam,
		CredentialChallengeParams: credentialChallengeParam,
	})
	if err != nil {
		return nil, err
	}
	return account, nil
}
