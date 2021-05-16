package httphandler

import (
	golog "log"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/locales/es"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	es_translations "github.com/go-playground/validator/v10/translations/es"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
)

var (
	uni   *ut.UniversalTranslator
	trans ut.Translator
)

func init() {
	uni = ut.New(es.New())

	v := binding.Validator.Engine().(*validator.Validate)
	trans, _ = uni.GetTranslator("es")
	es_translations.RegisterDefaultTranslations(v, trans)
}

// Handler gin.context wrapper.
type Handler struct {
	*gin.Context
}

// New wraps the gin.Context.
func New(c *gin.Context) *Handler {
	return &Handler{c}
}

// HandlerFunc defines the handler used by gin middleware as return value.
type HandlerFunc func(*Handler)

// Func returns the original gin.HandlerFunc.
func (h HandlerFunc) Func() gin.HandlerFunc {
	return func(c *gin.Context) {
		h(&Handler{c})
	}
}

// BindJSON calls the gin.Context.Bind method.
//
// If an error occurs during the process a StatusBadRequest is responded to the client,
// and returns true.
func (handler *Handler) BindJSON(v interface{}) bool {
	err := handler.ShouldBindWith(v, binding.JSON)
	if err == nil {
		return false
	}

	e, ok := err.(validator.ValidationErrors)
	if !ok {
		handler.ErrBadRequest(err, "malformed JSON")
		return true
	}

	for _, v := range e.Translate(trans) {
		handler.ErrBadRequestMsg(v)
		handler.Abort()
		return true
	}

	handler.ErrBadRequest(e, e.Error())
	handler.Abort()
	return true
}

// ParamInt returns the value of the URL param converted to int.
//
// If an error occurs during the conversion a StatusBadRequest is responded to the client,
// and returns true.
func (handler *Handler) ParamInt(paramName string) (int, bool) {
	num, err := strconv.Atoi(handler.Param(paramName))
	if handler.ErrBadRequest(err, paramName+" is not a number") {
		return 0, true
	}
	return num, false
}

// ParamInt64 returns the value of the URL param converted to int64.
//
// If an error occurs during the conversion a StatusBadRequest is responded to the client,
// and returns true.
func (handler *Handler) ParamInt64(paramName string) (int64, bool) {
	num, err := strconv.ParseInt(handler.Param(paramName), 10, 64)
	if handler.ErrBadRequest(err, paramName+" is not a number") {
		return 0, true
	}
	return num, false
}

// ErrBadRequest reponds with status 400 with a msg only if err != nil.
func (handler *Handler) ErrBadRequest(err error, clientMsg string) bool {
	if err == nil {
		return false
	}
	if ctxinfo.PrintLog(handler) {
		golog.Println(err)
	}
	handler.JSON(http.StatusBadRequest, gin.H{
		"msg": clientMsg,
	})
	return true
}

// ErrBadRequestMsg reponds with status 400 with a msg.
func (handler *Handler) ErrBadRequestMsg(clientMsg string) {
	handler.JSON(http.StatusBadRequest, gin.H{
		"msg": clientMsg,
	})
}
