package live_house_staff_usecase

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type UuidRepositoryMock struct{}

func (u UuidRepositoryMock) Generate() (string, error) {
	return "uuid", nil
}

/*
type LiveHouseStaffRepositoryMock struct{}

func (l LiveHouseStaffRepositoryMock) Save(owner live_house_staff_domain.LiveHouseStaffIntf, ctx context.Context) error {
	return nil
}

func (l LiveHouseStaffRepositoryMock) FindByEmail(emailAddress live_house_staff_domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_domain.LiveHouseStaffIntf, error) {
	return nil, nil
}
*/

func TestRegisterAccount(t *testing.T) {

	store := Store{
		Users: map[string]live_house_staff_domain.LiveHouseStaffIntf{},
	}
	for i := 0; i < 1; i++ {
		staff, err := live_house_staff_domain.NewLiveHouseStaff(fmt.Sprintf("%d", i), "name", fmt.Sprintf("%dtest@test.com", i), "password")
		if err != nil {
			t.Fatal(err)
		}
		store.Users[fmt.Sprintf("%d", i)] = staff
	}

	uuidRepositoryMock := UuidRepositoryMock{}
	liveHouseStaffRepositoryMock := LiveHouseStaffRepositoryMock{
		Store: store,
	}

	liveHouseStaffEmailAddressChecker := live_house_staff_domain.NewLiveHouseStaffEmailAddressChecker(liveHouseStaffRepositoryMock)

	usecase := NewLiveHouseStaffUsecase(uuidRepositoryMock, liveHouseStaffRepositoryMock, liveHouseStaffEmailAddressChecker)
	id, err := usecase.RegisterAccount("name", "2test@test.com", "password", context.Background())
	if err != nil {
		t.Error(err)
	}

	log.Println(id)
}
