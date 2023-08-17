package pgclient

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func ParseRowsTo[T any](ctx context.Context, rows pgx.Rows) ([]*T, error) {
	return pgx.CollectRows(rows, pgx.RowToAddrOfStructByName[T])
}

func ParseRowTo[T any](ctx context.Context, rows pgx.Rows) (*T, error) {
	return pgx.CollectOneRow(rows, pgx.RowToAddrOfStructByName[T])
}
