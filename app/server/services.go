package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/services/authorization"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndhttp"
)

type Service struct {
	Name       string
	Configurer interface {
		// API registers the http general routes of the component.
		// The General routes are available only for the authenticated users.
		API(*gin.RouterGroup)
		// APIAdmin registers the http admin routes of the component.
		// The Admin routes are available only for the administrators of the system.
		APIAdmin(*gin.RouterGroup)
		// APIPublic registers the http public routes of the component.
		// The Public Routes are available for any user no matter if they are authenticated.
		APIPublic(*gin.RouterGroup)
	}
}

func (server *Server) regiterServices(jwtSvc authorization.JWTService) {
	server.Services = map[string]Service{
		"/foundation": {
			Name:       "Foundation",
			Configurer: fndhttp.NewService(server.MainDB, jwtSvc),
		},
	}
}

func (server *Server) configureHttpServer() {
	if server.Config.JWTKey == "" {
		panic("empty JWT key")
	}
	jwtSvc := authorization.JWTService(server.Config.JWTKey)

	server.regiterServices(jwtSvc)

	r := gin.New()

	gin.SetMode(server.Config.AppMode)
	server.api(r.Group("/api"))
	server.configureMiddlewares(r, jwtSvc)
	h := http.TimeoutHandler(r, time.Duration(server.Config.RequestTimeout)*time.Second, `"request timeout"`)
	server.HTTPServer = &http.Server{
		Addr:           fmt.Sprintf(":%d", server.Config.Port),
		Handler:        h,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, //1 MB
	}
}

func (server *Server) api(apiGroup *gin.RouterGroup) {
	// public routes.
	adminGroup := apiGroup.Group("/admin")
	publicGroup := apiGroup.Group("/public")
	for route, service := range server.Services {
		service.Configurer.API(apiGroup.Group(route))
		service.Configurer.APIAdmin(adminGroup.Group(route))
		service.Configurer.APIPublic(publicGroup.Group(route))
	}
}
