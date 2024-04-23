package di

import (
	"github.com/RodolfoBonis/go_boilerplate/core/config"
	"github.com/RodolfoBonis/go_boilerplate/core/services"
	"github.com/RodolfoBonis/go_boilerplate/features/auth/domain/usecases"
)

func AuthInjection() usecases.AuthUseCase {
	return usecases.AuthUseCase{
		KeycloakClient:     services.AuthClient,
		KeycloakAccessData: config.EnvKeyCloak(),
	}
}
