package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Component system component.
type Component struct {
	ID            int                `json:"id"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	Status        base.Status        `json:"status"`
	RecordActions base.RecordActions `json:"recordActions"`
}
