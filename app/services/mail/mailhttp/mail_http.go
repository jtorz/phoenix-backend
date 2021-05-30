package mailhttp

import (
	"database/sql"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

type Service struct {
	DB      *sql.DB
	JwtSvc  baseservice.JWTGeneratorSvc
	MailSvc baseservice.MailSenderSvc
}

func NewHttpService(db *sql.DB) Service {
	return Service{
		DB: db,
	}
}

// APIPublic registers the http public routes of the component.
// The Public Routes are available for any user no matter if they are authenticated.
//
// current path: /api/public/mail
func (s Service) APIPublic(apiGroup *gin.RouterGroup) {
}

// APIAdmin registers the http admin routes of the component.
// The Admin routes are available only for the administrators of the system.
//
// current path: /api/admin/mail
func (s Service) APIAdmin(apiGroup *gin.RouterGroup) {
}

// API registers the http general routes of the component.
// The General routes are available only for the authenticated users.
//
// current path: /api/mail
func (s Service) API(apiGroup *gin.RouterGroup) {

	httpSender := newHttpSender(s.DB)
	{
		apiGroup.GET("/senders/sender/:id", httpSender.GetByID().Func())
		apiGroup.GET("/senders", httpSender.ListAll().Func())
		apiGroup.GET("/senders/active-records", httpSender.ListActive().Func())
		apiGroup.POST("/senders/sender", httpSender.New().Func())
		apiGroup.PUT("/senders/sender", httpSender.Edit().Func())
		apiGroup.PUT("/senders/sender/validate", httpSender.SetStatus(base.StatusActive).Func())
		apiGroup.PUT("/senders/sender/invalidate", httpSender.SetStatus(base.StatusInactive).Func())
		apiGroup.PUT("/senders/sender/soft-delete", httpSender.SetStatus(base.StatusDroppped).Func())
		apiGroup.PUT("/senders/sender/hard-delete", httpSender.Delete().Func())
	}
}
