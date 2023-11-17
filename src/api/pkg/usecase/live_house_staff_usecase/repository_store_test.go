package live_house_staff_usecase_test

import (
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

type Store struct {
	Staffs         []live_house_staff_domain.LiveHouseStaffIntf
	uuidRepository live_house_staff_usecase.UuidRepository
}

func NewStore(uuidRepository live_house_staff_usecase.UuidRepository) Store {
	return Store{
		uuidRepository: uuidRepository,
	}
}

var staffs = []map[string]string{
	{
		"id":           "",
		"name":         "name1",
		"emailAddress": "test-1@test.com",
		"password":     "password",
	},
	{
		"id":           "",
		"name":         "name2",
		"emailAddress": "test-2@test.com",
		"password":     "password",
	},
	{
		"id":           "",
		"name":         "name3",
		"emailAddress": "duplicate@test.com",
		"password":     "password",
	},
}

func (s *Store) Clear() {
	s.Staffs = []live_house_staff_domain.LiveHouseStaffIntf{}
}

func (s *Store) SetupStore(uuidRepository live_house_staff_usecase.UuidRepository, t *testing.T) {
	s.setupStaffs(uuidRepository, t)
}

func (s *Store) setupStaffs(uuidRepository live_house_staff_usecase.UuidRepository, t *testing.T) {

	if s.Staffs == nil {
		s.Staffs = []live_house_staff_domain.LiveHouseStaffIntf{}
	}

	for _, staff := range staffs {

		id, err := uuidRepository.Generate()
		if err != nil {
			t.Fatal(err)
		}

		staff, err := live_house_staff_domain.NewLiveHouseStaff(id, staff["name"], staff["emailAddress"], staff["password"])
		if err != nil {
			t.Fatal(err)
		}

		s.Staffs = append(s.Staffs, staff)
	}

}
