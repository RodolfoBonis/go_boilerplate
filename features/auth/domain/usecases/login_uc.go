package usecases

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/entities"
)

type AuthUseCase struct {
	KeycloakClient     *gocloak.GoCloak
	KeycloakAccessData entities.KeyCloakDataEntity
}
