package repository

import (
	"context"

	"github.com/JhonatanRSantos/gorest/internal/platform/database/pgclient"
	"github.com/JhonatanRSantos/gorest/internal/services/user/model"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PostgresDB interface {
	Close()
	GetConnection(ctx context.Context) (*pgxpool.Conn, error)
}

type Repository struct {
	pgDB PostgresDB
}

func NewRepository(pgDB PostgresDB) *Repository {
	return &Repository{
		pgDB: pgDB,
	}
}

func (r *Repository) GetUserV1(ctx context.Context, req model.GetUserRequestV1) (*model.DBUser, error) {
	if conn, err := r.pgDB.GetConnection(ctx); err != nil {
		return nil, err
	} else {
		defer conn.Release()

		args := pgx.NamedArgs{"email": req.Email}
		query := getUserByEmail

		if req.UserID != nil && *req.UserID != "" {
			args = pgx.NamedArgs{"user_id": *req.UserID}
			query = getUserByID
		}

		if rows, err := conn.Query(ctx, query, args); err != nil {
			return nil, err
		} else {
			return pgclient.ParseRowTo[model.DBUser](ctx, rows)
		}
	}
}

func (r *Repository) CreateUserV1(ctx context.Context, params model.CreateUserRequestV1) (*model.DBUser, error) {
	if conn, err := r.pgDB.GetConnection(ctx); err != nil {
		return nil, err
	} else {
		defer conn.Release()
		args := pgx.NamedArgs{
			"name":             params.Name,
			"country":          params.Country,
			"default_language": params.DefaultLanguage,
			"email":            params.Email,
			"phones":           params.Phones,
			"password":         params.Password,
		}
		if rows, err := conn.Query(ctx, createUser, args); err != nil {
			return nil, err
		} else {
			return pgclient.ParseRowTo[model.DBUser](ctx, rows)
		}
	}
}

func (r *Repository) DeactivateUserV1(ctx context.Context, req model.DeactivateUserRequestV1) error {
	if conn, err := r.pgDB.GetConnection(ctx); err != nil {
		return err
	} else {
		defer conn.Release()

		args := pgx.NamedArgs{"email": req.Email}
		query := deactivateDeleteUserByEmail

		if req.UserID != nil && *req.UserID != "" {
			args = pgx.NamedArgs{"user_id": *req.UserID}
			query = deactivateDeleteUserByID
		}

		if result, err := conn.Exec(ctx, query, args); err != nil {
			return err
		} else {
			if result.RowsAffected() != 1 {
				return model.ErrNoRowsAffected
			}
			return nil
		}
	}
}
