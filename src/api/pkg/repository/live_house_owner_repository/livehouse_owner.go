package live_house_owner_repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_owner_domain"
)

type liveHouseOwner struct {
	db *sql.DB
}

func NewLiveHouseOwner(db *sql.DB) *liveHouseOwner {
	return &liveHouseOwner{
		db: db,
	}
}

func (repo liveHouseOwner) Save(owner live_house_owner_domain.LiveHouseOwner, ctx context.Context) (*live_house_owner_domain.LiveHouseOwnerId, error) {

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

	ownerId, err := live_house_owner_domain.NewLiveHouseOwnerId(uint64(dbId))
	if err != nil {
		return nil, err
	}

	return &ownerId, nil
}
