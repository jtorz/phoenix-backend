package fnddao

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"

	//lint:ignore ST1001 dot import allowed only in dao packages for
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoPassword Data Access structure.
type DaoPassword struct{}

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
	row, err := DoInsertReturning(ctx, tx, ins, FndPassword.PasID)
	if err != nil {
		return DebugErr(ctx, err)
	}
	if err := row.Scan(&p.ID); err != nil {
		return err
	}
	p.UpdatedAt = now
	return nil
}

// InvalidateForUser invalidates all the passwords of a user.
func (dao DaoPassword) InvalidateForUser(ctx context.Context, tx *sql.Tx,
	userID string,
) error {
	query := NewUpdate(T.FndPassword).
		Set(goqu.Record{
			FndPassword.PasInvalidationDate: goqu.L("CURRENT_TIMESTAMP"),
			FndPassword.PasStatus:           base.StatusInactive,
		}).
		Where(
			goqu.C(FndPassword.PasUserID).Eq(userID),
			goqu.C(FndPassword.PasStatus).Eq(base.StatusActive),
		)
	_, err := DoUpdate(ctx, tx, query)
	return DebugErr(ctx, err)
}
