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
	navigator := newHttpNavigator(s.DB)
	{
		apiGroup.GET("/navigators/navigator/:id", navigator.GetByID().Func())
		apiGroup.GET("/navigators", navigator.List(false).Func())
		apiGroup.GET("/navigators/active-records", navigator.List(true).Func())
		apiGroup.POST("/navigators/active-records", navigator.List(true).Func())
		apiGroup.POST("/navigators/navigator", navigator.New().Func())
		apiGroup.PUT("/navigators/navigator", navigator.Edit().Func())
		apiGroup.PUT("/navigators/navigator/validate", navigator.SetStatus(base.StatusActive).Func())
		apiGroup.PUT("/navigators/navigator/invalidate", navigator.SetStatus(base.StatusInactive).Func())
		apiGroup.PUT("/navigators/navigator/soft-delete", navigator.SetStatus(base.StatusDroppped).Func())
		apiGroup.PUT("/navigators/navigator/hard-delete", navigator.Delete().Func())
	}
}
