package main

import (
	"context"
	"database/sql"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/takuyamashita/hacobi/src/api/db"
	"github.com/takuyamashita/hacobi/src/api/pkg/adapter/web"
	"github.com/takuyamashita/hacobi/src/api/pkg/domain/live_house_staff_domain"
	"github.com/takuyamashita/hacobi/src/api/pkg/repository"
	"github.com/takuyamashita/hacobi/src/api/pkg/usecase/live_house_staff_usecase"
)

type Application interface {
	start()
}

type application struct {
	server *echo.Echo
	db     *sql.DB
}

func newApplication() Application {

	ctx := context.Background()
	db := db.NewDatabase(ctx)

	app := &application{
		server: echo.New(),
		db:     db,
	}
	app.setupMiddlewares()
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

	uuid_repository := repository.NewUuidRepository()
	liveHouseStaffRepository := repository.NewliveHouseStaff(app.db)

	liveHouseStaffEmailAddressChecker := live_house_staff_domain.NewLiveHouseStaffEmailAddressChecker(liveHouseStaffRepository)

	liveHouseStaffUsecase := live_house_staff_usecase.NewLiveHouseStaffUsecase(uuid_repository, liveHouseStaffRepository, liveHouseStaffEmailAddressChecker)

	liveHouseStaffController := web.NewliveHouseStaffController(liveHouseStaffUsecase)

	app.server.POST("/api/live_house_owner/account", liveHouseStaffController.RegisterAccount)
}
