package fndhttp

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

type Service struct {
	DB      *sql.DB
	JwtSvc  baseservice.JWTGeneratorSvc
	MailSvc baseservice.MailSenderSvc
}

func NewService(db *sql.DB, jwtSvc baseservice.JWTGeneratorSvc, mailSvc baseservice.MailSenderSvc) Service {
	return Service{
		DB:      db,
		JwtSvc:  jwtSvc,
		MailSvc: mailSvc,
	}
}

// APIPublic registers the http public routes of the component.
// The Public Routes are available for any user no matter if they are authenticated.
//
// current path: /api/public/foundation
func (s Service) APIPublic(apiGroup *gin.RouterGroup) {
	httpPublic := newHttpPublic(s.DB, s.JwtSvc, s.MailSvc)
	{
		//g.POST("/account/logout", httpPublic.Logout().Func())
		apiGroup.POST("/account/login", httphandler.Secret(httpPublic.Login()))
		apiGroup.POST("/account/signup", httphandler.Secret(httpPublic.Signup()))
		apiGroup.POST("/account/restore/request", httpPublic.RequestRestoreAccount().Func())
		apiGroup.POST("/account/restore", httpPublic.RestoreAccount().Func())
		//apiGroup.GET("/account/session", httpPublic.GetSession().Func())
	}
}

// APIAdmin registers the http admin routes of the component.
// The Admin routes are available only for the administrators of the system.
//
// current path: /api/admin/foundation
func (s Service) APIAdmin(apiGroup *gin.RouterGroup) {

}

// API registers the http general routes of the component.
// The General routes are available only for the authenticated users.
//
// current path: /api/foundation
func (s Service) API(apiGroup *gin.RouterGroup) {

	httpNavigator := newHttpNavigator(s.DB)
	{
		apiGroup.GET("/navigators/navigator/:id", httpNavigator.GetByID().Func())
		apiGroup.GET("/navigators", httpNavigator.ListAll().Func())
		apiGroup.POST("/navigators", httpNavigator.ListAll().Func())
		apiGroup.GET("/navigators/active-records", httpNavigator.ListActive().Func())
		apiGroup.POST("/navigators/active-records", httpNavigator.ListActive().Func())
		apiGroup.POST("/navigators/navigator", httpNavigator.New().Func())
		apiGroup.PUT("/navigators/navigator", httpNavigator.Edit().Func())
		apiGroup.PUT("/navigators/navigator/validate", httpNavigator.SetStatus(base.StatusActive).Func())
		apiGroup.PUT("/navigators/navigator/invalidate", httpNavigator.SetStatus(base.StatusInactive).Func())
		apiGroup.PUT("/navigators/navigator/soft-delete", httpNavigator.SetStatus(base.StatusDroppped).Func())
		apiGroup.PUT("/navigators/navigator/hard-delete", httpNavigator.Delete().Func())
	}

	httpModule := newHttpModule(s.DB)
	{
		apiGroup.GET("/modules/module/:id", httpModule.GetByID().Func())
		apiGroup.GET("/modules", httpModule.ListAll().Func())
		apiGroup.POST("/modules", httpModule.ListAll().Func())
		apiGroup.GET("/modules/active-records", httpModule.ListActive().Func())
		apiGroup.POST("/modules/active-records", httpModule.ListActive().Func())
		apiGroup.POST("/modules/module", httpModule.New().Func())
		apiGroup.PUT("/modules/module", httpModule.Edit().Func())
		apiGroup.PUT("/modules/module/validate", httpModule.SetStatus(base.StatusActive).Func())
		apiGroup.PUT("/modules/module/invalidate", httpModule.SetStatus(base.StatusInactive).Func())
		apiGroup.PUT("/modules/module/soft-delete", httpModule.SetStatus(base.StatusDroppped).Func())
		apiGroup.PUT("/modules/module/hard-delete", httpModule.Delete().Func())
	}
}
