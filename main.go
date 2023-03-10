package main

import (
	"management/db"
	"management/handler"

	"go.uber.org/fx"
	"go.uber.org/fx/fxevent"
	"go.uber.org/zap"
)

func main() {
	fx.New(
		fx.WithLogger(func(log *zap.Logger) fxevent.Logger {
			return &fxevent.ZapLogger{Logger: log}
		}),
		fx.Provide(
			db.NewDatabaseConnection,
			zap.NewProduction,
			handler.NewUserHandler,
		),
		fx.Invoke(
			NewServer,
		),
	).Run()
}
