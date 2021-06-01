package maildao

import (
	"context"
	"database/sql"
	"time"

	"github.com/doug-martin/goqu/v9"
	"github.com/doug-martin/goqu/v9/exp"
	"github.com/jtorz/phoenix-backend/app/services/mail/mailmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"

	//lint:ignore ST1001 dot import allowed only in dao packages for
	. "github.com/jtorz/phoenix-backend/app/shared/lex"
)

// DaoSender Data Access structure.
type DaoSender struct{}

// GetByID retrives the record information using its ID.
func (dao *DaoSender) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*mailmodel.Sender, error) {
	rec := mailmodel.Sender{}
	query := NewSelect(
		MailSender.SenName,
		MailSender.SenDescription,
		MailSender.SenHost,
		MailSender.SenPort,
		MailSender.SenUser,
		MailSender.SenPassword,
		MailSender.SenFrom,
		MailSender.SenCreatedAt,
		MailSender.SenUpdatedAt,
		MailSender.SenStatus,
	).
		From(goqu.T(T.MailSender)).
		Where(
			goqu.C(MailSender.SenID).Eq(id),
		)

	row, err := QueryRowContext(ctx, exe, query)
	if err != nil {
		DebugErr(ctx, err)
		return nil, err
	}
	err = row.Scan(
		&rec.Name,
		&rec.Description,
		&rec.Host,
		&rec.Port,
		&rec.User,
		&rec.Password,
		&rec.From,
		&rec.CreatedAt,
		&rec.UpdatedAt,
		&rec.Status,
	)
	if err != nil {
		DebugErr(ctx, err)
		return nil, WrapNotFound(ctx, T.MailSender, err)
	}

	rec.ID = id
	return &rec, nil
}

// List returns the list of records.
func (dao *DaoSender) List(ctx context.Context, exe base.Executor,
	onlyActive bool,
) (mailmodel.Senders, error) {
	res := make(mailmodel.Senders, 0)
	filter := exp.NewExpressionList(exp.AndType)
	if onlyActive {
		filter = filter.Append(goqu.C(MailSender.SenStatus).Eq(base.StatusActive))
	}
	query := NewSelect(
		MailSender.SenID,
		MailSender.SenName,
		MailSender.SenDescription,
		MailSender.SenHost,
		MailSender.SenPort,
		MailSender.SenUser,
		MailSender.SenPassword,
		MailSender.SenFrom,
		MailSender.SenCreatedAt,
		MailSender.SenUpdatedAt,
		MailSender.SenStatus,
	).From(T.MailSender).
		Where(filter)

	rows, err := QueryContext(ctx, exe, query)
	if err != nil {
		if err == sql.ErrNoRows {
			return res, nil
		}
		DebugErr(ctx, err)
		return nil, err
	}
	defer rows.Close()
	for rows.Next() {
		rec := mailmodel.Sender{}
		err := rows.Scan(
			&rec.ID,
			&rec.Name,
			&rec.Description,
			&rec.Host,
			&rec.Port,
			&rec.User,
			&rec.Password,
			&rec.From,
			&rec.CreatedAt,
			&rec.UpdatedAt,
			&rec.Status,
		)
		if err != nil {
			DebugErr(ctx, err)
			return nil, err
		}
		res = append(res, rec)
	}
	return res, nil
}

// New creates a new record.
func (dao *DaoSender) New(ctx context.Context, tx *sql.Tx,
	rec *mailmodel.Sender,
) error {
	now := time.Now()
	record := goqu.Record{
		MailSender.SenID:          rec.ID,
		MailSender.SenName:        rec.Name,
		MailSender.SenDescription: rec.Description,
		MailSender.SenHost:        rec.Host,
		MailSender.SenPort:        rec.Port,
		MailSender.SenUser:        rec.User,
		MailSender.SenPassword:    rec.Password,
		MailSender.SenFrom:        rec.From,
		MailSender.SenCreatedAt:   rec.CreatedAt,
		MailSender.SenUpdatedAt:   now,
		MailSender.SenStatus:      rec.Status,
	}
	ins := NewInsert(T.MailSender).Rows(record)
	_, err := DoInsert(ctx, tx, ins)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return nil
}

// Edit edits the record.
func (dao *DaoSender) Edit(ctx context.Context, tx *sql.Tx,
	rec *mailmodel.Sender,
) error {
	now := time.Now()
	record := goqu.Record{
		MailSender.SenName:        rec.Name,
		MailSender.SenDescription: rec.Description,
		MailSender.SenHost:        rec.Host,
		MailSender.SenPort:        rec.Port,
		MailSender.SenUser:        rec.User,
		MailSender.SenPassword:    rec.Password,
		MailSender.SenFrom:        rec.From,
		MailSender.SenUpdatedAt:   now,
		MailSender.SenStatus:      rec.Status,
	}

	query := NewUpdate(T.MailSender).
		Set(record).
		Where(
			goqu.C(MailSender.SenID).Eq(rec.ID),

			goqu.C(MailSender.SenUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.MailSender, res)
}

// SetStatus updates the logical status of the record.
func (dao *DaoSender) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *mailmodel.Sender,
) error {
	now := time.Now()
	query := NewUpdate(T.MailSender).
		Set(goqu.Record{
			MailSender.SenUpdatedAt: now,
			MailSender.SenStatus:    rec.Status,
		}).
		Where(
			goqu.C(MailSender.SenID).Eq(rec.ID),
			goqu.C(MailSender.SenUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoUpdate(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	rec.UpdatedAt = now
	return CheckOneRowUpdated(ctx, T.MailSender, res)
}

// Delete performs a physical delete of the record.
func (dao *DaoSender) Delete(ctx context.Context, tx *sql.Tx,
	rec *mailmodel.Sender,
) error {
	query := NewDelete(T.MailSender).
		Where(
			goqu.C(MailSender.SenID).Eq(rec.ID),
			goqu.C(MailSender.SenUpdatedAt).Eq(rec.UpdatedAt),
		)
	res, err := DoDelete(ctx, tx, query)
	if err != nil {
		DebugErr(ctx, err)
		return err
	}
	return CheckOneRowUpdated(ctx, T.MailSender, res)
}
