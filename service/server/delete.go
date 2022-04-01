package server

import (
	"net/http"

	"hackathon/service/models"
	server_errors "hackathon/service/server/errors"

	"github.com/labstack/echo/v4"
)

func (s *Server) DeleteProfile(c echo.Context) error {
	err := s.api.DeleteProfileAPI(c)
	if err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	result := models.Response{
		Message: "Profile has been Deleted",
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) DeleteTask(c echo.Context) error {
	taskId := c.Param("todo_id")
	err := s.api.DeleteTaskAPI(c, taskId)
	if err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	result := models.Response{
		Message: "task has been Deleted",
	}
	return c.JSON(http.StatusOK, result)
}
