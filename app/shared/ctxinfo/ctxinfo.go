package ctxinfo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/config"
)

const modeKey = config.EnvPrefix + "_mode_"

func PrintLog(ctx context.Context) bool {
	mode, _ := ctx.Value(modeKey).(config.Mode)
	return config.IsModeDebug(mode)
}

func SetPrintLog(c *gin.Context, mode config.Mode) {
	c.Set(modeKey, mode)
}

type Agent struct {
	UserID string
	AuthService
}
type AuthService interface {
	GetPrivilegeByPriority(...string) (string, error)
	HasPrivilege(string) (bool, error)
	IsAdmin() (bool, error)
}

const agentKey = config.EnvPrefix + "_agent_"

func SetAgent(c *gin.Context, userID string, authSvc AuthService) {
	c.Set(agentKey, &Agent{
		UserID:      userID,
		AuthService: authSvc,
	})
}
