package usecases

import (
	"github.com/RodolfoBonis/go_boilerplate/core/utils"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

// Logout godoc
// @Summary Logout
// @Schemes
// @Description Logout the User
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} bool
// @Failure 400 {object} utils.HttpError
// @Failure 401 {object} utils.HttpError
// @Failure 403 {object} utils.HttpError
// @Failure 409 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /logout [post]
func (uc *AuthUseCase) Logout(c *gin.Context) {
	authHeader := c.GetHeader("Authorization")

	if len(authHeader) < 1 {
		err := utils.BadRequestError("Refresh token invalid")
		c.AbortWithStatusJSON(err.StatusCode, err)
		c.Abort()
		return
	}

	refreshToken := strings.Split(authHeader, " ")[1]

	err := uc.KeycloakClient.Logout(
		c,
		uc.KeycloakAccessData.ClientID,
		uc.KeycloakAccessData.ClientSecret,
		uc.KeycloakAccessData.Realm,
		refreshToken,
	)

	if err != nil {
		currentError := utils.BadRequestError(err.Error())
		c.AbortWithStatusJSON(currentError.StatusCode, currentError)
		c.Abort()
		return
	}

	c.JSON(http.StatusOK, true)
}
