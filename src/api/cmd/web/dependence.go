package main

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/repository"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

func (app *application) setupDI() {

	//~~~~~~~~~~~~~~~~~~ repository ~~~~~~~~~~~~~~~~~~//

	app.container.Bind(func() live_house_staff_usecase.LiveHouseStaffRepositoryIntf {
		return repository.NewliveHouseStaff(app.db)
	})

	app.container.Bind(func() live_house_staff_domain.LiveHouseStaffRepositoryIntf {
		return repository.NewliveHouseStaff(app.db)
	})

	app.container.Bind(func() live_house_staff_usecase.UuidRepositoryIntf {
		return repository.NewUuidRepository()
	})
	// or
	// app.container.Bind(repository.NewUuidRepository)

	//~~~~~~~~~~~~~~~~~~ domain ~~~~~~~~~~~~~~~~~~//

	app.container.Bind(func(liveHouseStaffRepository live_house_staff_domain.LiveHouseStaffRepositoryIntf) live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf {
		return live_house_staff_domain.NewLiveHouseStaffEmailAddressChecker(liveHouseStaffRepository)
	})

	app.container.Bind(func(
		uuidRepository live_house_staff_usecase.UuidRepositoryIntf,
		liveHouseStaffRepository live_house_staff_usecase.LiveHouseStaffRepositoryIntf,
		liveHouseStaffEmailAddressChecker live_house_staff_domain.LiveHouseStaffEmailAddressCheckerIntf,
	) live_house_staff_usecase.LiveHouseStaffUsecaseIntf {
		return live_house_staff_usecase.NewLiveHouseStaffUsecase(
			uuidRepository,
			liveHouseStaffRepository,
			liveHouseStaffEmailAddressChecker,
		)
	})
}
