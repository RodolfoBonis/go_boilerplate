package usecases

import (
	"github.com/RodolfoBonis/go_boilerplate/core/utils"
	"github.com/RodolfoBonis/go_boilerplate/features/login/domain/entities"
	"github.com/gin-gonic/gin"
	"net/http"
)

// ValidateLogin godoc
// @Summary Validate login
// @Schemes
// @Description performs login of user
// @Tags login
// @Accept json
// @Produce json
// @Param _ body entities.RequestLoginEntity true "Login Data"
// @Success 200 {object} entities.LoginResponseEntity
// @Failure 400 {object} utils.HttpError
// @Failure 409 {object} utils.HttpError
// @Failure 500 {object} utils.HttpError
// @Router /login [post]
func (uc *LoginUseCase) ValidateLogin(c *gin.Context) {
	loginData := new(entities.RequestLoginEntity)

	err := c.BindJSON(loginData)

	if err != nil {
		internalError := utils.BadRequestError(err.Error())
		c.AbortWithStatusJSON(internalError.StatusCode, internalError)
		return
	}

	jwt, err := uc.KeycloakClient.Login(
		c,
		uc.KeycloakAccessData.ClientID,
		uc.KeycloakAccessData.ClientSecret,
		uc.KeycloakAccessData.Realm,
		loginData.Email,
		loginData.Password,
	)

	if err != nil {
		internalError := utils.ForbiddenError(err.Error())
		c.AbortWithStatusJSON(internalError.StatusCode, internalError)
		return
	}

	c.JSON(http.StatusOK, entities.LoginResponseEntity{
		AccessToken:  jwt.AccessToken,
		RefreshToken: jwt.RefreshToken,
		ExpiresIn:    jwt.ExpiresIn,
	})
}
