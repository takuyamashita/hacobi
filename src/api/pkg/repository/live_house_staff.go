package repository

import (
	"context"
	"database/sql"
	"log"

	"github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type _compositIntf interface {
	usecase.LiveHouseStaffRepositoryIntf
	live_house_staff_domain.LiveHouseStaffRepositoryIntf
}

var _ _compositIntf = (*LiveHouseStaff)(nil)

type LiveHouseStaff struct {
	db *db.MySQL
}

func NewliveHouseStaff(db *db.MySQL) *LiveHouseStaff {
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

/*
mysql> desc live_house_staffs;
+------------+--------------+------+-----+----------------------+--------------------------------------------------+
| Field      | Type         | Null | Key | Default              | Extra                                            |
+------------+--------------+------+-----+----------------------+--------------------------------------------------+
| id         | varchar(36)  | NO   | PRI | NULL                 |                                                  |
| name       | varchar(255) | NO   |     | NULL                 |                                                  |
| email      | varchar(255) | NO   | UNI | NULL                 |                                                  |
| password   | varchar(255) | NO   |     | NULL                 |                                                  |
| created_at | datetime(6)  | NO   |     | CURRENT_TIMESTAMP(6) | DEFAULT_GENERATED                                |
| updated_at | datetime(6)  | NO   |     | CURRENT_TIMESTAMP(6) | DEFAULT_GENERATED on update CURRENT_TIMESTAMP(6) |
+------------+--------------+------+-----+----------------------+--------------------------------------------------+
6 rows in set (0.01 sec)
*/
func (repo LiveHouseStaff) FindById(id live_house_staff_domain.LiveHouseStaffId) (live_house_staff_domain.LiveHouseStaffIntf, error) {

	var dbId string
	var name string
	var email string
	var password string

	rows := repo.db.QueryRow(
		"SELECT id, name, email, password FROM live_house_staffs WHERE id = ?",
		id.String(),
	)
	err := rows.Scan(&dbId, &name, &email, &password)
	if err != nil {
		return nil, err
	}

	liveHouseStaff, err := live_house_staff_domain.NewLiveHouseStaff(
		dbId,
		name,
		email,
		password,
	)
	if err != nil {
		return nil, err
	}

	return liveHouseStaff, nil
}
