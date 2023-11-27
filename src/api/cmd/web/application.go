package main

import (
	"context"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/takuyamashita/hacobi/src/api/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/adapter/web"
	"github.com/takuyamashita/hacobi/src/api/pkg/container"
	mysql "github.com/takuyamashita/hacobi/src/api/pkg/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/dependency"

	"github.com/go-webauthn/webauthn/protocol"
	"github.com/go-webauthn/webauthn/protocol/webauthncose"
)

type Application interface {
	start()
}

type application struct {
	server    *echo.Echo
	db        *mysql.MySQL
	container container.Container
}

func newApplication() Application {

	ctx := context.Background()
	db := mysql.NewMySQL(db.NewDatabase(ctx))

	app := &application{
		server:    echo.New(),
		db:        &db,
		container: container.NewContainer(),
	}
	app.setupMiddlewares()
	dependency.SetupDI(app.container, app.db)
	app.setupRoutes()

	return app
}

func (app *application) start() {

	defer app.db.Close()

	ctx, _ := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)

	go func() {
		if err := app.server.Start(":8080"); err != nil {
			switch err {
			case http.ErrServerClosed:
				app.server.Logger.Info("server shutdown")
			default:
				app.server.Logger.Fatal(err)
			}
		}
	}()

	<-ctx.Done()

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := app.server.Shutdown(ctx); err != nil {
		app.server.Logger.Fatal(err)
	}
}

func (app *application) setupMiddlewares() {
	app.server.Use(middleware.Logger())
	app.server.Use(middleware.Recover())
}

func (app *application) setupRoutes() {
	app.server.GET("/api/hello", func(c echo.Context) error {
		return c.String(200, "Hello, World!")
	})

	liveHouseStaffController := web.NewliveHouseStaffController(app.container)

	app.server.POST("/api/v1/send_live_house_staff_email_authorization", liveHouseStaffController.SendLiveHouseStaffRegisterMail)
	app.server.POST("/api/v1/live_house_staff", liveHouseStaffController.RegisterStaff)
	app.server.POST("/api/v1/live_house_account", liveHouseStaffController.RegisterAccount)

	app.server.POST("/api/v1/ceremony/start", func(c echo.Context) error {

		challenge, err := protocol.CreateChallenge()
		if err != nil {
			return c.JSON(500, err)
		}

		c.Request().AddCookie(&http.Cookie{
			Name:  "challenge",
			Value: challenge.String(),
		})

		option := protocol.PublicKeyCredentialCreationOptions{
			Challenge: challenge,
			RelyingParty: protocol.RelyingPartyEntity{
				CredentialEntity: protocol.CredentialEntity{
					Name: "localhost",
					Icon: "https://localhost/favicon.ico",
				},
			},
			User: protocol.UserEntity{
				ID:          []byte("1234567890"),
				DisplayName: "test-user",
			},
			CredentialExcludeList: []protocol.CredentialDescriptor{},
			Parameters: []protocol.CredentialParameter{
				{
					Type:      protocol.PublicKeyCredentialType,
					Algorithm: webauthncose.AlgES256,
				},
				{
					Type:      protocol.PublicKeyCredentialType,
					Algorithm: webauthncose.AlgRS256,
				},
			},
			Timeout: int((5 * time.Minute).Milliseconds()),
		}

		return c.JSON(200, option)
	})

	app.server.POST("/api/v1/auth", func(c echo.Context) error {

		parsedResponse, err := protocol.ParseCredentialCreationResponseBody(c.Request().Body)
		if err != nil {
			log.Println(err)
			return c.JSON(500, err)
		}

		return c.JSON(200, parsedResponse)
	})
}
