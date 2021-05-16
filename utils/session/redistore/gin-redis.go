package redistore

import (
	"errors"

	"github.com/gin-contrib/sessions"
	"github.com/gomodule/redigo/redis"
)

// Store gin session store wrapper.
type Store interface {
	sessions.Store
}

// NewStore creates a new gin session store
//
// size: maximum number of idle connections.
// network: tcp or udp
// address: host:port
// password: redis-password
// Keys are defined in pairs to allow key rotation, but the common case is to set a single
// authentication key and optionally an encryption key.
//
// The first key in a pair is used for authentication and the second for encryption. The
// encryption key can be set to nil or omitted in the last pair, but the authentication key
// is required in all pairs.
//
// It is recommended to use an authentication key with 32 or 64 bytes. The encryption key,
// if set, must be either 16, 24, or 32 bytes to select AES-128, AES-192, or AES-256 modes.
func NewStore(size int, network, address, password string, keyPairs ...[]byte) (Store, *redis.Pool, error) {
	s, p, err := NewRediStore(size, network, address, password, keyPairs...)
	if err != nil {
		return nil, nil, err
	}
	return &store{s}, p, nil
}

// NewStoreWithDB - like NewStore but accepts `DB` parameter to select
// redis DB instead of using the default one ("0")
//
// Ref: https://godoc.org/github.com/boj/redistore#NewRediStoreWithDB
func NewStoreWithDB(size int, network, address, password, DB string, keyPairs ...[]byte) (Store, error) {
	s, err := NewRediStoreWithDB(size, network, address, password, DB, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &store{s}, nil
}

// NewStoreWithPool instantiates a RediStore with a *redis.Pool passed in.
//
// Ref: https://godoc.org/github.com/boj/redistore#NewRediStoreWithPool
func NewStoreWithPool(pool *redis.Pool, keyPairs ...[]byte) (Store, error) {
	s, err := NewRediStoreWithPool(pool, keyPairs...)
	if err != nil {
		return nil, err
	}
	return &store{s}, nil
}

type store struct {
	*RediStore
}

// GetRedisStore get the actual woking store.
// Ref: https://godoc.org/github.com/boj/redistore#RediStore
func GetRedisStore(s Store) (rediStore *RediStore, err error) {
	realStore, ok := s.(*store)
	if !ok {
		err = errors.New("unable to get the redis store: Store isn't *store")
		return
	}

	rediStore = realStore.RediStore
	return
}

// SetKeyPrefix sets the key prefix in the redis database.
func SetKeyPrefix(s Store, prefix string) error {
	rediStore, err := GetRedisStore(s)
	if err != nil {
		return err
	}

	rediStore.SetKeyPrefix(prefix)
	return nil
}

func (c *store) Options(options sessions.Options) {
	c.RediStore.Options = options.ToGorillaOptions()
}
