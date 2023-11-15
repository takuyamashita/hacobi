package web

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

type liveHouseStaffController struct {
	accountUseCase live_house_staff_usecase.LiveHouseStaffUsecase
}

func NewliveHouseStaffController(
	accountUseCase live_house_staff_usecase.LiveHouseStaffUsecase,
) liveHouseStaffController {
	return liveHouseStaffController{
		accountUseCase: accountUseCase,
	}
}

func (ctrl liveHouseStaffController) RegisterAccount(c echo.Context) error {

	id, err := ctrl.accountUseCase.RegisterAccount("name", "emailAddress@test.com", "password", c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}
