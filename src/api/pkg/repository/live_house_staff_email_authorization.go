package repository

import (
	"context"

	"github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_email_authorization_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type LiveHouseStaffEmailAuthorizationRepositoryIntf interface {
	usecase.LiveHouseStaffEmailAuthorizationRepositoryIntf
}

type liveHouseStaffEmailAuthorizationRepositoryImpl struct {
	db *db.MySQL
}

func NewLiveHouseStaffEmailAuthorizationRepository(db *db.MySQL) LiveHouseStaffEmailAuthorizationRepositoryIntf {
	return &liveHouseStaffEmailAuthorizationRepositoryImpl{
		db: db,
	}
}

func (repo liveHouseStaffEmailAuthorizationRepositoryImpl) Save(
	auth live_house_staff_email_authorization_domain.LiveHouseStaffEmailAuthorizationIntf,
	ctx context.Context,
) error {

	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO live_house_staff_email_authorizations (email, token) VALUES (?, ?)",
		auth.EmailAddress().String(),
		auth.Token().String(),
	)
	if err != nil {
		return err
	}

	return nil
}
