package corehttp

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/jtorz/jsont/v2"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/services/core/corebiz"
	"github.com/jtorz/phoenix-backend/app/services/core/coremodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// httpAction http handler component.
type httpAction struct {
	DB *sql.DB
}

func newHttpAction(db *sql.DB) httpAction {
	return httpAction{
		DB: db,
	}
}

// GetByID retrives the record information using its ID.
func (handler httpAction) GetByID() httphandler.HandlerFunc {
	resp := jsont.F{
		"ModuleID":      nil,
		"ActionID":      nil,
		"Name":          nil,
		"Description":   nil,
		"Order":         nil,
		"CreatedAt":     nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		moduleID := c.Param("moduleID")
		actionID := c.Param("actionID")
		biz := corebiz.NewBizAction()
		rec, err := biz.GetByID(c, handler.DB, moduleID, actionID)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(rec, resp)
	}
}

// ListAll returns the list of records that can be filtered by the user.
func (handler httpAction) ListAll() httphandler.HandlerFunc {
	resp := jsont.F{
		"ModuleID":      nil,
		"ActionID":      nil,
		"Name":          nil,
		"Description":   nil,
		"Order":         nil,
		"CreatedAt":     nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		moduleID := c.Param("moduleID")
		var err error
		qry := base.ClientQuery{OnlyActive: false}
		qry.RQL, err = c.GetRawData()
		if c.HandleError(err) {
			return
		}

		biz := corebiz.NewBizAction()
		recs, err := biz.List(c, handler.DB, qry, moduleID)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// ListActive returns the list of records that can be filtered by the user.
func (handler httpAction) ListActive() httphandler.HandlerFunc {
	resp := jsont.F{
		"ModuleID":    nil,
		"ActionID":    nil,
		"Name":        nil,
		"Description": nil,
		"Order":       nil,
	}
	return func(c *httphandler.Context) {
		moduleID := c.Param("moduleID")
		var err error
		qry := base.ClientQuery{OnlyActive: true}
		qry.RQL, err = c.GetRawData()
		if c.HandleError(err) {
			return
		}

		biz := corebiz.NewBizAction()
		recs, err := biz.List(c, handler.DB, qry, moduleID)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// New creates a new record.
func (handler httpAction) New() httphandler.HandlerFunc {
	type Req struct {
		ModuleID    string `binding:"required"`
		ActionID    string `binding:"required"`
		Name        string `binding:"required"`
		Description string `binding:"required"`
		Order       int
	}
	resp := jsont.F{
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}

		rec := coremodel.Action{
			ModuleID:    req.ModuleID,
			ActionID:    req.ActionID,
			Name:        req.Name,
			Description: req.Description,
			Order:       req.Order,
		}
		biz := corebiz.NewBizAction()
		tx := c.BeginTx(handler.DB)
		err := biz.New(c, tx.Tx, &rec)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.JSONWithFields(rec, resp)
	}
}

// Edit edits the record.
func (handler httpAction) Edit() httphandler.HandlerFunc {
	type Req struct {
		ModuleID    string `binding:"required"`
		ActionID    string `binding:"required"`
		Name        string `binding:"required"`
		Description string `binding:"required"`
		Order       int
		UpdatedAt   time.Time `binding:"required"`
	}
	resp := jsont.F{
		"UpdatedAt": nil,
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := coremodel.Action{
			ModuleID:    req.ModuleID,
			ActionID:    req.ActionID,
			Name:        req.Name,
			Description: req.Description,
			Order:       req.Order,
			UpdatedAt:   req.UpdatedAt,
		}

		biz := corebiz.NewBizAction()
		tx := c.BeginTx(handler.DB)
		err := biz.Edit(c, tx.Tx, &rec)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.JSONWithFields(rec, resp)
	}
}

// SetStatus updates the logical status of the record.
func (handler httpAction) SetStatus(status base.Status) httphandler.HandlerFunc {
	type Req struct {
		ModuleID  string    `binding:"required"`
		ActionID  string    `binding:"required"`
		UpdatedAt time.Time `binding:"required"`
	}
	resp := jsont.F{
		"UpdatedAt":     nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := coremodel.Action{
			ModuleID:  req.ModuleID,
			ActionID:  req.ActionID,
			UpdatedAt: req.UpdatedAt,
			Status:    status,
		}
		biz := corebiz.NewBizAction()
		tx := c.BeginTx(handler.DB)
		err := biz.SetStatus(c, tx.Tx, &rec)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.JSONWithFields(rec, resp)
	}
}

// Delete performs a physical delete of the record.
func (handler httpAction) Delete() httphandler.HandlerFunc {
	type Req struct {
		ModuleID  string    `binding:"required"`
		ActionID  string    `binding:"required"`
		UpdatedAt time.Time `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := coremodel.Action{
			ModuleID:  req.ModuleID,
			ActionID:  req.ActionID,
			UpdatedAt: req.UpdatedAt,
		}
		biz := corebiz.NewBizAction()
		tx := c.BeginTx(handler.DB)
		err := biz.Delete(c, tx.Tx, &rec)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.Status(http.StatusOK)
	}
}
