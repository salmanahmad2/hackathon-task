package api

import (
	"context"
	"hackathon/pkg/utils"
	"hackathon/service/cache"
	"hackathon/service/db"
	"hackathon/service/lib/nftredis"
	"hackathon/service/middlewares"
	"hackathon/service/models"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

const (
	PblicKeyPath = "config/mypubkey.pem"
	PvtKeyPath   = "config/mykey.pem"
)

//go:generate mockery --name NonFungibleTokenAPI --case snake --output ../../mocks/service/api
type NonFungibleTokenAPI interface {
	SignUpUserAPI(c echo.Context, user *models.User, signup *models.SignUpRequest) error
	ConfirmOTPAPI(c echo.Context, user *models.User, signup *models.SignUpRequest) error
	SignInUserAPI(c echo.Context, user *models.User) (*string, error)

	GetAllUsersAPI(c echo.Context) ([]*models.User, error)

	GetNewTokenAPI(c echo.Context) (*string, error)
	SendEmailForSignUpAPI(c echo.Context, user *models.User) error
	SendEmailForForgotPasswordAPI(c echo.Context, user *models.User) error
	GenerateUserTokenAPI(c echo.Context, userData *models.User) (*string, error)
	LogoutAPI(c echo.Context) error

	DeleteProfileAPI(c echo.Context) error
}

type NonFungibleTokenAPIImpl struct {
	db    *db.NonFungibleTokenDBImpl
	redis *redis.Client
	jwt   *middlewares.JWT
	clock *utils.ClockImpl
	cache *cache.NonFungibleTokenCacheImpl
}

var Cancel context.CancelFunc

func NewNonFungibleTokenAPIImpl() *NonFungibleTokenAPIImpl {

	dbImpl := db.NewNonFungibleTokenDBImpl()
	clockImp := utils.NewClock()
	cacheImpl := cache.NewNonFungibleTokenCacheImpl()
	return &NonFungibleTokenAPIImpl{
		db:    dbImpl,
		redis: nftredis.NewClient(),
		jwt:   middlewares.NewJWT(PvtKeyPath, PblicKeyPath),
		clock: clockImp,
		cache: cacheImpl,
	}
}

var _ NonFungibleTokenAPI = &NonFungibleTokenAPIImpl{}
