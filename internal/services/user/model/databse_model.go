package model

import (
	"time"

	"github.com/JhonatanRSantos/gorest/internal/platform/gtypes"
)

type DBUser struct {
	UserID          gtypes.UUID `db:"user_id"`
	Name            string      `db:"name"`
	Country         string      `db:"country"`
	DefaultLanguage string      `db:"default_language"`
	Email           string      `db:"email"`
	ConfirmedEmail  bool        `db:"confirmed_email"`
	Phones          []string    `db:"phones"`
	Password        string      `db:"password"`
	CreatedAt       *time.Time  `db:"created_at"`
	UpdatedAt       *time.Time  `db:"updated_at"`
	DeletedAt       *time.Time  `db:"deleted_at"`
}
