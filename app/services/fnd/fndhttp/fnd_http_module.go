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

// httpModule http handler component.
type httpModule struct {
	DB *sql.DB
}

func newHttpModule(db *sql.DB) httpModule {
	return httpModule{
		DB: db,
	}
}

// GetByID retrives the record information using its ID.
func (handler httpModule) GetByID() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":            nil,
		"Name":          nil,
		"Description":   nil,
		"Order":         nil,
		"ParentID":      nil,
		"CreatedAt":     nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		id := c.Param("id")
		biz := fndbiz.NewBizModule()
		rec, err := biz.GetByID(c, handler.DB, id)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(rec, resp)
	}
}

// ListAll returns the list of records that can be filtered by the user.
func (handler httpModule) ListAll() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":            nil,
		"Name":          nil,
		"Description":   nil,
		"Order":         nil,
		"ParentID":      nil,
		"CreatedAt":     nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		var err error
		qry := base.ClientQuery{OnlyActive: false}
		qry.RQL, err = c.GetRawData()
		if c.HandleError(err) {
			return
		}

		biz := fndbiz.NewBizModule()
		recs, err := biz.List(c, handler.DB, qry)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// ListActive returns the list of records that can be filtered by the user.
func (handler httpModule) ListActive() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":          nil,
		"Name":        nil,
		"Description": nil,
		"Order":       nil,
		"ParentID":    nil,
	}
	return func(c *httphandler.Context) {
		var err error
		qry := base.ClientQuery{OnlyActive: true}
		qry.RQL, err = c.GetRawData()
		if c.HandleError(err) {
			return
		}

		biz := fndbiz.NewBizModule()
		recs, err := biz.List(c, handler.DB, qry)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// New creates a new record.
func (handler httpModule) New() httphandler.HandlerFunc {
	type Req struct {
		ID          string `binding:"required"`
		Name        string `binding:"required"`
		Description string `binding:"required"`
		Order       int
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

		rec := fndmodel.Module{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			Order:       req.Order,
			Parent:      &fndmodel.Module{ID: req.ParentID},
		}
		biz := fndbiz.NewBizModule()
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
func (handler httpModule) Edit() httphandler.HandlerFunc {
	type Req struct {
		ID          string `binding:"required"`
		Name        string `binding:"required"`
		Description string `binding:"required"`
		Order       int
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
		rec := fndmodel.Module{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			Order:       req.Order,
			Parent:      &fndmodel.Module{ID: req.ParentID},
			UpdatedAt:   req.UpdatedAt,
		}

		biz := fndbiz.NewBizModule()
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
func (handler httpModule) SetStatus(status base.Status) httphandler.HandlerFunc {
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
		rec := fndmodel.Module{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
			Status:    status,
		}
		biz := fndbiz.NewBizModule()
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
func (handler httpModule) Delete() httphandler.HandlerFunc {
	type Req struct {
		ID        string    `binding:"required"`
		UpdatedAt time.Time `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := fndmodel.Module{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
		}
		biz := fndbiz.NewBizModule()
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
