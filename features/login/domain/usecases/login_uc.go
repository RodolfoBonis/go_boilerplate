package usecases

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/RodolfoBonis/go_boilerplate/core/entities"
)

type LoginUseCase struct {
	KeycloakClient     *gocloak.GoCloak
	KeycloakAccessData entities.KeyCloakDataEntity
}
