package api

import (
	api_errors "hackathon/service/api/error"
	"hackathon/service/cache"
	"hackathon/service/middlewares"
	"hackathon/service/models"

	"github.com/golang-jwt/jwt"
	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) GetNewTokenAPI(c echo.Context) (*string, error) {
	userData := models.NewUser()
	pubkey := api.jwt.GetPublicKey()
	accessToken := c.Request().Header["Authorization"][0]
	accessClaims, err := middlewares.ValidateToken(accessToken, pubkey)
	v, _ := err.(*jwt.ValidationError)
	if v.Errors != jwt.ValidationErrorExpired {
		return nil, api_errors.NewUnauthorizedError("You are not authorized")
	} else if err != nil && v.Errors != jwt.ValidationErrorExpired {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	*userData.UserId = accessClaims["user_id"].(string)
	cacheKey := cache.RefreshTokenKey + *userData.UserId
	cacheData, cacheErr := api.cache.CheckCache(c, cacheKey)
	if cacheErr != nil {
		return nil, api_errors.NewUnauthorizedError("Token has been expired.Please login again.")
	}
	refreshToken := string(cacheData)

	refreshClaims, err := middlewares.ValidateToken(refreshToken, pubkey)
	if err != nil {
		return nil, api_errors.NewUnauthorizedError(err.Error())
	}
	currentTime := api.clock.NowUnix()
	timeLeft := refreshClaims["exp"].(float64) - float64(currentTime)
	var duration float64
	if timeLeft < accessClaims["duration"].(float64) {
		duration = timeLeft
	} else {
		duration = accessClaims["duration"].(float64)
	}
	expTime := api.clock.NowUnix() + int64(duration)
	*userData.UserId = refreshClaims["user_id"].(string)

	newToken, err := api.jwt.GenerateToken(userData, expTime)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return &newToken, nil
}
