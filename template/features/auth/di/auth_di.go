package di

import (
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/config"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/core/services"
	"github.com/{{cookiecutter.github_username}}/{{cookiecutter.package_name}}/features/auth/domain/usecases"
)

func AuthInjection() usecases.AuthUseCase {
	return usecases.AuthUseCase{
		KeycloakClient:     services.AuthClient,
		KeycloakAccessData: config.EnvKeyCloak(),
	}
}
