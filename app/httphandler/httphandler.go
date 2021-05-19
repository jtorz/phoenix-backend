package httphandler

import (
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
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
