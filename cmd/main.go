package main

import (
	"fmt"

	"github.com/JhonatanRSantos/gorest/cmd/api/handlers"
	"github.com/JhonatanRSantos/gorest/cmd/api/routes"
	"github.com/JhonatanRSantos/gorest/internal/config"
	"github.com/JhonatanRSantos/gorest/internal/platform/golog"
	"github.com/JhonatanRSantos/gorest/internal/platform/webserver"

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
	config := config.GetConfig()
	golog.SetEnv(config.AppEnv)

	ws := webserver.NewWebServer(webserver.DefaultConfig(webserver.WebServerDefaultConfig{
		AppName: config.AppName,
		Swagger: webserver.WebServerSwaggerConfig{
			Title: "Gorest API",
		},
	}))

	ws.AddRoutes(routes.NewBaseRouter(ws.GetApp(), handlers.NewBasicHandlers()))

	if err := ws.Listen(fmt.Sprintf(":%s", config.AppPort)); err != nil {
		fmt.Printf("failed to graceful shutdown. Cause: %s\n", err)
	}
}
