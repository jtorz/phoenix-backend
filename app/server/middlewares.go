package server

import (
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/services/authorization"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
)

type agent string

// configureMiddlewares configures the http middlewares.
func (server *Server) configureMiddlewares(r *gin.Engine, jwtSvc authorization.JWTService) {
	r.Use(gin.Recovery()) // default recovery
	r.Use(func(ginCtx *gin.Context) {
		ctxinfo.SetPrintLog(ginCtx, config.Mode(server.Config.AppMode))
	})
	if server.Config.AppModeDebug() {
		r.Use(gin.Logger())
	}

	server.serveStaticFiles(r)

	r.NoRoute(func(c *gin.Context) {
		if isAPIRoute(c) {
			//c.Set(controller.KeyRequestError, error404(c.Request.URL.String()))
			c.JSON(http.StatusNotFound, "Not found")
			return
		}
		c.Redirect(http.StatusFound, "/?redirect="+url.QueryEscape(c.Request.URL.String()))
	})

	// midleware used to store
	r.Use(func(ginCtx *gin.Context) {

		if isAPIPublicRoute(ginCtx) {
			ginCtx.Next()
			return
		}
		c := httphandler.New(ginCtx)

		auth, err := authorization.NewAuthService(c, jwtSvc, server.MainDB) // check if has jwt
		if err == nil {
			ctxinfo.SetAgent(c.Context, auth.ID, auth)
			c.Next()
			return
		}
		defer c.Abort()
		if err == baseerrors.ErrAuth {
			c.JSON(http.StatusUnauthorized, "unauthorized")
			return
		}

		if err == baseerrors.ErrPrivilege {
			c.JSON(http.StatusUnauthorized, "fobidden: "+c.Request.RequestURI)
			return
		}

	})
}

func (server *Server) serveStaticFiles(r *gin.Engine) {
	// TODO:
}

func isAPIRoute(c *gin.Context) bool {
	return strings.HasPrefix(c.Request.URL.String(), "/api/")
}

func isAPIAdminRoute(c *gin.Context) bool {
	return strings.HasPrefix(c.Request.URL.String(), "/api/admin/")
}

func isAPIPublicRoute(c *gin.Context) bool {
	if !isAPIRoute(c) {
		return true
	}
	return strings.HasPrefix(c.Request.URL.String(), "/api/public/")
}
