package gtypes

import (
	"fmt"

	uuid "github.com/jackc/pgx-gofrs-uuid"
)

type UUID struct {
	uuid.UUID `json:"-"`
}

func (u UUID) String() string {
	if uuid, err := u.UUIDValue(); err != nil {
		return ""
	} else {
		if value, err := uuid.Value(); err != nil {
			return ""
		} else {
			return fmt.Sprint(value)
		}
	}
}

func (u UUID) MarshalJSON() ([]byte, error) {
	if value, err := u.UUIDValue(); err != nil {
		return nil, err
	} else {
		return value.MarshalJSON()
	}
}
