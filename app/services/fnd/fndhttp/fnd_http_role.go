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

/* CUT & PASTE TO fnd_http.go
httpRole := newHttpRole(s.DB)
{
	apiGroup.GET("/roles/role/:id/"remove_slash, httpRole.GetByID().Func())
	apiGroup.GET("/roles", httpRole.ListAll().Func())
	apiGroup.POST("/roles", httpRole.ListAll().Func())
	apiGroup.GET("/roles/active-records", httpRole.ListActive().Func())
	apiGroup.POST("/roles/active-records", httpRole.ListActive().Func())
	apiGroup.POST("/roles/role", httpRole.New().Func())
	apiGroup.PUT("/roles/role", httpRole.Edit().Func())
	apiGroup.PUT("/roles/role/validate", httpRole.SetStatus(base.StatusActive).Func())
	apiGroup.PUT("/roles/role/invalidate", httpRole.SetStatus(base.StatusInactive).Func())
	apiGroup.PUT("/roles/role/soft-delete", httpRole.SetStatus(base.StatusDroppped).Func())
	apiGroup.PUT("/roles/role/hard-delete", httpRole.Delete().Func())
}
*/

// httpRole http handler component.
type httpRole struct {
	DB *sql.DB
}

func newHttpRole(db *sql.DB) httpRole {
	return httpRole{
		DB: db,
	}
}

// GetByID retrives the record information using its ID.
func (handler httpRole) GetByID() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":            nil,
		"Name":          nil,
		"Description":   nil,
		"Icon":          nil,
		"CreatedAt":     nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		id := c.Param("id")
		biz := fndbiz.NewBizRole()
		rec, err := biz.GetByID(c, handler.DB, id)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(rec, resp)
	}
}

// ListAll returns the list of records that can be filtered by the user.
func (handler httpRole) ListAll() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":            nil,
		"Name":          nil,
		"Description":   nil,
		"Icon":          nil,
		"CreatedAt":     nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		biz := fndbiz.NewBizRole()
		recs, err := biz.List(c, handler.DB, false)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// ListActive returns the list of records that can be filtered by the user.
func (handler httpRole) ListActive() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":          nil,
		"Name":        nil,
		"Description": nil,
		"Icon":        nil,
	}
	return func(c *httphandler.Context) {
		biz := fndbiz.NewBizRole()
		recs, err := biz.List(c, handler.DB, true)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// New creates a new record.
func (handler httpRole) New() httphandler.HandlerFunc {
	type Req struct {
		Name        string `binding:"required"`
		Description string `binding:"required"`
		Icon        string `binding:"required"`
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

		rec := fndmodel.Role{
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
		}
		biz := fndbiz.NewBizRole()
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
func (handler httpRole) Edit() httphandler.HandlerFunc {
	type Req struct {
		ID          string    `binding:"required"`
		Name        string    `binding:"required"`
		Description string    `binding:"required"`
		Icon        string    `binding:"required"`
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
		rec := fndmodel.Role{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
			UpdatedAt:   req.UpdatedAt,
		}

		biz := fndbiz.NewBizRole()
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
func (handler httpRole) SetStatus(status base.Status) httphandler.HandlerFunc {
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
		rec := fndmodel.Role{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
			Status:    status,
		}
		biz := fndbiz.NewBizRole()
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
func (handler httpRole) Delete() httphandler.HandlerFunc {
	type Req struct {
		ID        string    `binding:"required"`
		UpdatedAt time.Time `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := fndmodel.Role{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
		}
		biz := fndbiz.NewBizRole()
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
