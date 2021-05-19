package fndbiz

import (
	"context"
	"testing"

	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/stretchr/testify/assert"
)

func TestLogin(t *testing.T) {
	bizUser := newBizUserTest()
	var err error
	var u *fndmodel.User
	u, err = bizUser.Login(context.Background(), "john.doe", "1")
	assert.Nil(t, err, "error must be nil")
	assert.NotNil(t, u, " user must not be nil when login without error")

	u, err = bizUser.Login(context.Background(), "john.doe", "2")
	assert.ErrorIsf(t, err, baseerrors.ErrAuth, "when password mismatch error must be %s", baseerrors.ErrAuth)
	assert.Nil(t, u, "user must be nil when login with error")

	u, err = bizUser.Login(context.Background(), "jane.doe", "1")
	assert.ErrorIsf(t, err, baseerrors.ErrActionNotAllowedStatus, "when user is not active error must be %s", baseerrors.ErrActionNotAllowedStatus)
	assert.Nil(t, u, "user must be nil when login with error")

	u, err = bizUser.Login(context.Background(), "unexistentuser", "1")
	assert.ErrorIsf(t, err, baseerrors.ErrAuth, "when user not exist error must be %s", baseerrors.ErrAuth)
	assert.Nil(t, u, "user must be nil when login with error")
}
