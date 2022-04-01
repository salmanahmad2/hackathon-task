package api

import (
	api_errors "hackathon/service/api/error"
	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) GetTodoTaskFilteredAPI(c echo.Context, status string) ([]*models.Todo, error) {
	userId := c.Get("user-id").(string)
	todoList, err := api.db.GetTodoTaskFilteredDB(c, userId, status)
	if err != nil {
		return nil, api_errors.NewError(err.Error())
	}

	if todoList == nil {
		return nil, api_errors.NewUnProcessableRequest("todoList not found")
	}

	return todoList, nil
}
