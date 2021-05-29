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
	httpAccount := newHttpAccount(s.DB)
	{
		apiGroup.GET("/account/session", httpAccount.GetSessionData().Func())
	}

	httpNavElement := newHttpNavElement(s.DB)
	{
		apiGroup.POST("/navigator/upsert", httpNavElement.UpsertAll().Func())
		apiGroup.GET("/navigator/elements/element/:id", httpNavElement.GetByID().Func())
		apiGroup.GET("/navigator/elements", httpNavElement.ListAll().Func())
		apiGroup.GET("/navigator/elements/role/:roleID", httpNavElement.ListAll().Func())
		apiGroup.POST("/navigator/elements/element", httpNavElement.New().Func())
		apiGroup.PUT("/navigator/elements/element", httpNavElement.Edit().Func())
		apiGroup.PUT("/navigator/elements/element/validate", httpNavElement.SetStatus(base.StatusActive).Func())
		apiGroup.PUT("/navigator/elements/element/invalidate", httpNavElement.SetStatus(base.StatusInactive).Func())
		apiGroup.PUT("/navigator/elements/element/soft-delete", httpNavElement.SetStatus(base.StatusDroppped).Func())
		apiGroup.PUT("/navigator/elements/element/hard-delete", httpNavElement.Delete().Func())
		apiGroup.PUT("/navigator/elements/element/associate-role", httpNavElement.AssociateRole().Func())
		apiGroup.PUT("/navigator/elements/element/dissociate-role", httpNavElement.DissociateRole().Func())
	}

	httpModule := newHttpModule(s.DB)
	{
		apiGroup.GET("/modules/module/:moduleID", httpModule.GetByID().Func())
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

		httpAction := newHttpAction(s.DB)
		{
			apiGroup.GET("/modules/module/:moduleID/actions/action/:actionID", httpAction.GetByID().Func())
			apiGroup.GET("/modules/module/:moduleID/actions", httpAction.ListAll().Func())
			apiGroup.POST("/modules/module/:moduleID/actions", httpAction.ListAll().Func())
			apiGroup.GET("/modules/module/:moduleID/actions/active-records", httpAction.ListActive().Func())
			apiGroup.POST("/modules/module/:moduleID/actions/active-records", httpAction.ListActive().Func())
			apiGroup.POST("/modules/actions/action", httpAction.New().Func())
			apiGroup.PUT("/modules/actions/action", httpAction.Edit().Func())
			apiGroup.PUT("/modules/actions/action/validate", httpAction.SetStatus(base.StatusActive).Func())
			apiGroup.PUT("/modules/actions/action/invalidate", httpAction.SetStatus(base.StatusInactive).Func())
			apiGroup.PUT("/modules/actions/action/soft-delete", httpAction.SetStatus(base.StatusDroppped).Func())
			apiGroup.PUT("/modules/actions/action/hard-delete", httpAction.Delete().Func())
		}
	}

	httpRole := newHttpRole(s.DB)
	{
		apiGroup.GET("/roles/role/:id", httpRole.GetByID().Func())
		apiGroup.GET("/roles", httpRole.ListAll().Func())
		apiGroup.GET("/roles/active-records", httpRole.ListActive().Func())
		apiGroup.POST("/roles/role", httpRole.New().Func())
		apiGroup.PUT("/roles/role", httpRole.Edit().Func())
		apiGroup.PUT("/roles/role/validate", httpRole.SetStatus(base.StatusActive).Func())
		apiGroup.PUT("/roles/role/invalidate", httpRole.SetStatus(base.StatusInactive).Func())
		apiGroup.PUT("/roles/role/soft-delete", httpRole.SetStatus(base.StatusDroppped).Func())
		apiGroup.PUT("/roles/role/hard-delete", httpRole.Delete().Func())
	}
}
