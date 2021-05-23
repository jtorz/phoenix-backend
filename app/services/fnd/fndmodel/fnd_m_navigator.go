package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

type Navigator struct {
	ID            string
	Name          string
	Description   string
	Icon          string
	Order         int
	URL           string
	Parent        *Navigator
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        base.Status
	RecordActions base.RecordActions
}
