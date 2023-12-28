package main

import (
	"context"
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
	// xxx: リクエストスコープ毎にDBをわたさないと、リクエストを超えてトランザクションを貼れてしまう
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

	app.server.POST("/api/v1/send_live_house_staff_email_authorization", liveHouseStaffController.SendMailLiveHouseStaffAccountRegisterPage)
	app.server.POST("/api/v1/live_house_account/credential/start_register", liveHouseStaffController.StartRegister)
	app.server.POST("/api/v1/live_house_account/credential/finish_register", liveHouseStaffController.FinishRegister)
	app.server.POST("/api/v1/live_house_account/credential/start_login", liveHouseStaffController.StartLogin)
	app.server.POST("/api/v1/live_house_account/credential/finish_login", liveHouseStaffController.FinishLogin)

	liveHouseStaffAccountOnly := app.server.Group("/api/v1/live_house_staff_account/enable")

	liveHouseStaffAccountOnly.Use(middleware.CORS())
	liveHouseStaffAccountOnly.Use(web.NewAuthJwtMiddleware())

	liveHouseStaffAccountOnly.GET("/live_house_staff", liveHouseStaffController.GetStaff)
}
