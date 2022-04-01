package server

import (
	"net/http"

	"hackathon/service/models"

	server_errors "hackathon/service/server/errors"

	"github.com/labstack/echo/v4"
)

func (s *Server) Signup(c echo.Context) error {
	body := models.NewSignUpRequest()
	if err := c.Bind(&body); err != nil {
		return server_errors.NewBadRequestError(err)
	}
	if err := c.Validate(body); err != nil {
		return server_errors.NewUnprocessablerequestError("Please enter valid credentials")
	}
	newUser := models.NewUser()
	newUser.Email = body.Email
	newUser.Password = body.Password

	err := s.api.SignUpUserAPI(c, newUser, body)
	if err != nil {
		return server_errors.NewUnprocessablerequestError(err.Error())
	}
	var success string
	if *body.OTP == "" {
		success = "Code has been sent to your email"
	} else {
		success = "Signup is successfull"
	}
	result := models.Response{
		Message: success,
	}
	return c.JSON(http.StatusOK, result)
}

func (s *Server) SignIn(c echo.Context) error {
	body := models.NewSignInRequest()
	if err := c.Bind(&body); err != nil {
		return server_errors.NewInternalServerError(err.Error())
	}
	if err := c.Validate(body); err != nil {
		return server_errors.NewUnprocessablerequestError("Please enter valid credentials")
	}
	user := models.NewUser()
	user.Email = body.Email
	user.Password = body.Password
	tokenString, err := s.api.SignInUserAPI(c, user)
	if err != nil {
		return server_errors.NewUnauthorizedError(err.Error())
	}
	success := tokenString
	result := models.Response{
		Message: *success,
	}
	return c.JSON(http.StatusOK, result)
}
