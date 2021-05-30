package mailbiz

import (
	"context"
	"database/sql"

	"github.com/jtorz/phoenix-backend/app/services/mail/maildao"
	"github.com/jtorz/phoenix-backend/app/services/mail/mailmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// BizSender business component.
type BizSender struct {
	Exe base.Executor
	dao *maildao.DaoSender
}

// NewBizSender creates the business component.
func NewBizSender() BizSender {
	return BizSender{
		dao: &maildao.DaoSender{},
	}
}

// GetByID retrives the record information using its ID.
func (biz *BizSender) GetByID(ctx context.Context, exe base.Executor,
	id string,
) (*mailmodel.Sender, error) {
	rec, err := biz.dao.GetByID(ctx, exe, id)
	if err != nil {
		return nil, err
	}
	biz.setRecordActions(ctx, rec)
	return rec, nil
}

// List returns the list of records.
func (biz *BizSender) List(ctx context.Context, exe base.Executor,
	onlyActive bool,
) (mailmodel.Senders, error) {
	recs, err := biz.dao.List(ctx, exe, onlyActive)
	if err != nil {
		return nil, err
	}
	biz.setRecordActionsSenders(ctx, recs)
	return recs, nil
}

// New creates a new record.
func (biz *BizSender) New(ctx context.Context, tx *sql.Tx,
	rec *mailmodel.Sender,
) error {
	rec.Status = base.StatusCaptured
	if err := biz.dao.New(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Edit edits the record.
func (biz *BizSender) Edit(ctx context.Context, tx *sql.Tx,
	rec *mailmodel.Sender,
) error {
	return biz.dao.Edit(ctx, tx, rec)
}

// SetStatus updates the logical status of the record.
func (biz *BizSender) SetStatus(ctx context.Context, tx *sql.Tx,
	rec *mailmodel.Sender,
) error {
	if err := biz.dao.SetStatus(ctx, tx, rec); err != nil {
		return err
	}
	biz.setRecordActions(ctx, rec)
	return nil
}

// Delete performs a physical delete of the record.
func (biz *BizSender) Delete(ctx context.Context, tx *sql.Tx,
	rec *mailmodel.Sender,
) error {
	if err := biz.dao.Delete(ctx, tx, rec); err != nil {
		return err
	}
	return nil
}

// setRecordActionsSenders sets the record actiosn to every element in the Senders slice.
func (biz *BizSender) setRecordActionsSenders(ctx context.Context,
	recs mailmodel.Senders,
) {
	for i := range recs {
		biz.setRecordActions(ctx, &recs[i])
	}
}

// setRecordActions sets the record action sto Sender record.
func (biz *BizSender) setRecordActions(ctx context.Context,
	rec *mailmodel.Sender,
) {
	rec.RecordActions = base.NewValidateInvalidate(rec.Status)
}
