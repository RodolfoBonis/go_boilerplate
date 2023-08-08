package entities

import (
	"github.com/RodolfoBonis/go_boilerplate/core/types"
	uuid "github.com/satori/go.uuid"
)

type JWTClaim struct {
	ID             uuid.UUID `json:"sub"`
	Verified       bool      `json:"email_verified"`
	Name           string    `json:"name"`
	Username       string    `json:"preferred_username"`
	FirstName      string    `json:"given_name"`
	FamilyName     string    `json:"family_name"`
	Email          string    `json:"email"`
	ResourceAccess client    `json:"resource_access,omitempty"`
}

type client struct {
	Api clientRoles `json:"api,omitempty"`
}
type clientRoles struct {
	Roles types.Array `json:"roles,omitempty"`
}
