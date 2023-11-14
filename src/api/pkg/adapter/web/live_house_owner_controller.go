package web

import (
	"github.com/labstack/echo/v4"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_owner_usecase"
)

type LiveHouseOwnerController struct {
	accountUseCase live_house_owner_usecase.AccountUseCase
}

func NewLiveHouseOwnerController(
	accountUseCase live_house_owner_usecase.AccountUseCase,
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
