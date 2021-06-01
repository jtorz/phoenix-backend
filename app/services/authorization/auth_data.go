package authorization

import (
	"bytes"
	"context"
	"encoding/gob"
	"log"
	"strings"

	"github.com/doug-martin/goqu/v9"
	"github.com/gomodule/redigo/redis"

	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"

	//lint:ignore ST1001 dot import allowed only in dal packages for lex.
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

func (svc *Service) getRoles(ctx context.Context) ([]string, error) {
	svc.mux.Lock()
	defer svc.mux.Unlock()
	if svc.roles != nil {
		return svc.roles, nil
	}

	roles, err := svc.getRolesCache(ctx)
	if err != nil {
		if err != redis.ErrNil && ctxinfo.LogginAllowed(ctx, config.LogError) {
			log.Printf("REDIS GET USER ROLES: %s", err)
		}
	} else {
		svc.roles = roles
		return roles, nil
	}

	svc.roles, err = svc.getRolesDB(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		err := svc.saveRolesCache(svc.roles)
		if err != nil && err != redis.ErrNil && ctxinfo.LogginAllowed(ctx, config.LogError) {
			log.Printf("REDIS SAVE USER ROLES: %s", err)
		}
	}()

	return svc.roles, nil
}

func (svc *Service) getPrivileges(ctx context.Context) (privileges, error) {
	svc.mux.Lock()
	defer svc.mux.Unlock()
	if svc.privileges != nil {
		return svc.privileges, nil
	}

	privs, err := svc.getPrivilegesCache(ctx)
	if err != nil {
		if err != redis.ErrNil && ctxinfo.LogginAllowed(ctx, config.LogError) {
			log.Printf("REDIS GET USER PRIVILEGES: %s", err)
		}
	} else {
		svc.privileges = privs
		return privs, nil
	}

	svc.privileges, err = svc.getPrivilegesDB(ctx)
	if err != nil {
		return nil, err
	}

	go func() {
		err := svc.savePrivilegesCache(svc.privileges)
		if err != nil && err != redis.ErrNil && ctxinfo.LogginAllowed(ctx, config.LogError) {
			log.Printf("REDIS SAVE USER PRIVILEGES: %s", err)
		}
	}()

	return svc.privileges, nil
}

// roleNS role namespace.
const roleNS string = string(baseservice.RedisNSUserPriv) + "ROL:"

func (svc *Service) getRolesCache(ctx context.Context) ([]string, error) {
	conn := svc.redis.Get()
	defer conn.Close()
	conn.Flush()
	var roles []string

	data, err := conn.Do("GET", roleNS+svc.ID)
	if err != nil {
		return nil, err
	}
	b, err := redis.Bytes(data, err)
	if err != nil {
		return nil, err
	}
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&roles)
	if err != nil {
		return nil, err
	}
	return roles, nil
}

// privilegeNS privilege namespace.
const privilegekey string = string(baseservice.RedisNSUserPriv) + "PRV:"

func (svc *Service) getPrivilegesCache(ctx context.Context) ([]privilege, error) {
	conn := svc.redis.Get()
	defer conn.Close()
	var privs []privilege

	data, err := conn.Do("GET", privilegekey+svc.ID)
	if err != nil {
		return nil, err
	}
	b, err := redis.Bytes(data, err)
	if err != nil {
		return nil, err
	}
	dec := gob.NewDecoder(bytes.NewBuffer(b))
	err = dec.Decode(&privs)
	if err != nil {
		return nil, err
	}
	return privs, nil
}

func (svc *Service) saveRolesCache(roles []string) error {
	conn := svc.redis.Get()
	defer conn.Close()
	bytez := bytes.Buffer{}
	enc := gob.NewEncoder(&bytez)
	err := enc.Encode(roles)
	if err != nil {
		return err
	}

	_, err = conn.Do("SET", roleNS+svc.ID, bytez.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) savePrivilegesCache(privs []privilege) error {
	conn := svc.redis.Get()
	defer conn.Close()
	bytez := bytes.Buffer{}
	enc := gob.NewEncoder(&bytez)
	err := enc.Encode(privs)
	if err != nil {
		return err
	}
	_, err = conn.Do("SET", privilegekey+svc.ID, bytez.Bytes())
	if err != nil {
		return err
	}
	return nil
}

func (svc *Service) getRolesDB(ctx context.Context) ([]string, error) {
	recs := make([]string, 0)
	query := NewSelect(FndUserRole.UsrRoleID).
		From(T.FndUserRole).
		Where(goqu.C(FndUserRole.UsrUserID).Eq(svc.ID))

	rows, err := QueryContext(ctx, svc.exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		var rec string
		err := rows.Scan(
			&rec,
		)
		if err != nil {
			DebugErr(ctx, err)
			return nil, err
		}
		recs = append(recs, rec)
	}
	return recs, nil

}

func (svc *Service) getPrivilegesDB(ctx context.Context) ([]privilege, error) {
	recs := []privilege{}
	query := NewSelect(
		goqu.L(FndVPrivilegeRole.PrrModuleID+"||'.'||"+FndVPrivilegeRole.PrrActionID),
		CoalesceStr(FndVPrivilegeRole.PrrMethod),
		CoalesceStr(FndVPrivilegeRole.PrrRoute),
	).
		From(V.FndVPrivilegeRole).
		InnerJoin(goqu.T(T.FndUserRole), goqu.On(goqu.C(FndUserRole.UsrRoleID).Eq(goqu.C(FndVPrivilegeRole.PrrRoleID)))).
		Where(goqu.C(FndUserRole.UsrUserID).Eq(svc.ID)).GroupBy()

	rows, err := QueryContext(ctx, svc.exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rec := privilege{}
		err := rows.Scan(
			&rec.Key,
			&rec.Method,
			&rec.Route,
		)
		if err != nil {
			DebugErr(ctx, err)
			return nil, err
		}
		recs = append(recs, rec)
	}
	log.Println(strings.Repeat("*", 100))
	log.Println(recs)
	return recs, nil
}
