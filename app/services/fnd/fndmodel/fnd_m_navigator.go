package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

type Navigator struct {
	ID            string     `rql:"filter,sort,column=nav_id"`
	Name          string     `rql:"filter,sort,column=nav_name"`
	Description   string     `rql:"filter,sort,column=nav_description"`
	Icon          string     `rql:"filter,sort,column=nav_icon"`
	Order         int        `rql:"filter,sort,column=nav_order"`
	URL           string     `rql:"filter,sort,column=nav_url"`
	Parent        *Navigator `rql:"filter,sort,datatype=string,column=nav_parent_id"`
	CreatedAt     time.Time  `rql:"filter,sort,column=nav_created_at"`
	UpdatedAt     time.Time  `rql:"filter,sort,column=nav_updated_at"`
	Status        base.Status
	RecordActions base.RecordActions
}
