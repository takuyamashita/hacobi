package livehouseownergo

import (
	"context"
	"database/sql"

	livehouseownerdomain "github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_owner"
)

type liveHouseOwner struct {
	db *sql.DB
}

func NewLiveHouseOwner(db *sql.DB) *liveHouseOwner {
	return &liveHouseOwner{
		db: db,
	}
}

func (repo liveHouseOwner) Save(owner livehouseownerdomain.LiveHouseOwner, ctx context.Context) (*livehouseownerdomain.LiveHouseOwnerId, error) {

	result, err := repo.db.ExecContext(ctx, "INSERT INTO live_house_owners (name, email_address, password) VALUES (?, ?, ?)", owner.Name(), owner.EmailAddress(), owner.Password())
	if err != nil {
		return nil, err
	}

	dbId, err := result.LastInsertId()
	if err != nil {
		return nil, err
	}

	ownerId, err := livehouseownerdomain.NewLiveHouseOwnerId(uint64(dbId))
	if err != nil {
		return nil, err
	}

	return &ownerId, nil
}
