package web

import (
	"log"
	"net/http"
	"time"

	"github.com/golang-jwt/jwt/v5"
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

// curl -X POST -H "Content-Type: application/json" -d '{}' localhost/api/v1/send_live_house_staff_email_authorization
func (ctrl liveHouseStaffController) SendMailLiveHouseStaffAccountRegisterPage(c echo.Context) error {

	req := struct {
		EmailAddress string `json:"emailAddress"`
	}{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	err := usecase.RegisterProvisionalLiveHouseStaffAccount(req.EmailAddress, c.Request().Context(), ctrl.container)
	if err != nil {
		c.Error(err)
		return err
	}

	return c.JSON(200, "ok")
}

type jwtClaims struct {
	AccountId string `json:"accountId"`
	jwt.RegisteredClaims
}

func (ctrl liveHouseStaffController) StartRegister(c echo.Context) error {

	req := struct {
		Token string `json:"token"`
	}{}

	if err := c.Bind(&req); err != nil {
		return err
	}

	option, accountId, err := usecase.StartRegister(req.Token, c.Request().Context(), ctrl.container)
	if err != nil {
		return err
	}

	claims := jwtClaims{
		accountId,
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Millisecond * time.Duration(option.Timeout))),
		},
	}
	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	c.SetCookie(&http.Cookie{
		Name:  "jwt",
		Value: jwtToken,
	})

	return c.JSON(200, option)
}

func (ctrl liveHouseStaffController) FinishRegister(c echo.Context) error {

	jwtCookie, err := c.Cookie("jwt")
	if err != nil {
		return err
	}
	// get accountId from jwt
	claims := jwtClaims{}
	_, err = jwt.ParseWithClaims(jwtCookie.Value, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte("secret"), nil
	})
	if err != nil {
		return err
	}

	log.Println(claims.AccountId)

	if err := usecase.FinishRegisterLiveHouseStaffAccount(c.Request().Body, claims.AccountId, c.Request().Context(), ctrl.container); err != nil {
		return err
	}

	return c.JSON(200, "ok")
}

func (ctrl liveHouseStaffController) StartLogin(c echo.Context) error {

	req := struct {
		EmailAddress string `json:"emailAddress"`
	}{}

	if err := c.Bind(&req); err != nil {
		return err
	}
	log.Println(req.EmailAddress)
	option, err := usecase.StartLiveHouseStaffAccountLogin(req.EmailAddress, c.Request().Context(), ctrl.container)
	if err != nil {
		return err
	}

	return c.JSON(200, option)
}

func (ctrl liveHouseStaffController) FinishLogin(c echo.Context) error {

	err, account := usecase.LoginLiveHouseStaffAccount(c.Request().Body, c.Request().Context(), ctrl.container)
	if err != nil {
		return err
	}

	claims := AuthJwtClaims{
		AccountId: account.Id().String(),
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * 1)),
		},
	}

	// xxx: secretは別で定義する
	jwtToken, err := jwt.NewWithClaims(jwt.SigningMethodHS256, claims).SignedString([]byte("secret"))
	if err != nil {
		return err
	}

	c.SetCookie(&http.Cookie{
		Name:     JwtTokenCookieName,
		Value:    jwtToken,
		HttpOnly: true,
		Path:     "/",
		Expires:  time.Now().Add(time.Hour * 500),
	})

	return c.JSON(200, "ok")
}

// curl -X POST -H "Content-Type: application/json" -d '{}' localhost/api/v1/live_house_account
func (ctrl liveHouseStaffController) RegisterAccount(c echo.Context) error {

	id, err := usecase.RegisterLiveHouseAccount("298e12d6-ec49-4dd7-8a39-84b090d47b36", c.Request().Context(), ctrl.container)
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}

// curl -X POST -H "Content-Type: application/json" -d '{}' localhost/api/v1/live_house_staff
func (ctrl liveHouseStaffController) RegisterStaff(c echo.Context) error {

	id, err := usecase.RegisterLiveHouseStaff("name", "emailAddress@test.com", "password", c.Request().Context(), ctrl.container)
	if err != nil {
		return err
	}

	return c.JSON(200, id)
}

func (ctrl liveHouseStaffController) GetStaff(c echo.Context) error {

	jwt := c.Get(AuthJwtKey).(*jwt.Token)
	auth := jwt.Claims.(*AuthJwtClaims)

	account, err := usecase.GetLiveHouseStaffAccount(auth.AccountId, c.Request().Context(), ctrl.container)

	if err != nil {
		return err
	}

	type Response struct {
		Id           string `json:"id"`
		Name         string `json:"name"`
		EmailAddress string `json:"emailAddress"`
	}

	return c.JSON(200, Response{
		Id:           account.Id().String(),
		Name:         string(account.Profile().DisplayName()),
		EmailAddress: account.EmailAddress().String(),
	})
}
