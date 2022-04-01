package server

import (
	"net/http"

	"hackathon/service/models"
	server_errors "hackathon/service/server/errors"

	"github.com/labstack/echo/v4"
)

func (s *Server) GetAllUsers(c echo.Context) error {
	users, err := s.api.GetAllUsersAPI(c)
	if err != nil {
		return server_errors.NewUnprocessablerequestError(err.Error())
	}
	var messageResp string
	if len(users) == 0 {
		messageResp = "No user found"
	} else {
		messageResp = "success"
	}
	result := models.Response{
		Message: messageResp,
		Data:    users,
	}
	return c.JSON(http.StatusOK, result)
}
func (s *Server) GetProfile(c echo.Context) error {
	userId := c.Param("user_id")
	isValid := IsValidUUID(userId)
	if !isValid {
		return server_errors.NewUnprocessablerequestError("Please enter valid user id")
	}
	getProfile, err := s.api.GetProfileAPI(c, userId)
	if err != nil {
		return server_errors.NewUnprocessablerequestError(err.Error())
	}
	result := models.Response{
		Message: "Success",
		Data:    getProfile,
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) GetTodoList(c echo.Context) error {
	getList, err := s.api.GetTodoListAPI(c)
	if err != nil {
		return server_errors.NewUnprocessablerequestError(err.Error())
	}
	result := models.Response{
		Message: "Success",
		Data:    getList,
	}
	return c.JSON(http.StatusOK, result)
}
func (s *Server) GetTodoFilteredList(c echo.Context) error {
	taskId := c.Param("status")
	getList, err := s.api.GetTodoTaskFilteredAPI(c, taskId)
	if err != nil {
		return server_errors.NewUnprocessablerequestError(err.Error())
	}
	result := models.Response{
		Message: "Success",
		Data:    getList,
	}
	return c.JSON(http.StatusOK, result)
}
