package db

// TODO: Add postgres integration
// type NFTDatabase interface {
// }

import (
	"fmt"
	"log"
	"os"

	"hackathon/config"
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
	file, err := os.OpenFile("schema.sql", os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		log.Fatal(err)
	}
	config.Cfg.Writer = file
	cfg := config.Cfg.Database
	psqlInfo := fmt.Sprintf("postgres://%v:%v@%v:%v/%v?sslmode=%v", cfg.DbUser, cfg.Password, cfg.Host, cfg.Port, cfg.DbName, cfg.SslEnable)
	log.Println(psqlInfo)
	conn, err := sqlx.Connect("postgres", psqlInfo)
	if err != nil {
		log.Println(err)
		fmt.Fprintf(os.Stderr, "Unable to connect to database: %v\n", err)
		//os.Exit(1)
	} else {
		log.Println("Connected")
	}

	return &NonFungibleTokenDBImpl{
		conn: conn,
	}

}

var _ NonFungibleTokenDB = &NonFungibleTokenDBImpl{}
