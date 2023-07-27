package handlers

import (
	"github.com/gofiber/fiber/v2"
)

type BasicHandlers struct{}

func NewBasicHandlers() *BasicHandlers {
	return &BasicHandlers{}
}

// Base Root http handler
// @Summary Show the status of server.
// @Description Get the current server status.
// @Tags root
// @Success 200 {string} string "Server status"
// @Failure 500 {object} string "Server error"
// @Router / [get]
func (bh *BasicHandlers) Base(c *fiber.Ctx) error {
	return c.Status(fiber.StatusOK).SendString("It's working fine!")
}
