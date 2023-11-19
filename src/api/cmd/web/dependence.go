package main

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/repository"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_account_usecase"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

func (app *application) setupDI() {

	//~~~~~~~~~~~~~~~~~~ repository ~~~~~~~~~~~~~~~~~~//

	app.container.Bind(func() usecase.UuidRepositoryIntf {
		return repository.NewUuidRepository()
	})

	app.container.Bind(func() live_house_staff_domain.LiveHouseStaffRepositoryIntf {
		return repository.NewliveHouseStaff(app.db)
	})

	app.container.Bind(func() usecase.LiveHouseStaffRepositoryIntf {
		return repository.NewliveHouseStaff(app.db)
	})

	app.container.Bind(func() usecase.LiveHouseAccountRepositoryIntf {
		return repository.NewliveHouseAccount(app.db)
	})

	//~~~~~~~~~~~~~~~~~~ domain ~~~~~~~~~~~~~~~~~~//

	app.container.Bind(func(liveHouseStaffRepository live_house_staff_domain.LiveHouseStaffRepositoryIntf) live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf {
		return live_house_staff_domain.NewLiveHouseStaffEmailAddressChecker(liveHouseStaffRepository)
	})

	//~~~~~~~~~~~~~~~~~~ usecase ~~~~~~~~~~~~~~~~~~//

	app.container.Bind(live_house_staff_usecase.NewLiveHouseStaffUsecase)
	app.container.Bind(live_house_account_usecase.NewLiveHouseAccountUsecase)
}
