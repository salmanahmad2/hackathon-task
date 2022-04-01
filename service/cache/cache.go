package cache

import (
	"time"

	"github.com/labstack/echo/v4"
)

func (cache *NonFungibleTokenCacheImpl) SaveToCache(c echo.Context, key string, value []byte, duration time.Duration) error {
	cacheErr := cache.redis.Set(c.Request().Context(), key, value, duration).Err()
	if cacheErr != nil {
		return cacheErr
	}
	return nil
}

func (cache *NonFungibleTokenCacheImpl) CheckCache(c echo.Context, key string) ([]byte, error) {
	cacheData, cacheErr := cache.redis.Get(c.Request().Context(), key).Result()
	return []byte(cacheData), cacheErr
}

func (cache *NonFungibleTokenCacheImpl) DeleteCache(c echo.Context, key string) {
	_ = cache.redis.Del(c.Request().Context(), key)
}
