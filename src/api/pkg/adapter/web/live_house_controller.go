package web

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase"
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

	id, err := usecase.RegisterLiveHouseStaff("name", "emailAddress@test.com", "password", c.Request().Context(), ctrl.container)
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}

// curl -X POST -H "Content-Type: application/json" -d '{}' localhost/api/v1/live_house_account
func (ctrl liveHouseStaffController) RegisterAccount(c echo.Context) error {

	id, err := usecase.RegisterLiveHouseAccount("298e12d6-ec49-4dd7-8a39-84b090d47b36", c.Request().Context(), ctrl.container)
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}
