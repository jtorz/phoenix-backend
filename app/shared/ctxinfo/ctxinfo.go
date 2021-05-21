package ctxinfo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/config"
)

const modeKey = config.EnvPrefix + "_mode_"

// LoggingLevel returns the level of logging int he context.
func LoggingLevel(ctx context.Context) config.LogginLvl {
	lvl, _ := ctx.Value(modeKey).(config.LogginLvl)
	if lvl == "" {
		return config.LogDebug
	}
	return lvl
}

// SetLoggingLevel sets level of logging to the gin.Context.
func SetLoggingLevel(c *gin.Context, lvl config.LogginLvl) {
	c.Set(modeKey, lvl)
}

// Agent has the information of the user that is executing an operation,
// and an AuthService to retrive their privileges.
type Agent struct {
	UserID string
	AuthService
}

// AuthService is used to retrieve the privileges of an agent.
type AuthService interface {
	GetPrivilegeByPriority(...string) (string, error)
	HasPrivilege(string) (bool, error)
	IsAdmin() (bool, error)
}

const agentKey = config.EnvPrefix + "_agent_"

// SetAgent sets the agent to the gin.Context.
func SetAgent(c *gin.Context, userID string, authSvc AuthService) {
	c.Set(agentKey, &Agent{
		UserID:      userID,
		AuthService: authSvc,
	})
}

// GetAgent returs the agent from the context.
func GetAgent(c context.Context) *Agent {
	a, ok := c.Value(agentKey).(*Agent)
	if !ok {
		return nil
	}
	return a
}
