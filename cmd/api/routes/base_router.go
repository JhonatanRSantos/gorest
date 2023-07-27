package routes

import "github.com/gofiber/fiber/v2"

type BasicHandlers interface {
	Base(ctx *fiber.Ctx) error
}

type BaseRouter struct {
	app      *fiber.App
	handlers BasicHandlers
}

func NewBaseRouter(app *fiber.App, handlers BasicHandlers) *BaseRouter {
	return &BaseRouter{
		app:      app,
		handlers: handlers,
	}
}

func (br *BaseRouter) Start() {
	baseGroup := br.app.Group("/")
	baseGroup.Get("/", br.handlers.Base)
}
