package live_house_staff_usecase

import (
	"context"
	"testing"

	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
)

type UuidRepositoryMock struct{}

func (u UuidRepositoryMock) Generate() (string, error) {
	return "uuid", nil
}

type LiveHouseStaffRepositoryMock struct{}

func (l LiveHouseStaffRepositoryMock) Save(owner live_house_staff_domain.LiveHouseStaff, ctx context.Context) error {
	return nil
}

func (l LiveHouseStaffRepositoryMock) FindByEmail(emailAddress live_house_staff_domain.LiveHouseStaffEmailAddress, ctx context.Context) (live_house_staff_domain.LiveHouseStaff, error) {
	return nil, nil
}

func TestRegisterAccount(t *testing.T) {

	uuidRepositoryMock := UuidRepositoryMock{}
	liveHouseStaffRepositoryMock := LiveHouseStaffRepositoryMock{}

	usecase := NewLiveHouseStaffUsecase(uuidRepositoryMock, liveHouseStaffRepositoryMock)
	usecase.RegisterAccount("name", "email", "password", context.Background())

	t.Error("test")
}
