package main

import (
	"github.com/denizgursoy/go-touch-projects/internal/handlers"
	"github.com/denizgursoy/go-touch-projects/internal/repositories"
	"github.com/denizgursoy/go-touch-projects/internal/server"
	"github.com/denizgursoy/go-touch-projects/internal/services"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.RecoverFromPanics(),
		fx.Provide(
			fx.Annotate(services.NewResourceService, fx.As(new(handlers.ResourceService))),
			repositories.NewResourceRepository,
			server.CreateServer,
			zap.NewProduction,
		),
		fx.Invoke(
			server.StartServer,
			handlers.NewResourceHandler,
		),
	).Run()
}
