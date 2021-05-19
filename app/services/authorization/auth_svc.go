package authorization

import (
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

// Service authorization module service.
//
// Implements ctxinfo.Service interface.
type Service struct {
	//DB *sql.DB
	baseservice.JWTData
}

// NewAuthService creates a new Service.
func NewAuthService(c *httphandler.Context, jwtSvc JWTSvc, db *sql.DB) (*Service, error) {
	jwtData, err := jwtSvc.AuthJWT(c)
	if err != nil {
		return nil, err
	}

	svc := Service{
		JWTData: *jwtData,
	}

	privs, err := svc.GetPrivileges()
	if err != nil {
		return nil, err
	}
	for i := range privs {
		if privs[i].Method == c.Request.Method && privs[i].Route == c.FullPath() {
			return &svc, nil
		}
	}
	return nil, baseerrors.ErrPrivilege
}

func (svc *Service) GetPrivileges() ([]Privilege, error) {
	return []Privilege{}, nil
}

func (svc *Service) GetPrivilegeByPriority(privileges ...string) (string, error) {
	if len(privileges) > 0 {
		return privileges[0], nil
	}
	return "", nil
}
func (svc *Service) HasPrivilege(string) (bool, error) {
	return true, nil
}
func (svc *Service) IsAdmin() (bool, error) {
	return true, nil
}
