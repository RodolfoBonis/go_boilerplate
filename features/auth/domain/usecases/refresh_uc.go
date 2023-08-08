package usecases

import (
	"github.com/RodolfoBonis/go_boilerplate/core/utils"
	"github.com/RodolfoBonis/go_boilerplate/features/auth/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// RefreshAuthToken godoc
// @Summary Refresh Login Access Token
// @Schemes
// @Description Refresh User Token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} entities.LoginResponseEntity
// @Failure 400 {object} utils.HttpError
// @Failure 401 {object} utils.HttpError
// @Failure 403 {object} utils.HttpError
// @Failure 409 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /refresh [post]
func (uc *AuthUseCase) RefreshAuthToken(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if len(authHeader) < 1 {
		err := utils.BadRequestError("Refresh token invalid")
		c.AbortWithStatusJSON(err.StatusCode, err)
		c.Abort()
		return
	}

	refreshToken := strings.Split(authHeader, " ")[1]

	token, err := uc.KeycloakClient.RefreshToken(
		c,
		refreshToken,
		uc.KeycloakAccessData.ClientID,
		uc.KeycloakAccessData.ClientSecret,
		uc.KeycloakAccessData.Realm,
	)

	if err != nil {
		currentError := utils.BadRequestError(err.Error())
		c.AbortWithStatusJSON(currentError.StatusCode, currentError)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, entities.LoginResponseEntity{
		AccessToken:  token.AccessToken,
		RefreshToken: token.RefreshToken,
		ExpiresIn:    token.ExpiresIn,
	})
}
