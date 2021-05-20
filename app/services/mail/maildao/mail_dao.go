package maildao

import (
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/daohelper"
)

// DaoMail Data Access structure.
type DaoMail struct {
	Exe base.Executor
	h   daohelper.QueryHelper
}
