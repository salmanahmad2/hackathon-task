package api

import (
	api_errors "hackathon/service/api/error"
	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) UpdateTaskAPI(c echo.Context, todo *models.Todo) error {
	*todo.UserId = c.Get("user-id").(string)
	err := api.db.UpdateTodoDB(c, todo)
	if err != nil {
		return api_errors.NewInternalServerError(err.Error())
	}
	return nil
}
