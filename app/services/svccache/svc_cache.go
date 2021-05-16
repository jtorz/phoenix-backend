package svccache

import (
	"context"
	"encoding/json"
	"log"
	"time"

	"github.com/allegro/bigcache/v2"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
)

// CacheSvc Cache service.
type CacheSvc struct {
	*bigcache.BigCache
}

func NewCacheSvc() (*CacheSvc, error) {
	bc, err := bigcache.NewBigCache(bigcache.DefaultConfig(40 * time.Minute))
	if err != nil {
		return nil, err
	}
	return &CacheSvc{BigCache: bc}, nil
}

func (cacheSvc CacheSvc) Get(ctx context.Context, v interface{}, key string) (found bool) {
	b, found, err := cacheSvc.GetBytes(key)
	if err != nil {
		if ctxinfo.PrintLog(ctx) {
			log.Printf("cache error: can't retrieve value: %s", err)
		}
		return false
	}

	if !found {
		return false
	}

	err = json.Unmarshal(b, v)
	if err != nil {
		if ctxinfo.PrintLog(ctx) {
			log.Printf("cache error: can't unmarshal %T: %s", v, err)
		}
		return false
	}
	return true
}

// Save stores the values in the cache.
func (cacheSvc CacheSvc) Save(ctx context.Context, v interface{}, key string) {
	b, err := json.Marshal(v)
	if err != nil {
		if ctxinfo.PrintLog(ctx) {
			log.Printf("cache error: can't unmarshal %T: %s", v, err)
		}
	}
	err = cacheSvc.SetBytes(key, b)
	if err != nil {
		if ctxinfo.PrintLog(ctx) {
			log.Printf("cache error: can't set value: %s", err)
		}
	}
}

// GetBytes returns the stored bytes in the cache.
func (cacheSvc CacheSvc) GetBytes(k string) (bytez []byte, found bool, err error) {
	bytez, err = cacheSvc.BigCache.Get(k)
	if err != nil {
		if err == bigcache.ErrEntryNotFound {
			return nil, false, nil
		}
		return nil, false, err
	}
	return bytez, true, nil
}

// SetBytes stores the data in the cache.
func (cacheSvc CacheSvc) SetBytes(key string, bytez []byte) error {
	return cacheSvc.BigCache.Set(key, bytez)
}
