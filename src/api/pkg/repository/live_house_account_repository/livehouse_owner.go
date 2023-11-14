package live_house_account_repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_account_domain"
)

type LiveHouseStaff struct {
	db *sql.DB
}

func NewliveHouseStaff(db *sql.DB) *LiveHouseStaff {
	return &LiveHouseStaff{
		db: db,
	}
}

func (repo LiveHouseStaff) Save(owner live_house_account_domain.LiveHouseStaff, ctx context.Context) (*live_house_account_domain.LiveHouseStaffId, error) {

	result, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO live_house_owners (name, email, password) VALUES (?, ?, ?)",
		owner.Name().String(), owner.EmailAddress().String(), owner.Password().String(),
	)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	dbId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	ownerId, err := live_house_account_domain.NewliveHouseStaffId(uint64(dbId))
	if err != nil {
		return nil, err
	}

	return &ownerId, nil
}
