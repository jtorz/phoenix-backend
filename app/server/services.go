package server

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/services/authorization"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndhttp"
	"github.com/jtorz/phoenix-backend/app/services/mail"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
	"github.com/spf13/viper"
)

func (server *Server) configureServices() {
	if server.Config.JWTKey == "" {
		panic("empty JWT key")
	}
	jwtSvc := authorization.JWTSvc(server.Config.JWTKey)

	gin.SetMode(server.Config.AppMode)

	r := gin.New()
	server.configureMiddlewares(r, jwtSvc)

	apiGroup := r.Group("/api")
	var mailSenderSvc baseservice.MailSenderSvc

	{
		mailConfig := mail.ServiceConfig{}
		mailConfig.SetDefaults(viper.SetDefault)
		server.registerEnvs(mailConfig)
		viper.Unmarshal(&mailConfig)
		mailSenderSvc = mail.NewService(server.MainDB, mailConfig, server.Config.Domain)

		/* route := "/email"
		adminGroup := apiGroup.Group("/admin")
		publicGroup := apiGroup.Group("/public")
		fndSVC := fndhttp.NewService(server.MainDB, jwtSvc)
		fndSVC.API(apiGroup.Group(route))
		fndSVC.APIAdmin(adminGroup.Group(route))
		fndSVC.APIPublic(publicGroup.Group(route)) */
	}

	{
		route := "/foundation"
		fndSVC := fndhttp.NewService(server.MainDB, jwtSvc, mailSenderSvc)

		adminGroup := apiGroup.Group("/admin")
		publicGroup := apiGroup.Group("/public")
		fndSVC.API(apiGroup.Group(route))
		fndSVC.APIAdmin(adminGroup.Group(route))
		fndSVC.APIPublic(publicGroup.Group(route))
	}

	h := http.TimeoutHandler(r, time.Duration(server.Config.RequestTimeout)*time.Second, `"request timeout"`)

	server.HTTPServer = &http.Server{
		Addr:           fmt.Sprintf(":%d", server.Config.Port),
		Handler:        h,
		ReadTimeout:    30 * time.Second,
		WriteTimeout:   30 * time.Second,
		MaxHeaderBytes: 1 << 20, //1 MB
	}
}
