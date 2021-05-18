package httphandler

import (
	"context"
	"log"
	golog "log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
)

// ErrBadRequest reponds with status 400 with a msg only if err != nil.
func (c *Context) ErrBadRequest(err error, clientMsg string) bool {
	if err == nil {
		return false
	}
	if ctxinfo.PrintLog(c) {
		golog.Println(err)
	}
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": clientMsg,
	})
	return true
}

// ErrBadRequestMsg reponds with status 400 with a msg.
func (c *Context) ErrBadRequestMsg(clientMsg string) {
	c.JSON(http.StatusBadRequest, gin.H{
		"msg": clientMsg,
	})
}

// HandleError handles the error with the basic.
// returns true if an error occured
func (c *Context) HandleError(err error) bool {
	if err == nil {
		return false
	}
	if baseerrors.IsErrPrivilege(err) {
		c.Status(http.StatusForbidden) // 403
		return true
	}
	if baseerrors.IsErrStatus(err) {
		c.JSON(http.StatusConflict, "status") //409
		return true
	}
	if baseerrors.IsErrNotUpdated(err) {
		c.JSON(http.StatusConflict, "not updated")
		return true
	}
	return c.UnexpectedError(err)
}

// UnexpectedError si hay error regresa true, responde con StatusInternalServerError
// y realiza el log del error
//
// Esta funcion solo debe llamarse cuando el origen del error es desconocido o no esta siendo manejado
func (c *Context) UnexpectedError(err error) bool {
	if err == nil {
		return false
	}
	if c.Request.Context().Err() == context.Canceled {
		return true
	}
	if ctxinfo.PrintLog(c) {
		log.Println(err)
	}
	c.Status(http.StatusInternalServerError)
	c.Set(KeyRequestError, err)
	return true
}
