package web

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
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

func (ctrl liveHouseStaffController) RegisterAccount(c echo.Context) error {

	var usecase live_house_staff_usecase.LiveHouseStaffUsecaseIntf
	ctrl.container.Make(&usecase)

	id, err := usecase.RegisterAccount("name", "emailAddress@test.com", "password", c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}
