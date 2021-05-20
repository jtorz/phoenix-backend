package fndbiz

import (
	"context"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// BizAccountAccess business component.
type BizAccountAccess struct {
	Exe base.Executor
	dao DaoBizAccountAccess
}

// NewBizAccountAccess creates business component.
func NewBizAccountAccess(exe base.Executor) BizAccountAccess {
	return BizAccountAccess{
		Exe: exe,
		dao: &fnddao.DaoAccountAccess{Exe: exe},
	}
}

type DaoBizAccountAccess interface {
	Insert(context.Context, *fndmodel.AccountAccess) error
	UseAccountAccess(context.Context, fndmodel.AccountAccessType, string) (userID string, _ error)
	GetAccessByUserID(_ context.Context, _ fndmodel.AccountAccessType, userID string) (*fndmodel.AccountAccess, error)
}

// GetAccessByUserID returns the user's active access.
func (biz *BizAccountAccess) GetAccessByUserID(ctx context.Context,
	accType fndmodel.AccountAccessType, userID string,
) (*fndmodel.AccountAccess, error) {
	return biz.dao.GetAccessByUserID(ctx, accType, userID)
}

// UseAccountAccess returns the user ID for the account access and inactivates the access.
func (biz *BizAccountAccess) UseAccountAccess(ctx context.Context,
	accType fndmodel.AccountAccessType, key string,
) (string, error) {
	return biz.dao.UseAccountAccess(ctx, accType, key)
}

// NewAccountAccess create the new account access,
func (biz *BizAccountAccess) NewAccountAccess(ctx context.Context,
	u fndmodel.User,
) (*fndmodel.AccountAccess, error) {

	ac, err := fndmodel.NewAccountAccess(u, fndmodel.AccAccRestoreAccount)
	if err != nil {
		return nil, err
	}

	if err = biz.dao.Insert(ctx, ac); err != nil {
		return nil, err
	}
	return ac, nil
}
