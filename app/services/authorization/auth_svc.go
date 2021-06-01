package authorization

import (
	"context"
	"net/http"
	"sync"

	"github.com/gomodule/redigo/redis"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// Service authorization module service.
//
// Implements ctxinfo.Service interface.
//
// Roles and Privileges are loaded once and only the firts time they are nedded.
type Service struct {
	baseservice.JWTData

	exe   base.Executor
	redis *redis.Pool

	// mutex to protect the writing of Roles and privileges form concurrent calls.
	mux *sync.Mutex

	roles      []string
	privileges privileges
}

// NewAuthService creates a new Service.
func NewAuthService(req *http.Request, jwtSvc JWTSvc, exe base.Executor, redis *redis.Pool) (*Service, error) {
	jwtData, err := jwtSvc.AuthJWT(req)
	if err != nil {
		return nil, err
	}
	return &Service{
		JWTData:    *jwtData,
		exe:        exe,
		redis:      redis,
		mux:        &sync.Mutex{},
		roles:      nil, // loaded only when neeeded with function: IsAdmin
		privileges: nil, // loaded only when neeeded with functions: CheckAuthorization, GetPrivilegeByPriority
	}, nil
}

// CheckAuthorization checks if the user is authorized to execute the handler.
//
// Admin users are allowed to execute all the handlers.
func (svc *Service) CheckAuthorization(ctx context.Context, httMethod, httpRoute string) error {
	if ok, err := svc.IsAdmin(ctx); err != nil {
		return err
	} else if ok {
		return nil
	}
	privs, err := svc.getPrivileges(ctx)
	if err != nil {
		return err
	}
	for i := range privs {
		if privs[i].Method == httMethod && privs[i].Route == httpRoute {
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
