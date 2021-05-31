package authorization

import (
	"context"
	"database/sql"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// Service authorization module service.
//
// Implements ctxinfo.Service interface.
type Service struct {
	baseservice.JWTData

	db    *sql.DB
	redis *redis.Pool
	mux   *sync.Mutex

	roles      []string
	privileges privileges
}

// NewAuthService creates a new Service.
func NewAuthService(c *httphandler.Context, jwtSvc JWTSvc, db *sql.DB, redis *redis.Pool) (*Service, error) {
	jwtData, err := jwtSvc.AuthJWT(c.Request)
	if err != nil {
		return nil, err
	}
	return &Service{
		JWTData:    *jwtData,
		db:         db,
		redis:      redis,
		mux:        &sync.Mutex{},
		roles:      nil,
		privileges: nil,
	}, nil
}

// CheckAuthorization check if the user
func (svc *Service) CheckAuthorization(c *httphandler.Context) error {
	if ok, err := svc.IsAdmin(c); err != nil {
		return err
	} else if ok {
		return nil
	}
	privs, err := svc.getPrivileges(c)
	if err != nil {
		return err
	}
	for i := range privs {
		if privs[i].Method == c.Request.Method && privs[i].Route == c.FullPath() {
			return nil
		}
	}
	return baseerrors.ErrPrivilege
}

// GetPrivilegeByPriority baseservice.AuthService implementation.
// More info given in its documentation.
func (svc *Service) GetPrivilegeByPriority(ctx context.Context, privileges ...string) (string, error) {
	if len(privileges) == 0 {
		return "", nil
	}
	access, err := svc.getPrivileges(ctx)
	if err != nil {
		return "", err
	}
	return access.getPrivilegeByPriority(privileges...), nil
}

// HasPrivilege baseservice.AuthService implementation.
// More info given in its documentation.
func (svc *Service) HasPrivilege(ctx context.Context, priv string) (bool, error) {
	if priv == "" {
		return false, nil
	}
	privs, err := svc.getPrivileges(ctx)
	if err != nil {
		return false, err
	}
	for i := range privs {
		if privs[i].Key == priv {
			return true, nil
		}
	}
	return false, nil
}

// IsAdmin baseservice.AuthService implementation.
// More info given in its documentation.
func (svc *Service) IsAdmin(ctx context.Context) (bool, error) {
	roles, err := svc.getRoles(ctx)
	if err != nil {
		return false, err
	}
	for i := range roles {
		if roles[i] == baseservice.RoleAdmin {
			return true, nil
		}
	}
	return false, nil
}
