package server

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"go.uber.org/fx"
)

func CreateServer() *echo.Echo {
	e := echo.New()
	e.Use(middleware.CORS())
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.RequestID())

	return e
}

func StartServer(e *echo.Echo, lifecycle fx.Lifecycle, shutdowner fx.Shutdowner) {
	lifecycle.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go func() {
				e.Start(":8080")
				shutdowner.Shutdown()
			}()

			return nil
		},
		OnStop: func(ctx context.Context) error {

			return e.Shutdown(ctx)
		},
	})
}
