package server

import (
	"hackathon/service/api"
	"hackathon/service/lib"
	"hackathon/service/middlewares"

	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

type ServerImpl interface {
	Signup(c echo.Context) error
	SignIn(c echo.Context) error
	GetAllUsers(c echo.Context) error
	GetProfile(c echo.Context) error
	DeleteProfile(c echo.Context) error
	RenewToken(c echo.Context) error
	Logout(c echo.Context) error
}
type Server struct {
	api *api.NonFungibleTokenAPIImpl
}

func NewServer() *Server {
	return &Server{
		api: api.NewNonFungibleTokenAPIImpl(),
	}
}

func NewServerImpl(e *echo.Echo) {

	server := NewServer()
	// v := validator.New()
	e.Validator = lib.NewCustomValidator()

	e.POST("/signup", server.Signup)
	e.POST("/signin", server.SignIn)
	e.GET("/profile/:user_id", server.GetProfile)

	e.DELETE("/profile", server.DeleteProfile, middlewares.UserAuth)
	e.GET("/getalluser", server.GetAllUsers, middlewares.UserAuth)

	e.GET("/renewtoken", server.RenewToken)
	e.GET("/logout", server.Logout, middlewares.UserAuth)
	e.GET("/todolist", server.GetTodoList, middlewares.UserAuth)
	e.GET("/todofilteredlist/:status:", server.GetTodoFilteredList, middlewares.UserAuth)
	e.DELETE("/todo/:todo_id", server.DeleteTask, middlewares.UserAuth)
	e.PUT("/todo/:todo_id", server.UpdateTask, middlewares.UserAuth)

	e.GET("/swagger/*", echoSwagger.WrapHandler)
}

var _ ServerImpl = &Server{}
