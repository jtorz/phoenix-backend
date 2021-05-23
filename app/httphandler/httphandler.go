package httphandler

import (
	"log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	"github.com/jtorz/jsont/v2"
	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
)

var (
	trans ut.Translator
)

func init() {
	v := binding.Validator.Engine().(*validator.Validate)
	trans, _ = ut.New(es.New()).GetTranslator("es")
	es_translations.RegisterDefaultTranslations(v, trans)
}

// Context gin.context wrapper created to add extra functions.
type Context struct {
	*gin.Context
}

// New wraps the gin.Context.
func New(c *gin.Context) *Context {
	return &Context{c}
}

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Context)

// Func returns the original gin.HandlerFunc.
func (h HandlerFunc) Func() gin.HandlerFunc {
	return func(c *gin.Context) {
		h(&Context{c})
	}
}

// BindJSON calls the gin.Context.Bind method.
//
// If an error occurs during the process a StatusBadRequest is responded to the client.
//
// returns true if an error occured.
func (c *Context) BindJSON(v interface{}) bool {
	err := c.ShouldBindWith(v, binding.JSON)
	if err == nil {
		return false
	}

	e, ok := err.(validator.ValidationErrors)
	if !ok {
		c.ErrBadRequest(err, "malformed JSON")
		return true
	}

	for _, v := range e.Translate(trans) {
		c.ErrBadRequestMsg(v)
		c.Abort()
		return true
	}

	c.ErrBadRequest(e, e.Error())
	c.Abort()
	return true
}

// ParamInt returns the value of the URL param converted to int.
//
// If an error occurs during the conversion a StatusBadRequest is responded to the client.
//
// returns true if an error occured.
func (c *Context) ParamInt(paramName string) (int, bool) {
	num, err := strconv.Atoi(c.Param(paramName))
	if c.ErrBadRequest(err, paramName+" is not a number") {
		return 0, true
	}
	return num, false
}

// ParamInt64 returns the value of the URL param converted to int64.
//
// If an error occurs during the conversion a StatusBadRequest is responded to the client.
//
// returns true if an error occured.
func (c *Context) ParamInt64(paramName string) (int64, bool) {
	num, err := strconv.ParseInt(c.Param(paramName), 10, 64)
	if c.ErrBadRequest(err, paramName+" is not a number") {
		return 0, true
	}
	return num, false
}

// JSONWithFields responds with the JSON encoding of v using the whitelist of fields to include.
func (c *Context) JSONWithFields(v interface{}, fields jsont.F) {
	json, err := jsont.MarshalFields(v, fields)
	if err != nil {
		if ctxinfo.LogginAllowed(c, config.LogError) {
			log.Printf("Marshaling response error %s\n", err)
		}
		c.JSON(http.StatusInternalServerError, http.StatusText(http.StatusInternalServerError))
		return
	}
	c.Header("Content-Type", "application/json; charset=utf-8")
	c.String(http.StatusOK, string(json))
}
