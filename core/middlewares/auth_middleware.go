package middlewares

import (
	"encoding/json"
	"github.com/RodolfoBonis/go_boilerplate/core/config"
	"github.com/RodolfoBonis/go_boilerplate/core/entities"
	"github.com/RodolfoBonis/go_boilerplate/core/services"
	"github.com/RodolfoBonis/go_boilerplate/core/utils"
	"github.com/gin-gonic/gin"
	"strings"

	jsonToken "github.com/golang-jwt/jwt/v4"
)

func Protect(handler gin.HandlerFunc, role string) gin.HandlerFunc {
	return func(c *gin.Context) {
		keycloakDataAccess := config.EnvKeyCloak()
		authHeader := c.GetHeader("Authorization")

		if len(authHeader) < 1 {
			err := utils.UnauthorizedError()
			c.AbortWithStatusJSON(err.StatusCode, err)
			c.Abort()
			return
		}

		accessToken := strings.Split(authHeader, " ")[1]

		rptResult, err := services.AuthClient.RetrospectToken(
			c,
			accessToken,
			keycloakDataAccess.ClientID,
			keycloakDataAccess.ClientSecret,
			keycloakDataAccess.Realm,
		)

		if err != nil {
			currentError := utils.BadRequestError(err.Error())
			c.AbortWithStatusJSON(currentError.StatusCode, currentError)
			c.Abort()
			return
		}

		isTokenValid := *rptResult.Active

		if !isTokenValid {
			currentError := utils.UnauthorizedError()
			c.AbortWithStatusJSON(currentError.StatusCode, currentError)
			c.Abort()
			return
		}

		token, _, err := services.AuthClient.DecodeAccessToken(
			c,
			accessToken,
			keycloakDataAccess.Realm,
		)

		if err != nil {
			currentError := utils.BadRequestError(err.Error())
			c.JSON(currentError.StatusCode, currentError)
			return
		}

		claims := token.Claims.(jsonToken.MapClaims)

		jsonData, _ := json.Marshal(claims)

		var userClaim entities.JWTClaim

		err = json.Unmarshal(jsonData, &userClaim)
		if err != nil {
			currentError := utils.BadRequestError(err.Error())
			c.JSON(currentError.StatusCode, currentError)
			return
		}

		containsRole := userClaim.ResourceAccess.Api.Roles.Contains(role)

		if !containsRole {
			currentError := utils.UnauthorizedError()
			c.AbortWithStatusJSON(currentError.StatusCode, currentError)
			c.Abort()
			return
		}

		c.Set("claims", userClaim)
		handler(c)
	}
}
