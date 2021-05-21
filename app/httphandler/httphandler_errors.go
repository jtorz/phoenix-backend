package httphandler

import (
	"context"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
)

// ErrBadRequest reponds with status 400 if the error exists.
//
// returns true if the error exists.
func (c *Context) ErrBadRequest(err error, clientMsg string) bool {
	if err == nil {
		return false
	}
	if ctxinfo.LoggingLevel(c) == config.LogDebug {
		log.Println(err)
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

// HandleError handles the basic errors.
//
// returns true if the error is not nil.
func (c *Context) HandleError(err error) bool {
	if err == nil {
		return false
	}
	if baseerrors.IsErrInvalidData(err) {
		c.JSON(http.StatusBadRequest, err.Error()) // 400
		return true
	}
	if baseerrors.IsErrPrivilege(err) {
		c.JSON(http.StatusForbidden, err.Error()) // 403
		return true
	}
	if baseerrors.IsErrStatus(err) {
		c.JSON(http.StatusConflict, err.Error()) //409
		return true
	}
	if baseerrors.IsErrNotUpdated(err) || baseerrors.IsErrDuplicated(err) {
		c.JSON(http.StatusConflict, err.Error())
		return true
	}
	return c.UnexpectedError(err)
}

// UnexpectedError handles the error when the origin of such error is unknown
//
// returns true if the error is not nil.
func (c *Context) UnexpectedError(err error) bool {
	if err == nil {
		return false
	}
	if c.Request.Context().Err() == context.Canceled {
		return true
	}
	if ctxinfo.LoggingLevel(c) == config.LogDebug {
		log.Println(err)
	}
	c.Status(http.StatusInternalServerError)
	c.Set(KeyRequestError, err)
	return true
}
