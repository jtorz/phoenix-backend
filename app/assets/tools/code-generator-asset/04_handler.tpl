{{define "Param"}}
{{- if eq .GoDataType "string"}}
	{{.GoVarName}} := c.Param("id")
{{- else}}
	{{.GoVarName}}, isErr := c.Param{{.GoDataType | upperfirst}}("{{.GoVarName}}")
	if isErr{
		return
	}
{{- end}}
{{end}}

package {{$.ServiceAbbr | lowercase}}http

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/jtorz/jsont/v2"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/services/{{$.ServiceAbbr | lowercase}}/{{$.ServiceAbbr | lowercase}}biz"
	"github.com/jtorz/phoenix-backend/app/services/{{$.ServiceAbbr | lowercase}}/{{$.ServiceAbbr | lowercase}}model"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// http{{$.Entity.GoStruct}} http handler component.
type http{{$.Entity.GoStruct}} struct {
	DB *sql.DB
}

func newHttp{{$.Entity.GoStruct}}(db *sql.DB) http{{$.Entity.GoStruct}} {
	return http{{$.Entity.GoStruct}}{
		DB: db,
	}
}

// GetByID retrives the record information using its ID.
func (handler http{{$.Entity.GoStruct}}) GetByID() httphandler.HandlerFunc {
	resp := jsont.F{
		{{- range $Col := $.Entity.Columns}}
			"{{.GoField}}":nil,
		{{- end}}
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		{{- range $Col := $.Entity.Columns}}
		{{- if $Col.IsPK}}
			{{- template "Param" $Col}}
		{{- end}}
		{{- end}}biz := {{$.ServiceAbbr | lowercase}}biz.NewBiz{{$.Entity.GoStruct}}()
		rec, err := biz.GetByID(c, handler.DB, {{- range $Col := $.Entity.Columns}}
		{{- if $Col.IsPK}}{{.GoVarName}}, {{- end}}
		{{- end}})
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(rec, resp)
	}
}

// ListAll returns the list of records that can be filtered by the user.
func (handler http{{$.Entity.GoStruct}}) ListAll() httphandler.HandlerFunc {
	resp := jsont.F{
		{{- range $Col := $.Entity.Columns}}
			"{{.GoField}}":nil,
		{{- end}}
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		var err error
		qry := base.ClientQuery{OnlyActive: false}
		qry.RQL, err = c.GetRawData()
		if c.HandleError(err) {
			return
		}

		biz := {{$.ServiceAbbr | lowercase}}biz.NewBiz{{$.Entity.GoStruct}}()
		recs, err := biz.List(c, handler.DB, qry)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}


// ListActive returns the list of records that can be filtered by the user.
func (handler http{{$.Entity.GoStruct}}) ListActive() httphandler.HandlerFunc {
	resp := jsont.F{
		{{- range $Col := $.Entity.Columns}}
			{{if and (ne $Col.GoField "Status") (ne $Col.GoField "UpdatedAt") (ne $Col.GoField "CreatedAt")}} "{{.GoField}}":nil, {{end}}
		{{- end}}
	}
	return func(c *httphandler.Context) {
		var err error
		qry := base.ClientQuery{OnlyActive: true}
		qry.RQL, err = c.GetRawData()
		if c.HandleError(err) {
			return
		}

		biz := {{$.ServiceAbbr | lowercase}}biz.NewBiz{{$.Entity.GoStruct}}()
		recs, err := biz.List(c, handler.DB, qry)
		if c.HandleError(err) {
			return
		}
		c.JSONWithFields(recs, resp)
	}
}

// New creates a new record.
func (handler http{{$.Entity.GoStruct}}) New() httphandler.HandlerFunc {
	type Req struct {
	{{- range $Col := $.Entity.Columns}}
		{{- if and (not $Col.IsPK) (ne $Col.GoField "Status") (ne $Col.GoField "UpdatedAt") (ne $Col.GoField "CreatedAt")}}
			{{$Col.GoField}}  {{$Col.GoDataType}} `binding:"required"`
		{{- end}}
	{{- end}}
	}
	resp := jsont.F{
		{{- range $Col := $.Entity.Columns}}
			{{- if or ($Col.IsPK) (eq $Col.GoField "Status") (eq $Col.GoField "UpdatedAt")}}
				"{{$Col.GoField}}": nil,
			{{- end}}
		{{- end}}
		"RecordActions": nil,
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}

		rec := {{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}}{
		{{- range $Col := $.Entity.Columns}}
			{{- if and (not $Col.IsPK) (ne $Col.GoField "Status") (ne $Col.GoField "UpdatedAt") (ne $Col.GoField "CreatedAt")}}
				{{$Col.GoField}}: req.{{$Col.GoField}},
			{{- end}}
		{{- end}}
		}
		biz := {{$.ServiceAbbr | lowercase}}biz.NewBiz{{$.Entity.GoStruct}}()
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
func (handler http{{$.Entity.GoStruct}}) Edit() httphandler.HandlerFunc {
	type Req struct {
	{{- range $Col := $.Entity.Columns}}
		{{- if and (ne $Col.GoField "Status") (ne $Col.GoField "CreatedAt")}}
			{{$Col.GoField}}  {{$Col.GoDataType}} `binding:"required"`
		{{- end}}
	{{- end}}
	}
	resp := jsont.F{
		"UpdatedAt": nil,
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := {{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}}{
		{{- range $Col := $.Entity.Columns}}
			{{- if and (ne $Col.GoField "Status") (ne $Col.GoField "CreatedAt")}}
				{{$Col.GoField}}: req.{{$Col.GoField}},
			{{- end}}
		{{- end}}
		}

		biz := {{$.ServiceAbbr | lowercase}}biz.NewBiz{{$.Entity.GoStruct}}()
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
func (handler http{{$.Entity.GoStruct}}) SetStatus(status base.Status) httphandler.HandlerFunc {
	type Req struct {
	{{- range $Col := $.Entity.Columns}}
		{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
			{{$Col.GoField}}  {{$Col.GoDataType}} `binding:"required"`
		{{- end}}
	{{- end}}
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
		rec := {{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}}{
		{{- range $Col := $.Entity.Columns}}
			{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
				{{$Col.GoField}}: req.{{$Col.GoField}},
			{{- end}}
		{{- end}}
			Status:    status,
		}
		biz := {{$.ServiceAbbr | lowercase}}biz.NewBiz{{$.Entity.GoStruct}}()
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
func (handler http{{$.Entity.GoStruct}}) Delete() httphandler.HandlerFunc {
	type Req struct {
	{{- range $Col := $.Entity.Columns}}
		{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
			{{$Col.GoField}}  {{$Col.GoDataType}} `binding:"required"`
		{{- end}}
	{{- end}}
	}
	return func(c *httphandler.Context) {
		req := Req{}
		if c.BindJSON(&req) {
			return
		}
		rec := {{$.ServiceAbbr | lowercase}}model.{{$.Entity.GoStruct}}{
		{{- range $Col := $.Entity.Columns}}
			{{- if or ($Col.IsPK) (eq $Col.GoField "UpdatedAt")}}
				{{$Col.GoField}}: req.{{$Col.GoField}},
			{{- end}}
		{{- end}}
		}
		biz := {{$.ServiceAbbr | lowercase}}biz.NewBiz{{$.Entity.GoStruct}}()
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
