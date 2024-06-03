package services

import (
	"github.com/Nerzal/gocloak/v13"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/config"
)

var AuthClient *gocloak.GoCloak

func InitializeOAuthServer() {
	keycloakDataAccess := config.EnvKeyCloak()

	AuthClient = gocloak.NewClient(keycloakDataAccess.Host)
}
