package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/JhonatanRSantos/gorest/internal/platform/golog"
	"github.com/JhonatanRSantos/gorest/internal/services/user/model"
)

type UserRepository interface {
	GetUserV1(ctx context.Context, req model.GetUserRequestV1) (*model.DBUser, error)
	CreateUserV1(ctx context.Context, params model.CreateUserRequestV1) (*model.DBUser, error)
	DeactivateUserV1(ctx context.Context, req model.DeactivateUserRequestV1) error
}

type Service struct {
	userRepository UserRepository
}

func NewService(userRepository UserRepository) *Service {
	return &Service{
		userRepository: userRepository,
	}
}

func (s *Service) CreateUserV1(ctx context.Context, req model.CreateUserRequestV1) (model.CreateUserResponseV1, error) {
	if err := req.Validate(); err != nil {
		golog.Log().Error(ctx, fmt.Sprintf("validation error. %s", err))
		return model.CreateUserResponseV1{}, fmt.Errorf("validation error. %w", err)
	}

	if dbUser, err := s.userRepository.CreateUserV1(ctx, req); err != nil {
		golog.Log().Error(ctx, fmt.Sprintf("failed to create user. %s", err))
		return model.CreateUserResponseV1{}, err
	} else {
		return model.CreateUserResponseV1{
			BasicUserInformation: model.BasicUserInformation{
				UserID:          dbUser.UserID,
				Name:            dbUser.Name,
				Country:         dbUser.Country,
				DefaultLanguage: dbUser.DefaultLanguage,
				Email:           dbUser.Email,
				Phones:          dbUser.Phones,
			},
		}, nil
	}
}

func (s *Service) GetUserV1(ctx context.Context, req model.GetUserRequestV1) (model.GetUserResponseV1, error) {
	if err := req.Validate(); err != nil {
		golog.Log().Error(ctx, fmt.Sprintf("validation error. %s", err))
		return model.GetUserResponseV1{}, fmt.Errorf("validation error. %w", err)
	}

	if dbUser, err := s.userRepository.GetUserV1(ctx, req); err != nil {
		golog.Log().Error(ctx, fmt.Sprintf("failed to get user. %s", err))
		return model.GetUserResponseV1{}, fmt.Errorf("failed to get user. %w", err)
	} else {
		return model.GetUserResponseV1{
			BasicUserInformation: model.BasicUserInformation{
				UserID:          dbUser.UserID,
				Name:            dbUser.Name,
				Country:         dbUser.Country,
				DefaultLanguage: dbUser.DefaultLanguage,
				Email:           dbUser.Email,
				Phones:          dbUser.Phones,
			},
		}, nil
	}
}

func (s *Service) DeactivateUserV1(ctx context.Context, req model.DeactivateUserRequestV1) error {
	if err := req.Validate(); err != nil {
		golog.Log().Error(ctx, fmt.Sprintf("validation error. %s", err))
		return fmt.Errorf("validation error. %w", err)
	}

	if err := s.userRepository.DeactivateUserV1(ctx, req); err != nil {
		if !errors.Is(err, model.ErrNoRowsAffected) {
			golog.Log().Error(ctx, fmt.Sprintf("failed to deactivate user. %s", err))
			return fmt.Errorf("failed to deactivate user. %w", err)
		}
	}
	return nil
}
