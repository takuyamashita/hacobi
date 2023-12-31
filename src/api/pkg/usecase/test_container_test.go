package usecase_test

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/dependency"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_account_domain"
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

	container.Bind(func() domain.MailIntf {
		return &MailMock{
			Store: store,
		}
	})

	container.Bind(func() usecase.TransationRepositoryIntf {
		return &TransactionRepositoryMock{}
	})

	container.Bind(func() live_house_staff_domain.LiveHouseStaffRepositoryIntf {
		return &LiveHouseStaffRepositoryMock{Store: store}
	})

	container.Bind(func() usecase.LiveHouseStaffAccountRepositoryIntf {
		return &LiveHouseStaffAccountRepositoryMock{Store: store}
	})

	container.Bind(func() live_house_staff_account_domain.LiveHouseStaffAccountRepositoryIntf {
		return &LiveHouseStaffAccountRepositoryMock{Store: store}
	})

	container.Bind(func() usecase.LiveHouseStaffRepositoryIntf {
		return &LiveHouseStaffRepositoryMock{Store: store}
	})

	return &container
}
