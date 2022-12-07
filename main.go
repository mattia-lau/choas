package main

import (
	"chaos-go/internal/config"
	"chaos-go/internal/database"
	"chaos-go/internal/logging"
	"chaos-go/internal/ticker"
	"context"
	"fmt"
	"net/http"
	"time"

	ginzap "github.com/gin-contrib/zap"
	"github.com/gin-gonic/gin"
	"go.uber.org/fx"
	"go.uber.org/zap"
)

func newServer(lc fx.Lifecycle, cfg *config.Config, logger *zap.Logger) *gin.Engine {
	r := gin.New()
	r.Use(ginzap.Ginzap(logger, time.RFC3339, true))
	r.Use(ginzap.RecoveryWithZap(logger, true))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", cfg.ServerConfig.Port),
		Handler: r,
	}
	lc.Append(fx.Hook{
		OnStart: func(ctx context.Context) error {
			go srv.ListenAndServe()
			return nil
		},
		OnStop: func(ctx context.Context) error {
			return srv.Shutdown(ctx)
		},
	})

	return r
}

func loadConfig() (*config.Config, error) {
	return config.Load("config.yml")
}

func main() {
	app := fx.New(
		fx.Provide(
			loadConfig,
			logging.SetupLogger,
			database.Init,
			ticker.NewHandler,
			ticker.NewTickerRepository,
			newServer,
		),
		fx.Invoke(
			ticker.Route,
			// ticker.CreateCronJob,
		),
	)
	app.Run()
}
