package api

import (
	api_errors "hackathon/service/api/error"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) DeleteTaskAPI(c echo.Context, taskId string) error {
	userId := c.Get("user-id").(string)
	err := api.db.DeleteTodoDB(c, userId, taskId)
	if err != nil {
		return api_errors.NewInternalServerError(err.Error())
	}
	return nil
}
