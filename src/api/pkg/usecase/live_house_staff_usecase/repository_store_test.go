package live_house_staff_usecase_test

import (
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

type Store struct {
	Staffs         []live_house_staff_domain.LiveHouseStaffIntf
	uuidRepository usecase.UuidRepositoryIntf
}

func NewStore(uuidRepository usecase.UuidRepositoryIntf) Store {
	return Store{
		uuidRepository: uuidRepository,
	}
}

func (s *Store) Clear() {
	s.Staffs = []live_house_staff_domain.LiveHouseStaffIntf{}
}

func (s *Store) SetupStore(uuidRepository usecase.UuidRepositoryIntf, t *testing.T) {
	s.setupStaffs(uuidRepository, t)
}

func (s *Store) setupStaffs(uuidRepository usecase.UuidRepositoryIntf, t *testing.T) {

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
