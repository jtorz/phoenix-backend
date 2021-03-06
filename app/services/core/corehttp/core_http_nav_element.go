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

// httpNavElement http handler component.
type httpNavElement struct {
	DB *sql.DB
}

func newHttpNavElement(db *sql.DB) httpNavElement {
	return httpNavElement{
		DB: db,
	}
}

// UpsertOrDeleteAll processes the all the records with the Create, Update or Delete actions.
//
// Delete: the record is deleted if the field NavElement.Deleted is true.
// Create: the record is created if doesn't exist.
// Update: the record is updated if already exists.
func (handler httpNavElement) UpsertOrDeleteAll() httphandler.HandlerFunc {
	return func(c *httphandler.Context) {
		req := coremodel.Navigator{}
		if c.BindJSON(&req) {
			return
		}
		biz := corebiz.NewBizNavElement()
		tx := c.BeginTx(handler.DB)
		err := biz.UpsertOrDeleteAll(c, tx.Tx, req)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.Status(http.StatusOK)
	}
}

// GetByID retrives the record information using its ID.
func (handler httpNavElement) GetByID() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":            nil,
		"Name":          nil,
		"Description":   nil,
		"Icon":          nil,
		"Order":         nil,
		"URL":           nil,
		"ParentID":      nil,
		"CreatedAt":     nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		id := c.Param("id")
		biz := corebiz.NewBizNavElement()
		rec, err := biz.GetByID(c, handler.DB, id)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(rec, resp)
	}
}

// ListAll returns the list of records that can be filtered by the user.
func (handler httpNavElement) ListAll() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":            nil,
		"Name":          nil,
		"Description":   nil,
		"Icon":          nil,
		"Order":         nil,
		"URL":           nil,
		"Children":      jsont.Recursive,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
		"IsAssigned":    nil,
	}
	return func(c *httphandler.Context) {
		var err error
		biz := corebiz.NewBizNavElement()
		recs, err := biz.ListAll(c, handler.DB, c.Param("roleID"))
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// SetStatus updates the logical status of the record.
func (handler httpNavElement) SetStatus(status base.Status) httphandler.HandlerFunc {
	type Req struct {
		ID        string    `binding:"required"`
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
		rec := coremodel.NavElement{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
			Status:    status,
		}
		biz := corebiz.NewBizNavElement()
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

// AssociateRole associates the nav element to the role.
func (handler httpNavElement) AssociateRole() httphandler.HandlerFunc {
	type Req struct {
		NavElementID string `binding:"required"`
		RoleID       string `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}

		biz := corebiz.NewBizNavElement()
		tx := c.BeginTx(handler.DB)
		err := biz.AssociateRole(c, tx.Tx, req.NavElementID, req.RoleID)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.Status(http.StatusOK)
	}
}

// DissociateRole dissociates the nav element from the role.
func (handler httpNavElement) DissociateRole() httphandler.HandlerFunc {
	type Req struct {
		NavElementID string `binding:"required"`
		RoleID       string `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}

		biz := corebiz.NewBizNavElement()
		tx := c.BeginTx(handler.DB)
		err := biz.DissociateRole(c, tx.Tx, req.NavElementID, req.RoleID)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.Status(http.StatusOK)
	}
}
