package api

import (
	"fmt"
	"time"

	api_errors "hackathon/service/api/error"
	"hackathon/service/cache"

	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) SignInUserAPI(c echo.Context, user *models.User) (*string, error) {

	userData, err := api.db.GetUserDB(c, user)
	if err != nil {
		return nil, handleAPIError(err)
	}
	if userData != nil {
		errHash := CheckPasswordHash(*user.Password, *userData.Password)
		if errHash != nil {
			return nil, errHash
		}
		token, err := api.GenerateUserTokenAPI(c, userData)
		if err != nil {
			return nil, api_errors.NewInternalServerError(err.Error())
		}
		return token, nil
	}
	return nil, api_errors.NewError("user not found")
}

func (api *NonFungibleTokenAPIImpl) GenerateUserTokenAPI(c echo.Context, userData *models.User) (*string, error) {
	//duration is time duration for access token
	duration := 1 * time.Hour
	expTime := SetExpTime(duration)
	tokenString, err := api.jwt.GenerateToken(userData, expTime)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	//duration is time duration for refresh token
	expTime = SetExpTime(cache.RefreshToken_Exp)

	refreshToken, err := api.jwt.GenerateToken(userData, expTime)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}

	cacheKey := cache.RefreshTokenKey + *userData.UserId
	err = api.cache.SaveToCache(c, cacheKey, []byte(fmt.Sprint(refreshToken)), cache.RefreshToken_Exp)
	if err != nil {
		return nil, api_errors.NewInternalServerError(err.Error())
	}
	return &tokenString, nil
}
