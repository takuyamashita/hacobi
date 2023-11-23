package dependency

import (
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/event"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_email_authorization_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/repository"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
)

func SetupDI(container container.Container, db *db.MySQL) {

	//~~~~~~~~~~~~~~~~~~ repository ~~~~~~~~~~~~~~~~~~//

	container.Bind(func() usecase.UuidRepositoryIntf {
		return repository.NewUuidRepository()
	})

	container.Bind(func() domain.RandomStringRepositoryIntf {
		return repository.NewRandomStringRepository()
	})

	container.Bind(func() usecase.LiveHouseStaffEmailAuthorizationRepositoryIntf {
		return repository.NewLiveHouseStaffEmailAuthorizationRepository(db)
	})

	container.Bind(func() usecase.TransationRepositoryIntf {
		return repository.NewTransaction(db)
	})

	container.Bind(func() live_house_staff_domain.LiveHouseStaffRepositoryIntf {
		return repository.NewliveHouseStaff(db)
	})

	container.Bind(func() usecase.LiveHouseStaffRepositoryIntf {
		return repository.NewliveHouseStaff(db)
	})

	container.Bind(func() usecase.LiveHouseAccountRepositoryIntf {
		return repository.NewliveHouseAccount(db)
	})

	//~~~~~~~~~~~~~~~~~~ domain ~~~~~~~~~~~~~~~~~~//

	container.Bind(live_house_staff_domain.NewLiveHouseStaffEmailAddressChecker)

	container.Bind(live_house_staff_email_authorization_domain.NewTokenGenerator)

	/*
		var (
			rndStrRepo                  domain.RandomStringRepositoryIntf
			liveHouseStaffEmailAuthRepo usecase.LiveHouseStaffEmailAuthorizationRepositoryIntf
			tknGen                      live_house_staff_email_authorization_domain.TokenGeneratorIntf
		)
		container.Make(&rndStrRepo)
		container.Make(&liveHouseStaffEmailAuthRepo)
		container.Make(&tknGen)
	*/

	//~~~~~~~~~~~~~~~~~~ event ~~~~~~~~~~~~~~~~~~//

	container.Bind(event.NewLiveHouseStaffEmailAuthorizationCreated)
}
