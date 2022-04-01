package server

import (
	"log"
	"net/http"

	"hackathon/service/models"

	server_errors "hackathon/service/server/errors"

	"github.com/labstack/echo/v4"
)

func (s *Server) RenewToken(c echo.Context) error {
	newToken, err := s.api.GetNewTokenAPI(c)
	if err != nil {
		return server_errors.NewUnprocessablerequestError(err.Error())
	}
	result := models.Response{
		Message: *newToken,
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) Logout(c echo.Context) error {
	err := s.api.LogoutAPI(c)
	if err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	result := models.Response{
		Message: "Successfully logged out",
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) UpdateTask(c echo.Context) error {
	todo := models.NewTodo()
	if err := c.Bind(&todo); err != nil {
		log.Println(err)
		return server_errors.NewBadRequestError(err)
	}
	err := s.api.UpdateTaskAPI(c, todo)
	if err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	result := models.Response{
		Message: "task has been updated",
	}
	return c.JSON(http.StatusOK, result)
}
