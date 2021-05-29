package server

import (
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/services/agentinfo"
	"github.com/jtorz/phoenix-backend/app/services/authorization"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
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
		c := httphandler.New(ginCtx)
		authSvc, err := authorization.NewAuthService(c, jwtSvc, server.MainDB) // check if has jwt

		if isAPIPublicRoute(ginCtx) {
			var agent *baseservice.Agent
			if err != nil {
				agent = baseservice.NewAgentAnonym()
			} else {
				agentinfoSvc := agentinfo.NewService(server.MainDB, authSvc.ID)
				ctxinfo.SetAgent(c.Context, baseservice.NewAgent(authSvc.ID, agentinfoSvc, authSvc))
			}
			ctxinfo.SetAgent(c.Context, agent)
			ginCtx.Next()
			return
		}

		if err != nil {
			defer c.Abort()
			if err == baseerrors.ErrAuth {
				c.JSON(http.StatusUnauthorized, "unauthorized")
				return
			}

			if err == baseerrors.ErrPrivilege {
				c.JSON(http.StatusForbidden, "fobidden: "+c.Request.RequestURI)
				return
			}

			c.Status(http.StatusUnauthorized)
			log.Printf("uknown error: %s", err)
			return
		}

		agentinfoSvc := agentinfo.NewService(server.MainDB, authSvc.ID)
		ctxinfo.SetAgent(c.Context, baseservice.NewAgent(authSvc.ID, agentinfoSvc, authSvc))
		c.Next()
	})
}

func (server *Server) serveStaticFiles(r *gin.Engine) {
	r.Use(static.Serve("/", static.LocalFile(server.Config.AppPathEnv+"/web/dist", true)))
	fmt.Println("TODO: serve dynamic assets")
	//path := conf.GetDirWebAssets()
	//files.CreateDirPanic(path)
	//r.Use(static.Serve("/dyn-assets/", static.LocalFile(path, true)))
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
