package routes

import "github.com/gofiber/fiber/v2"

type UserHandlers interface {
	GetUserV1(c *fiber.Ctx) error
	CreateUserV1(c *fiber.Ctx) error
	DeactivateUserV1(c *fiber.Ctx) error
}

type UserRouter struct {
	app      *fiber.App
	handlers UserHandlers
}

func NewUserRouter(app *fiber.App, handlers UserHandlers) *UserRouter {
	return &UserRouter{
		app:      app,
		handlers: handlers,
	}
}

func (br *UserRouter) Start() {
	baseGroupV1 := br.app.Group("/v1/users")

	baseGroupV1.Get("/", br.handlers.GetUserV1)
	baseGroupV1.Post("/", br.handlers.CreateUserV1)
	baseGroupV1.Delete("/", br.handlers.DeactivateUserV1)
}
