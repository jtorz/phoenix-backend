package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Actions slice.
type Actions []Action

// Action that can be executed on a module.
type Action struct {
	Module        Module             `json:"module"`
	ID            string             `json:"id"`
	Name          string             `json:"name"`
	Description   string             `json:"description"`
	Order         int                `json:"order"`
	Route         string             `json:"route"`
	Method        string             `json:"method"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	Status        base.Status        `json:"status"`
	RecordActions base.RecordActions `json:"recordActions"`
}
