package corehttp

import (
	"database/sql"
	"net/http"

	"github.com/jtorz/phoenix-backend/app/httphandler"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
)

// httpAccount http handler component.
type httpAccount struct {
	DB *sql.DB
}

func newHttpAccount(db *sql.DB) httpAccount {
	return httpAccount{
		DB: db,
	}
}

// GetSessionData return the user information of the agent.
func (handler httpAccount) GetSessionData() httphandler.HandlerFunc {
	return func(c *httphandler.Context) {
		a := ctxinfo.GetAgent(c)
		agentInfo, err := a.GetInfo(c)
		if c.HandleError(err) {
			return
		}
		c.JSON(http.StatusOK, agentInfo)
	}
}
