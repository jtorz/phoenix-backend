package mailhttp

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/jtorz/jsont/v2"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/services/mail/mailbiz"
	"github.com/jtorz/phoenix-backend/app/services/mail/mailmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// httpSender http handler component.
type httpSender struct {
	DB *sql.DB
}

func newHttpSender(db *sql.DB) httpSender {
	return httpSender{
		DB: db,
	}
}

// GetByID retrives the record information using its ID.
func (handler httpSender) GetByID() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":            nil,
		"Name":          nil,
		"Description":   nil,
		"Host":          nil,
		"Port":          nil,
		"User":          nil,
		"Password":      nil,
		"From":          nil,
		"CreatedAt":     nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		id := c.Param("id")
		biz := mailbiz.NewBizSender()
		rec, err := biz.GetByID(c, handler.DB, id)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(rec, resp)
	}
}

// ListAll returns the list of records that can be filtered by the user.
func (handler httpSender) ListAll() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":            nil,
		"Name":          nil,
		"Description":   nil,
		"Host":          nil,
		"Port":          nil,
		"User":          nil,
		"Password":      nil,
		"From":          nil,
		"CreatedAt":     nil,
		"UpdatedAt":     nil,
		"Status":        nil,
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		biz := mailbiz.NewBizSender()
		recs, err := biz.List(c, handler.DB, false)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// ListActive returns the list of records that can be filtered by the user.
func (handler httpSender) ListActive() httphandler.HandlerFunc {
	resp := jsont.F{
		"ID":          nil,
		"Name":        nil,
		"Description": nil,
		"Host":        nil,
		"Port":        nil,
		"User":        nil,
		"Password":    nil,
		"From":        nil,
	}
	return func(c *httphandler.Context) {
		biz := mailbiz.NewBizSender()
		recs, err := biz.List(c, handler.DB, true)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// New creates a new record.
func (handler httpSender) New() httphandler.HandlerFunc {
	type Req struct {
		ID          string `binding:"required"`
		Name        string `binding:"required"`
		Description string `binding:"required"`
		Host        string `binding:"required"`
		Port        int    `binding:"required"`
		User        string `binding:"required"`
		Password    string `binding:"required"`
		From        string `binding:"required"`
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

		rec := mailmodel.Sender{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			Host:        req.Host,
			Port:        req.Port,
			User:        req.User,
			Password:    req.Password,
			From:        req.From,
		}
		biz := mailbiz.NewBizSender()
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
func (handler httpSender) Edit() httphandler.HandlerFunc {
	type Req struct {
		ID          string    `binding:"required"`
		Name        string    `binding:"required"`
		Description string    `binding:"required"`
		Host        string    `binding:"required"`
		Port        int       `binding:"required"`
		User        string    `binding:"required"`
		Password    string    `binding:"required"`
		From        string    `binding:"required"`
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
		rec := mailmodel.Sender{
			ID:          req.ID,
			Name:        req.Name,
			Description: req.Description,
			Host:        req.Host,
			Port:        req.Port,
			User:        req.User,
			Password:    req.Password,
			From:        req.From,
			UpdatedAt:   req.UpdatedAt,
		}

		biz := mailbiz.NewBizSender()
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
func (handler httpSender) SetStatus(status base.Status) httphandler.HandlerFunc {
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
		rec := mailmodel.Sender{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
			Status:    status,
		}
		biz := mailbiz.NewBizSender()
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
func (handler httpSender) Delete() httphandler.HandlerFunc {
	type Req struct {
		ID        string    `binding:"required"`
		UpdatedAt time.Time `binding:"required"`
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := mailmodel.Sender{
			ID:        req.ID,
			UpdatedAt: req.UpdatedAt,
		}
		biz := mailbiz.NewBizSender()
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
