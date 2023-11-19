package web

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_account_usecase"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

type liveHouseStaffController struct {
	container container.Container
}

func NewliveHouseStaffController(container container.Container) liveHouseStaffController {
	return liveHouseStaffController{
		container: container,
	}
}

// curl -X POST -H "Content-Type: application/json" -d '{}' localhost/api/v1/live_house_staff
func (ctrl liveHouseStaffController) RegisterStaff(c echo.Context) error {

	var usecase live_house_staff_usecase.LiveHouseStaffUsecaseIntf
	ctrl.container.Make(&usecase)

	id, err := usecase.RegisterAccount("name", "emailAddress@test.com", "password", c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}

// curl -X POST -H "Content-Type: application/json" -d '{}' localhost/api/v1/live_house_account
func (ctrl liveHouseStaffController) RegisterAccount(c echo.Context) error {

	var usecase live_house_account_usecase.LiveHouseAccountUsecaseIntf
	ctrl.container.Make(&usecase)

	id, err := usecase.RegisterLiveHouseAccount("16d594eb-91d8-4aa1-a1c9-c021efb07cc2", c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}
