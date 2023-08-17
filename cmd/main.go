package main

import (
	"context"
	"fmt"
	"os"

	"github.com/JhonatanRSantos/gorest/cmd/api/handlers"
	"github.com/JhonatanRSantos/gorest/cmd/api/routes"
	"github.com/JhonatanRSantos/gorest/internal/config"
	"github.com/JhonatanRSantos/gorest/internal/platform/database/pgclient"
	"github.com/JhonatanRSantos/gorest/internal/platform/gocontext"
	"github.com/JhonatanRSantos/gorest/internal/platform/golog"
	"github.com/JhonatanRSantos/gorest/internal/platform/webserver"

	userSvc "github.com/JhonatanRSantos/gorest/internal/services/user"
	userRepo "github.com/JhonatanRSantos/gorest/internal/services/user/repository"

	_ "github.com/JhonatanRSantos/gorest/docs"
)

// More details: https://github.com/gofiber/swagger

// @title Gorest API
// @version 1.0
// @description This is an auto-generated swagger for the Gorest project
// @host localhost:8080
// @BasePath /
// @schemas http
func main() {
	var (
		err      error
		pgClient *pgclient.PGClient
	)

	ctx := gocontext.FromContext(context.Background())
	config := config.GetConfig()
	golog.SetEnv(config.AppEnv)

	if pgClient, err = pgclient.NewClient(
		ctx,
		pgclient.PGClientConfig{
			Host:   config.DBConfig.Host,
			Port:   uint16(config.DBConfig.Port),
			User:   config.DBConfig.User,
			Pass:   config.DBConfig.Pass,
			DBName: config.DBConfig.DBName,
		},
	); err != nil {
		fatalf(ctx, "failed to prepare database. Cause: %s", err)
	}

	defer pgClient.Close()

	userRepository := userRepo.NewRepository(pgClient)
	userService := userSvc.NewService(userRepository)

	ws := webserver.NewWebServer(webserver.DefaultConfig(webserver.WebServerDefaultConfig{
		AppName: config.AppName,
		Swagger: webserver.WebServerSwaggerConfig{
			Title: "Gorest API",
		},
	}))

	ws.AddRoutes(routes.NewBaseRouter(ws.GetApp(), handlers.NewBasicHandlers()))
	ws.AddRoutes(routes.NewUserRouter(ws.GetApp(), handlers.NewUserHandlers(userService)))

	if err := ws.Listen(fmt.Sprintf(":%s", config.AppPort)); err != nil {
		golog.Log().Error(ctx, fmt.Sprintf("failed to graceful shutdown. Cause: %s", err))
	}
}

func fatalf(ctx context.Context, message string, args ...any) {
	golog.Log().Error(ctx, fmt.Sprintf(message, args...))
	os.Exit(1)
}
