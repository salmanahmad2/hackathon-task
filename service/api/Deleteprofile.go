package api

import (
	api_errors "hackathon/service/api/error"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) DeleteProfileAPI(c echo.Context) error {
	userId := c.Get("user-id").(string)
	err := api.db.DeleteUserDB(c, userId)
	if err != nil {
		return api_errors.NewInternalServerError(err.Error())
	}
	return nil
}
