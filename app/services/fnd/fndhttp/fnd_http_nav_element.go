package fndhttp

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/jtorz/jsont/v2"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndbiz"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
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

// New creates a new record.
func (handler httpNavElement) UpsertAll() httphandler.HandlerFunc {
	return func(c *httphandler.Context) {
		req := fndmodel.Navigator{}
		if c.BindJSON(&req) {
			return
		}
		biz := fndbiz.NewBizNavElement()
		tx := c.BeginTx(handler.DB)
		err := biz.UpsertAll(c, tx.Tx, req)
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
		biz := fndbiz.NewBizNavElement()
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
		biz := fndbiz.NewBizNavElement()
		recs, err := biz.ListAll(c, handler.DB, c.Param("roleID"))
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// New creates a new record.
func (handler httpNavElement) New() httphandler.HandlerFunc {
	type Req struct {
		ID          string `binding:"required"`
		Name        string `binding:"required"`
		Description string `binding:"required"`
		Icon        string `binding:"required"`
		Order       int
		URL         string `binding:"required"`
		ParentID    string
	}
	resp := jsont.F{
		"ID":            nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}

		rec := fndmodel.NavElement{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
			Order:       req.Order,
			URL:         req.URL,
			ParentID:    req.ParentID,
		}
		biz := fndbiz.NewBizNavElement()
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
func (handler httpNavElement) Edit() httphandler.HandlerFunc {
	type Req struct {
		ID          string `binding:"required"`
		Name        string `binding:"required"`
		Description string `binding:"required"`
		Icon        string `binding:"required"`
		Order       int
		URL         string `binding:"required"`
		ParentID    string
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
		rec := fndmodel.NavElement{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
			Order:       req.Order,
			URL:         req.URL,
			ParentID:    req.ParentID,
			UpdatedAt:   req.UpdatedAt,
		}

		biz := fndbiz.NewBizNavElement()
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
		rec := fndmodel.NavElement{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
			Status:    status,
		}
		biz := fndbiz.NewBizNavElement()
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
func (handler httpNavElement) Delete() httphandler.HandlerFunc {
	type Req struct {
		ID        string    `binding:"required"`
		UpdatedAt time.Time `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := fndmodel.NavElement{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
		}
		biz := fndbiz.NewBizNavElement()
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

		biz := fndbiz.NewBizNavElement()
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

		biz := fndbiz.NewBizNavElement()
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
