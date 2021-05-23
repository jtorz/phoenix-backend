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

// httpNavigator http handler component.
type httpNavigator struct {
	DB *sql.DB
}

func newHttpNavigator(db *sql.DB) httpNavigator {
	return httpNavigator{
		DB: db,
	}
}

// GetByID retrives the record information using its ID.
func (handler httpNavigator) GetByID() httphandler.HandlerFunc {
	return func(c *httphandler.Context) {
		biz := fndbiz.NewBizNavigator()
		rec, err := biz.GetByID(c, handler.DB, c.Param("id"))
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(rec, jsont.F{
			"Id":            nil,
			"Name":          nil,
			"Description":   nil,
			"Icon":          nil,
			"Order":         nil,
			"Url":           nil,
			"Parent":        jsont.Recursive,
			"Status":        nil,
			"UpdatedAt":     nil,
			"RecordActions": nil,
		})
	}
}

// List returns the list of records that can be filtered by the user.
func (handler httpNavigator) List(queryOnlyActive bool) httphandler.HandlerFunc {
	return func(c *httphandler.Context) {
		var err error
		qry := base.ClientQuery{OnlyActive: queryOnlyActive}
		qry.RQL, err = c.GetRawData()
		if c.HandleError(err) {
			return
		}

		biz := fndbiz.NewBizNavigator()
		recs, err := biz.List(c, handler.DB, qry)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, jsont.F{
			"Id":            nil,
			"Name":          nil,
			"Description":   nil,
			"Icon":          nil,
			"Order":         nil,
			"Url":           nil,
			"Parent":        jsont.Recursive,
			"Status":        nil,
			"UpdatedAt":     nil,
			"RecordActions": nil,
		})
	}
}

// New creates a new record.
func (handler httpNavigator) New() httphandler.HandlerFunc {
	type Req struct {
		Name        string `binding:"required"`
		Description string `binding:"required"`
		Icon        string `binding:"required"`
		Order       int    `binding:"required"`
		URL         string `binding:"required"`
		ParentID    string `binding:""`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}

		rec := fndmodel.Navigator{
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
			Order:       req.Order,
			URL:         req.URL,
			Parent:      &fndmodel.Navigator{ID: req.ParentID},
		}
		biz := fndbiz.NewBizNavigator()
		tx := c.BeginTx(handler.DB)
		err := biz.New(c, tx.Tx, &rec)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.JSONWithFields(rec, jsont.F{
			"Id":            nil,
			"Status":        nil,
			"UpdatedAt":     nil,
			"RecordActions": nil,
		})
	}
}

// Edit edits the record.
func (handler httpNavigator) Edit() httphandler.HandlerFunc {
	type Req struct {
		ID          string    ` binding:"required"`
		Name        string    `binding:"required"`
		Description string    `binding:"required"`
		Icon        string    `binding:"required"`
		Order       int       `binding:"required"`
		URL         string    `binding:"required"`
		ParentID    string    `binding:""`
		UpdatedAt   time.Time `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}

		rec := fndmodel.Navigator{
			Name:        req.Name,
			Description: req.Description,
			Icon:        req.Icon,
			Order:       req.Order,
			URL:         req.URL,
			Parent:      &fndmodel.Navigator{ID: req.ParentID},
		}
		biz := fndbiz.NewBizNavigator()
		tx := c.BeginTx(handler.DB)
		err := biz.Edit(c, tx.Tx, &rec)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.JSONWithFields(rec, jsont.F{
			"UpdatedAt": nil,
		})
	}
}

// SetStatus updates the logical status of the record.
func (handler httpNavigator) SetStatus(status base.Status) httphandler.HandlerFunc {
	type Req struct {
		ID        string    ` binding:"required"`
		UpdatedAt time.Time `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := fndmodel.Navigator{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
			Status:    status,
		}
		biz := fndbiz.NewBizNavigator()
		tx := c.BeginTx(handler.DB)
		err := biz.SetStatus(c, tx.Tx, &rec)
		if c.HandleError(err) {
			tx.Rollback(c)
			return
		}
		tx.Commit(c)
		c.JSONWithFields(rec, jsont.F{
			"UpdatedAt":     nil,
			"RecordActions": nil,
		})
	}
}

// Delete performs a physical delete of the record.
func (handler httpNavigator) Delete() httphandler.HandlerFunc {
	type Req struct {
		ID        string    ` binding:"required"`
		UpdatedAt time.Time `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := fndmodel.Navigator{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
		}
		biz := fndbiz.NewBizNavigator()
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
