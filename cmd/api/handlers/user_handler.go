package handlers

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"

	"github.com/JhonatanRSantos/gorest/internal/platform/gocontext"
	"github.com/JhonatanRSantos/gorest/internal/platform/golog"
	"github.com/JhonatanRSantos/gorest/internal/services/user/model"

	"github.com/gofiber/fiber/v2"
)

type UserService interface {
	GetUserV1(ctx context.Context, req model.GetUserRequestV1) (model.GetUserResponseV1, error)
	CreateUserV1(ctx context.Context, params model.CreateUserRequestV1) (model.CreateUserResponseV1, error)
	DeactivateUserV1(ctx context.Context, req model.DeactivateUserRequestV1) error
}

type UserHandlers struct {
	userService UserService
}

func NewUserHandlers(userService UserService) *UserHandlers {
	return &UserHandlers{
		userService: userService,
	}
}

// Create a new user
// @Summary Create a new user
// @Tags users
// @Accepts json
// @Produce json
// @Param user_info body model.CreateUserRequestV1 true "User information"
// @Success 200 {object} model.CreateUserResponseV1 "User information"
// @Failure 500 {object} string "Error description"
// @Router /v1/users [post]
func (uh UserHandlers) CreateUserV1(c *fiber.Ctx) error {
	ctx := gocontext.FromContext(c.Context())
	createUserRequest := model.CreateUserRequestV1{}

	if err := json.Unmarshal(c.Body(), &createUserRequest); err != nil {
		golog.Log().Error(ctx, fmt.Sprintf("failed to parse request body. %s", err))
		return c.Status(fiber.StatusBadRequest).SendString("failed to parse request body")
	}

	if user, err := uh.userService.CreateUserV1(ctx, createUserRequest); err != nil {
		if errors.Is(err, model.ErrMissingRequiredParams) {
			return c.Status(fiber.StatusBadRequest).SendString(model.ErrMissingRequiredParams.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString("failed to create user")
	} else {
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

// Get the user information
// @Summary Get user information
// @Tags users
// @Produce json
// @Param email query string false "User email"
// @Param user_id query string false "User ID"
// @Success 200 {object} model.GetUserResponseV1 "User information"
// @Failure 500 {object} string "Error description"
// @Router /v1/users [get]
func (uh *UserHandlers) GetUserV1(c *fiber.Ctx) error {
	ctx := gocontext.FromContext(c.Context())
	email := c.Query("email")
	userID := c.Query("user_id")
	getUserRequest := model.GetUserRequestV1{
		UserID: &userID,
		Email:  &email,
	}

	if user, err := uh.userService.GetUserV1(ctx, getUserRequest); err != nil {
		if errors.Is(err, model.ErrMissingRequiredParams) {
			return c.Status(fiber.StatusBadRequest).SendString(model.ErrMissingRequiredParams.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString("failed to get user")
	} else {
		return c.Status(fiber.StatusOK).JSON(user)
	}
}

// Deactivate user
// @Summary Deactivate user
// @Tags users
// @Param email query string false "User email"
// @Param user_id query string false "User ID"
// @Success 204
// @Failure 500 {object} string "Error description"
// @Router /v1/users [delete]
func (uh *UserHandlers) DeactivateUserV1(c *fiber.Ctx) error {
	ctx := gocontext.FromContext(c.Context())
	email := c.Query("email")
	userID := c.Query("user_id")
	deactivateUserRequest := model.DeactivateUserRequestV1{
		UserID: &userID,
		Email:  &email,
	}

	if err := uh.userService.DeactivateUserV1(ctx, deactivateUserRequest); err != nil {
		if errors.Is(err, model.ErrMissingRequiredParams) {
			return c.Status(fiber.StatusBadRequest).SendString(model.ErrMissingRequiredParams.Error())
		}
		return c.Status(fiber.StatusInternalServerError).SendString("failed to deactivate user")
	} else {
		return c.SendStatus(fiber.StatusNoContent)
	}
}
