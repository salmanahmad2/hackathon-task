package api

import (
	api_errors "hackathon/service/api/error"
	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) GetProfileAPI(c echo.Context, userId string) (*models.User, error) {
	var err error
	var userProfile *models.User

	userProfile, err = api.db.GetProfileDB(c, userId)
	if err != nil {
		return nil, api_errors.NewError(err.Error())
	}

	if userProfile == nil {
		return nil, api_errors.NewUnProcessableRequest("user not found")
	}

	return userProfile, nil
}
