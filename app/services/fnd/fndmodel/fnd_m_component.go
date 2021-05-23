package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Component system component.
type Component struct {
	ID            int
	Name          string
	Description   string
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        base.Status
	RecordActions base.RecordActions
}
