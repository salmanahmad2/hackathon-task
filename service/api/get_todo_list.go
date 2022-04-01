package api

import (
	api_errors "hackathon/service/api/error"
	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) GetTodoListAPI(c echo.Context) ([]*models.Todo, error) {
	userId := c.Get("user-id").(string)
	todoList, err := api.db.GetTodoListDB(c, userId)
	if err != nil {
		return nil, api_errors.NewError(err.Error())
	}

	if todoList == nil {
		return nil, api_errors.NewUnProcessableRequest("list not found")
	}

	return todoList, nil
}
