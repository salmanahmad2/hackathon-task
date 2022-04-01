package db

// TODO: Add postgres integration
// type NFTDatabase interface {
// }

import (
	"fmt"

	"hackathon/service/models"

	"github.com/jmoiron/sqlx"
	"github.com/labstack/echo/v4"
	_ "github.com/lib/pq"
)

type NonFungibleTokenDB interface {
	SignupAPIDB(c echo.Context, user *models.User) error
	GetUserDB(c echo.Context, user *models.User) (*models.User, error)
	GetAllUsersDB(c echo.Context) ([]*models.User, error)

	UpdatePasswordDB(c echo.Context, user *models.User) error
	DeleteUserDB(c echo.Context, userId string) error
}

type NonFungibleTokenDBImpl struct {
	conn *sqlx.DB
}

func NewNonFungibleTokenDBImpl() *NonFungibleTokenDBImpl {
	conn, err := sqlx.Connect("mysql", "database:Salman_database@123@tcp(127.0.0.1:3306)/salmanDB")
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("db is connected")
	}
	return &NonFungibleTokenDBImpl{
		conn: conn,
	}
}

var _ NonFungibleTokenDB = &NonFungibleTokenDBImpl{}
