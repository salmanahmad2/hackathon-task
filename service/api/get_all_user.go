package api

import (
	"encoding/json"
	"log"

	api_errors "hackathon/service/api/error"
	"hackathon/service/cache"
	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) GetAllUsersAPI(c echo.Context) ([]*models.User, error) {
	var err error
	var usersData []*models.User
	if cachedData, cachedErr := api.cache.CheckCache(c, cache.AllUsersKey); cachedErr == nil && len(cachedData) > 0 {
		err = json.Unmarshal(cachedData, &usersData)
		if err != nil {
			return nil, api_errors.NewInternalServerError(err.Error())
		}
	} else {
		log.Println("Fetching all users from database")
		usersData, err = api.db.GetAllUsersDB(c)
		if err != nil {
			return nil, api_errors.NewInternalServerError(err.Error())
		}
		cacheValue, err := json.Marshal(&usersData)
		if err != nil {
			return nil, api_errors.NewInternalServerError(err.Error())
		}
		err = api.cache.SaveToCache(c, cache.AllUsersKey, cacheValue, cache.AllUsersKey_Exp)
		if err != nil {
			return nil, api_errors.NewInternalServerError(err.Error())
		}
	}
	return usersData, nil
}
