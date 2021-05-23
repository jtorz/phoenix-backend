package server

import (
	"log"
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

// configureMiddlewares configures the http middlewares.
func (server *Server) configureMiddlewares(r *gin.Engine, jwtSvc authorization.JWTSvc) {
	// Default gin recovery middleware
	r.Use(gin.Recovery())

	// Middleware used to add the app mode to the context.
	r.Use(func(ginCtx *gin.Context) {
		ctxinfo.SetLoggingLevel(ginCtx, config.LogginLvl(server.Config.LoggingLevel))
	})

	// gin.logger middleware added only on debug mode.
	if config.LogDebug >= server.Config.LoggingLevel {
		r.Use(gin.Logger())
	}

	// static file server
	server.serveStaticFiles(r)

	// If the route doesn't exists in the api a status 404 is returned,
	// If the requested resource is not from the api the client is redirected to the root.
	// In order to serve the frontend files, and it can search the original requested resource.
	r.NoRoute(func(c *gin.Context) {
		if isAPIRoute(c) {
			//c.Set(controller.KeyRequestError, error404(c.Request.URL.String()))
			c.JSON(http.StatusNotFound, "Not found")
			return
		}
		c.Redirect(http.StatusFound, "/?redirect="+url.QueryEscape(c.Request.URL.String()))
	})

	// Auhtentication and Authorization middleware.
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

		c.Status(http.StatusInternalServerError)
		log.Printf("uknown error: %s", err)
	})
}

func (server *Server) serveStaticFiles(r *gin.Engine) {
	// TODO:
}

func isAPIRoute(c *gin.Context) bool {
	return strings.HasPrefix(c.Request.URL.String(), "/api/")
}

func isAPIPublicRoute(c *gin.Context) bool {
	if !isAPIRoute(c) {
		return true
	}
	return strings.HasPrefix(c.Request.URL.String(), "/api/public/")
}
