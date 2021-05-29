package ctxinfo

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/shared/baseservice"
)

const modeKey = config.EnvPrefix + "_mode_"

// LoggingLevel returns the level of logging int he context.
func LoggingLevel(ctx context.Context) config.LogginLvl {
	lvl, _ := ctx.Value(modeKey).(config.LogginLvl)
	return lvl
}

// LogginAllowed verifies if the logging is allowed in the context.
// Only the loggin is allowed when the required log is bigger (or equal) than the context.LogginLvl.
func LogginAllowed(ctx context.Context, lvl config.LogginLvl) bool {
	ctxLvl := LoggingLevel(ctx)
	return lvl >= ctxLvl
}

// SetLoggingLevel sets level of logging to the gin.Context.
func SetLoggingLevel(c *gin.Context, lvl config.LogginLvl) {
	c.Set(modeKey, lvl)
}

const agentKey = config.EnvPrefix + "_agent_"

// SetAgent sets the agent to the gin.Context.
func SetAgent(c *gin.Context, agent *baseservice.Agent) {
	c.Set(agentKey, agent)
}

// GetAgent returs the agent from the context.
func GetAgent(c context.Context) *baseservice.Agent {
	a, ok := c.Value(agentKey).(*baseservice.Agent)
	if !ok {
		return nil
	}
	return a
}
