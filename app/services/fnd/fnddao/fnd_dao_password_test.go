package fnddao_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fnddao"
	"github.com/jtorz/phoenix-backend/app/services/fnd/fndmodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/stretchr/testify/assert"
)

var daoPassword = fnddao.DaoPassword{}

func TestPasswordNew(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("*Begin.Tx ERROR* %s", err)
	}
	any := sqlmock.AnyArg()
	mock.ExpectQuery(`INSERT INTO "fnd_password"`).WithArgs([]byte("{}"), 2, fndmodel.PassTypeScrypt2017, any, "848577e0-9ec9-5d25-bcb8-8ed913100609").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// now we execute our method
	pass := fndmodel.Password{
		Type:   fndmodel.PassTypeScrypt2017,
		Data:   base.JSONObject{},
		Status: 2,
	}
	if err = daoPassword.New(context.Background(), tx, "848577e0-9ec9-5d25-bcb8-8ed913100609", &pass); err != nil {
		t.Errorf("error was not expected while inserting password: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
	assert.Equal(t, pass.ID, 1)
}

func TestPasswordInvalidateForUser(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()

	mock.ExpectBegin()
	tx, err := db.Begin()
	if err != nil {
		t.Fatalf("*Begin.Tx ERROR* %s", err)
	}

	mock.ExpectExec(`UPDATE "fnd_password"`).
		WithArgs(base.StatusInactive, "848577e0-9ec9-5d25-bcb8-8ed913100609", base.StatusActive).
		WillReturnResult(sqlmock.NewResult(0, 10))

	if err = daoPassword.InvalidateForUser(context.Background(), tx, "848577e0-9ec9-5d25-bcb8-8ed913100609"); err != nil {
		t.Errorf("error was not expected while inserting password: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
