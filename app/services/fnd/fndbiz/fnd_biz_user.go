package fndbiz

import (
	"context"
	"fmt"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
)

// BizUser business component.
type BizUser struct {
	dao DaoUser
}

func NewBizUser() BizUser {
	return BizUser{dao: fnddao.DaoUser{}}
}

type DaoUser interface {
	Login(ctx context.Context, exe base.Executor, user string) (*fndmodel.User, error)
}

// Login retrieves the necessary data to log in a user given its email/username.
func (biz *BizUser) Login(ctx context.Context, exe base.Executor,
	user, pass string,
) (*fndmodel.User, error) {

	u, err := biz.dao.Login(ctx, exe, user)
	if err != nil {
		return nil, err
	}
	if u.Status == base.StatusCaptured || u.Status == base.StatusDroppped || u.Status == base.StatusInactive {
		return nil, fmt.Errorf("Can't login user in status %s: %w", u.Status, baseerrors.ErrActionNotAllowedStatus)
	}
	if err = u.Password.ComparePassword(pass); err != nil {
		return nil, err
	}
	return u, nil
}
