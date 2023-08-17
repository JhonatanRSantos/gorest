package model

import (
	"errors"

	"github.com/JhonatanRSantos/gorest/internal/platform/gtypes"
)

var (
	ErrMissingRequiredParams = errors.New("missing required params")
)

type BasicUserInformation struct {
	UserID          gtypes.UUID `json:"user_id,omitempty"`
	Name            string      `json:"name,omitempty"`
	Country         string      `json:"country,omitempty"`
	DefaultLanguage string      `json:"default_language,omitempty"`
	Email           string      `json:"email,omitempty"`
	Phones          []string    `json:"phones,omitempty"`
}

type GetUserRequestV1 struct {
	UserID *string `json:"user_id,omitempty"`
	Email  *string `json:"email,omitempty"`
}

func (gur *GetUserRequestV1) Validate() error {
	if (gur.UserID == nil && gur.Email == nil) || (*gur.UserID == "" && *gur.Email == "") {
		return ErrMissingRequiredParams
	}
	return nil
}

type GetUserResponseV1 struct {
	BasicUserInformation
}

type CreateUserRequestV1 struct {
	Name            string   `json:"name"`
	Country         string   `json:"country"`
	DefaultLanguage string   `json:"default_language,omitempty"`
	Email           string   `json:"email"`
	Phones          []string `json:"phones"`
	Password        string   `json:"password"`
}

func (cur *CreateUserRequestV1) Validate() error {
	if cur.Name == "" || cur.Country == "" || cur.DefaultLanguage == "" || cur.Email == "" || cur.Password == "" {
		return ErrMissingRequiredParams
	}

	if len(cur.Phones) == 0 {
		cur.Phones = nil
	}
	return nil
}

type CreateUserResponseV1 struct {
	BasicUserInformation
}

type DeactivateUserRequestV1 struct {
	UserID *string `json:"user_id,omitempty"`
	Email  *string `json:"email,omitempty"`
}

func (dur *DeactivateUserRequestV1) Validate() error {
	if (dur.UserID == nil && dur.Email == nil) || (*dur.UserID == "" && *dur.Email == "") {
		return ErrMissingRequiredParams
	}
	return nil
}
