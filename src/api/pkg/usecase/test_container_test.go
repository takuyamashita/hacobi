package usecase_test

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/dependency"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/repository"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

func NewTestContainer(store *Store) container.Container {

	container := container.NewContainer()
	dependency.SetupDI(&container, nil)

	container.Bind(func() usecase.UuidRepositoryIntf {
		return repository.NewUuidRepository()
	})

	container.Bind(func() usecase.TransationRepositoryIntf {
		return &TransactionRepositoryMock{}
	})

	container.Bind(func() live_house_staff_domain.LiveHouseStaffRepositoryIntf {
		return &LiveHouseStaffRepositoryMock{Store: store}
	})

	container.Bind(func() usecase.LiveHouseStaffEmailAuthorizationRepositoryIntf {
		return &LiveHouseStaffEmailAuthorizationRepositoryMock{Store: store}
	})

	container.Bind(func() usecase.LiveHouseStaffRepositoryIntf {
		return &LiveHouseStaffRepositoryMock{Store: store}
	})

	container.Bind(func() usecase.LiveHouseAccountRepositoryIntf {
		return &LiveHouseAccountRepositoryMock{Store: store}
	})

	return &container
}
