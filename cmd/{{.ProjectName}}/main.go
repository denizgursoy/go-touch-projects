package main

import (
	"go.uber.org/fx"
	"go.uber.org/zap"
	"{{.ModuleName}}/internal/handlers"
	"{{.ModuleName}}/internal/repositories"
	"{{.ModuleName}}/internal/server"
	"{{.ModuleName}}/internal/services"
)

func main() {
	fx.New(
		fx.RecoverFromPanics(),
		fx.Provide(
			fx.Annotate(services.NewResourceService, fx.As(new(handlers.ResourceService))),
			fx.Annotate(repositories.NewResourceRepository, fx.As(new(services.ResourceRepository))),
			server.CreateServer,
			zap.NewProduction,
		),
		fx.Invoke(
			server.StartServer,
			handlers.NewResourceHandler,
		),
	).Run()
}
