package coremodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Actions slice.
type Actions []Action

// Action that can be executed on a module.
type Action struct {
	ModuleID      string
	ActionID      string
	Name          string
	Description   string
	Order         int
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        base.Status
	RecordActions base.RecordActions
}

func (a Action) Key() string {
	return a.ModuleID + "." + a.ActionID
}
