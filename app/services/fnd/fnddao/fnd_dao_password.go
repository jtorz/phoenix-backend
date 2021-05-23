package fnddao

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"

	//lint:ignore ST1001 dot import allowed only in dao packages for lex.
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoPassword Data Access structure.
type DaoPassword struct {
	Exe base.Executor
}

/// New saves the password record in the database.
func (dao *DaoPassword) New(ctx context.Context, tx *sql.Tx,
	usuarioID string, p *fndmodel.Password,
) error {
	jsonData, err := json.Marshal(p.Data)
	if err != nil {
		return err
	}
	now := time.Now()
	ins := NewInsert(T.FndPassword).Rows(goqu.Record{
		FndPassword.PasData:      jsonData,
		FndPassword.PasUserID:    usuarioID,
		FndPassword.PasType:      p.Type,
		FndPassword.PasStatus:    p.Status,
		FndPassword.PasUpdatedAt: now,
	})
	row, err := DoInsertReturning(ctx, dao.Exe, ins, FndPassword.PasID)
	if err != nil {
		return WrapErr(ctx, err)
	}
	if err := row.Scan(&p.ID); err != nil {
		return err
	}
	p.UpdatedAt = now
	return nil
}
