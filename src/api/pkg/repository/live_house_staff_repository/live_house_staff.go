package live_house_staff_repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

type _compositIntf interface {
	live_house_staff_usecase.LiveHouseStaffRepository
	live_house_staff_domain.LiveHouseStaffRepository
}

var _ _compositIntf = (*LiveHouseStaff)(nil)

type LiveHouseStaff struct {
	db *sql.DB
}

func NewliveHouseStaff(db *sql.DB) *LiveHouseStaff {
	return &LiveHouseStaff{
		db: db,
	}
}

func (repo LiveHouseStaff) Save(staff live_house_staff_domain.LiveHouseStaffIntf, ctx context.Context) error {

	_, err := repo.db.ExecContext(
		ctx,
		"INSERT INTO live_house_staffs (id, name, email, password) VALUES (?, ?, ?, ?)",
		staff.Id().String(), staff.Name().String(), staff.EmailAddress().String(), staff.Password().String(),
	)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}

func (repo LiveHouseStaff) FindByEmail(emailAddress live_house_staff_domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_domain.LiveHouseStaffIntf, error) {

	var id string
	var name string
	var email string
	var password string

	err := repo.db.QueryRowContext(
		ctx,
		"SELECT id, name, email, password FROM live_house_staffs WHERE email = ?",
		emailAddress.String(),
	).Scan(&id, &name, &email, &password)

	if err != sql.ErrNoRows {
		return nil, err
	}

	if err == sql.ErrNoRows {
		return nil, nil
	}

	liveHouseStaff, err := live_house_staff_domain.NewLiveHouseStaff(
		id,
		name,
		email,
		password,
	)
	if err != nil {
		return nil, err
	}

	return liveHouseStaff, nil
}
