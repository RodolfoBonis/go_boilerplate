package di

import (
	"github.com/RodolfoBonis/go_boilerplate/core/config"
	"github.com/RodolfoBonis/go_boilerplate/core/services"
	"github.com/RodolfoBonis/go_boilerplate/features/login/domain/usecases"
)

func LoginInjection() usecases.LoginUseCase {
	return usecases.LoginUseCase{
		KeycloakClient:     services.AuthClient,
		KeycloakAccessData: config.EnvKeyCloak(),
	}
}
