package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Modules slice.
type Modules []Module

// Modules data.
type Module struct {
	ID            string    `rql:"filter,sort,column=mod_id"`
	Name          string    `rql:"filter,sort,column=mod_name"`
	Description   string    `rql:"filter,sort,column=mod_description"`
	Order         int       `rql:"filter,sort,column=mod_order"`
	Parent        *Module   `rql:"filter,sort,column=mod_parent_id,datatype=string"`
	CreatedAt     time.Time `rql:"filter,sort,column=mod_created_at"`
	UpdatedAt     time.Time `rql:"filter,sort,column=mod_updated_at"`
	Status        base.Status
	RecordActions base.RecordActions
}
