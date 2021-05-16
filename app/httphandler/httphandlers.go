package httphandler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// Secret wraps the controller function to set the flag that marks the request as secret
//
// A secret request avoids the logging of sensible data.
func Secret(f HandlerFunc) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(KeyRequestSecret, true)
		f(New(c))
	}
}

// NotImplementedGin simple gin.HandlerFunc that returns http.StatusNotImplemented status.
func NotImplementedGin(c *gin.Context) {
	c.Status(http.StatusNotImplemented)
}
