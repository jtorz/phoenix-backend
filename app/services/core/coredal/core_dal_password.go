package coredal

import (
	"context"
	"database/sql"
	"encoding/json"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/jtorz/phoenix-backend/app/services/core/coremodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"

	//lint:ignore ST1001 dot import allowed only in dal packages for
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DalPassword Data Access structure.
type DalPassword struct{}

/// New saves the password record in the database.
func (dal *DalPassword) New(ctx context.Context, tx *sql.Tx,
	usuarioID string, p *coremodel.Password,
) error {
	jsonData, err := json.Marshal(p.Data)
	if err != nil {
		return err
	}
	now := time.Now()
	ins := NewInsert(T.CorePassword).Rows(goqu.Record{
		CorePassword.PasData:      jsonData,
		CorePassword.PasUserID:    usuarioID,
		CorePassword.PasType:      p.Type,
		CorePassword.PasStatus:    p.Status,
		CorePassword.PasUpdatedAt: now,
	})
	row, err := DoInsertReturning(ctx, tx, ins, CorePassword.PasID)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	if err := row.Scan(&p.ID); err != nil {
		return err
	}
	p.UpdatedAt = now
	return nil
}

// InvalidateForUser invalidates all the passwords of a user.
func (dal DalPassword) InvalidateForUser(ctx context.Context, tx *sql.Tx,
	userID string,
) error {
	query := NewUpdate(T.CorePassword).
		Set(goqu.Record{
			CorePassword.PasInvalidationDate: goqu.L("CURRENT_TIMESTAMP"),
			CorePassword.PasStatus:           base.StatusInactive,
		}).
		Where(
			goqu.C(CorePassword.PasUserID).Eq(userID),
			goqu.C(CorePassword.PasStatus).Eq(base.StatusActive),
		)
	_, err := DoUpdate(ctx, tx, query)
	DebugErr(ctx, err)
	return err
}
