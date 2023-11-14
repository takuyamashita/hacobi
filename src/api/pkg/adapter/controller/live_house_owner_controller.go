package controller

import (
	"github.com/labstack/echo/v4"
	livehouseowner "github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_owner"
)

type LiveHouseOwnerController struct {
	accountUseCase livehouseowner.AccountUseCase
}

func NewLiveHouseOwnerController(
	accountUseCase livehouseowner.AccountUseCase,
) LiveHouseOwnerController {
	return LiveHouseOwnerController{
		accountUseCase: accountUseCase,
	}
}

func (ctrl LiveHouseOwnerController) RegisterAccount(c echo.Context) error {

	id, err := ctrl.accountUseCase.RegisterAccount("name", "emailAddress@test.com", "password", c.Request().Context())
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}
