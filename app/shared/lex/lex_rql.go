// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package lex

import (
	"fmt"

	//"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	rqlgq "github.com/jtorz/phoenix-backend/utils/rql-goqu"
)

var rql map[string]*rqlgq.Parser

func init() {
	rql = make(map[string]*rqlgq.Parser)

	initParser(T.FndAccountAccess, FndAccountAccess)
	initParser(T.FndAction, FndAction)
	initParser(T.FndModule, FndModule)
	initParser(T.FndNavigator, FndNavigator)
	initParser(T.FndPassword, FndPassword)
	initParser(T.FndPrivilege, FndPrivilege)
	initParser(T.FndRole, FndRole)
	initParser(T.FndRoleNavigator, FndRoleNavigator)
	initParser(T.FndUser, FndUser)
	initParser(T.FndUserRole, FndUserRole)

	initParser(V.FndVPrivilegeRole, FndVPrivilegeRole)
}

func initParser(tablename string, model interface{}) {
	p, err := rqlgq.NewParser(rqlgq.Config{Model: model})
	if err != nil {
		panic(fmt.Sprintf("can't init rql config %s", err.Error()))
	}
	rql[tablename] = p
}

// ParseFilter parser the client query for the table fnd_account_access.
func (TableFndAccountAccess) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndAccountAccess, qry)
}

// ParseFilter parser the client query for the table fnd_action.
func (TableFndAction) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndAction, qry)
}

// ParseFilter parser the client query for the table fnd_module.
func (TableFndModule) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndModule, qry)
}

// ParseFilter parser the client query for the table fnd_navigator.
func (TableFndNavigator) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndNavigator, qry)
}

// ParseFilter parser the client query for the table fnd_password.
func (TableFndPassword) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndPassword, qry)
}

// ParseFilter parser the client query for the table fnd_privilege.
func (TableFndPrivilege) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndPrivilege, qry)
}

// ParseFilter parser the client query for the table fnd_role.
func (TableFndRole) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndRole, qry)
}

// ParseFilter parser the client query for the table fnd_role_navigator.
func (TableFndRoleNavigator) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndRoleNavigator, qry)
}

// ParseFilter parser the client query for the table fnd_user.
func (TableFndUser) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndUser, qry)
}

// ParseFilter parser the client query for the table fnd_user_role.
func (TableFndUserRole) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(T.FndUserRole, qry)
}

// ParseFilter parser the client query for the table fnd_v_privilege_role.
func (ViewFndVPrivilegeRole) ParseFilter(qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	return ParseFilter(V.FndVPrivilegeRole, qry)
}

// ParseFilter parser the client query for the given table.
func ParseFilter(tablename string, qry base.ClientQuery) (rqlParams *rqlgq.Params, err error) {
	if len(qry.RQL) == 0 {
		return &rqlgq.Params{
			FilterExp: exp.NewExpressionList(exp.AndType),
		}, nil
	}
	p, err := rql[tablename].Parse(qry.RQL)
	if err != nil {
		return nil, err
	}
	return p, nil
}
