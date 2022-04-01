package cache

import (
	"time"

	"hackathon/service/lib/nftredis"

	"github.com/go-redis/redis/v8"
	"github.com/labstack/echo/v4"
)

type NonFungibleTokenCache interface {
	SaveToCache(c echo.Context, key string, value []byte, duration time.Duration) error
	CheckCache(c echo.Context, key string) ([]byte, error)
}

type NonFungibleTokenCacheImpl struct {
	redis *redis.Client
}

func NewNonFungibleTokenCacheImpl() *NonFungibleTokenCacheImpl {
	return &NonFungibleTokenCacheImpl{
		redis: nftredis.NewClient(),
	}
}

var _ NonFungibleTokenCache = &NonFungibleTokenCacheImpl{}
