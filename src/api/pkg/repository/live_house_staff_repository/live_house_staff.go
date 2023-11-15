package live_house_staff_repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type LiveHouseStaff struct {
	db *sql.DB
}

func NewliveHouseStaff(db *sql.DB) *LiveHouseStaff {
	return &LiveHouseStaff{
		db: db,
	}
}

func (repo LiveHouseStaff) Save(owner live_house_staff_domain.LiveHouseStaff, ctx context.Context) error {

	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO live_house_staffs (id, name, email, password) VALUES (?, ?, ?, ?)",
		owner.Id().String(), owner.Name().String(), owner.EmailAddress().String(), owner.Password().String(),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
