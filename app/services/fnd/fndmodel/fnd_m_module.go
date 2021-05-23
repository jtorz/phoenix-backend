package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Modules slice.
type Modules []Module

// Module del sistema
type Module struct {
	ID            string
	Name          string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        base.Status
	RecordActions base.RecordActions
}
