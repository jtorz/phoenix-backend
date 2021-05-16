package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

type Roles []Role

//Role user role.
type Role struct {
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Privileges    Privileges         `json:"privileges"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	Status        base.Status        `json:"status"`
	RecordActions base.RecordActions `json:"recordActions"`
}

func (roles Roles) Privileges() Privileges {
	var p Privileges
	for _, r := range roles {
		p = append(p, r.Privileges...)
	}
	return p
}
