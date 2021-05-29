package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

type Roles []Role

//Role user role.
type Role struct {
	ID          string
	Name        string
	Description string
	Icon        string
	Privileges  Privileges
	CreatedAt   time.Time
	UpdatedAt   time.Time
	Status      base.Status
	base.RecordActions
}

func (roles Roles) Privileges() Privileges {
	var p Privileges
	for _, r := range roles {
		p = append(p, r.Privileges...)
	}
	return p
}
