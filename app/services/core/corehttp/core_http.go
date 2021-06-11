package corehttp

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gomodule/redigo/redis"
	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

type Service struct {
	DB      *sql.DB
	Redis   *redis.Pool
	JwtSvc  baseservice.JWTGeneratorSvc
	MailSvc baseservice.MailSenderSvc
}

func NewHttpService(db *sql.DB, jwtSvc baseservice.JWTGeneratorSvc, mailSvc baseservice.MailSenderSvc, redis *redis.Pool) Service {
	return Service{
		DB:      db,
		Redis:   redis,
		JwtSvc:  jwtSvc,
		MailSvc: mailSvc,
	}
}

// APIPublic registers the http public routes of the component.
// The Public Routes are available for any user no matter if they are authenticated.
//
// current path: /api/public/core
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
// current path: /api/admin/core
func (s Service) APIAdmin(apiGroup *gin.RouterGroup) {
	apiGroup.POST("/cache/clear", httphandler.HandlerFunc(s.clearCache).Func())
}

func (s Service) clearCache(c *httphandler.Context) {
	conn := s.Redis.Get()
	defer conn.Close()

	_, err := conn.Do("FLUSHALL")
	if c.HandleError(err) {
		return
	}
	c.Status(http.StatusOK)
}

// APIAdmin registers the http routes  available to all the authenticated users.
// The Admin routes are available only for the administrators of the system.
//
// current path: /api/auth-all/core
func (s Service) APIAuthAll(apiGroup *gin.RouterGroup) {
	httpAccount := newHttpAccount(s.DB)
	{
		apiGroup.GET("/account/session", httpAccount.GetSessionData().Func())
	}
}

// API registers the http general routes of the component.
// The General routes are available only for the authenticated users.
//
// current path: /api/core
func (s Service) API(apiGroup *gin.RouterGroup) {

	httpNavElement := newHttpNavElement(s.DB)
	{
		apiGroup.POST("/navigator/upsert", httpNavElement.UpsertOrDeleteAll().Func())
		apiGroup.GET("/navigator/elements/element/:id", httpNavElement.GetByID().Func())
		apiGroup.GET("/navigator/elements", httpNavElement.ListAll().Func())
		apiGroup.GET("/navigator/elements/role/:roleID", httpNavElement.ListAll().Func())
		apiGroup.PUT("/navigator/elements/element/activate", httpNavElement.SetStatus(base.StatusActive).Func())
		apiGroup.PUT("/navigator/elements/element/inactivate", httpNavElement.SetStatus(base.StatusInactive).Func())
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
		apiGroup.PUT("/modules/module/activate", httpModule.SetStatus(base.StatusActive).Func())
		apiGroup.PUT("/modules/module/inactivate", httpModule.SetStatus(base.StatusInactive).Func())
		//apiGroup.PUT("/modules/module/soft-delete", httpModule.SetStatus(base.StatusDroppped).Func())
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
			apiGroup.PUT("/modules/actions/action/activate", httpAction.SetStatus(base.StatusActive).Func())
			apiGroup.PUT("/modules/actions/action/inactivate", httpAction.SetStatus(base.StatusInactive).Func())
			//apiGroup.PUT("/modules/actions/action/soft-delete", httpAction.SetStatus(base.StatusDroppped).Func())
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
		apiGroup.PUT("/roles/role/activate", httpRole.SetStatus(base.StatusActive).Func())
		apiGroup.PUT("/roles/role/inactivate", httpRole.SetStatus(base.StatusInactive).Func())
		apiGroup.PUT("/roles/role/soft-delete", httpRole.SetStatus(base.StatusDroppped).Func())
		apiGroup.PUT("/roles/role/hard-delete", httpRole.Delete().Func())
	}
}
