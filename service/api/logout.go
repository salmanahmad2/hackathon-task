package api

import (
	"hackathon/service/models"

	"github.com/labstack/echo/v4"
)

func (api *NonFungibleTokenAPIImpl) LogoutAPI(c echo.Context) error {
	user := models.NewUser()
	*user.Email = c.Get("email").(string)
	c.Request().Header["Authorization"][0] = ""
	_ = api.redis.Del(c.Request().Context(), *user.Email)
	return nil
}
