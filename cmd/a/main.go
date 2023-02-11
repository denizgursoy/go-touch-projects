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
		fx.Provide(server.CreateServer),
		fx.Provide(services.NewResourceService),
		fx.Provide(repositories.NewResourceRepository),
		fx.Provide(zap.NewProduction),
		fx.Invoke(server.StartServer),
		fx.Invoke(handlers.NewResourceHandler),
	).Run()
}
