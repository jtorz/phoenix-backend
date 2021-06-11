package coredal_test

import (
	"context"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/jtorz/phoenix-backend/app/services/core/coredal"
	"github.com/jtorz/phoenix-backend/app/services/core/coremodel"
	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/stretchr/testify/assert"
)

var dalPassword = coredal.DalPassword{}

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
	mock.ExpectQuery(`INSERT INTO "core_password"`).WithArgs([]byte("{}"), 2, coremodel.PassTypeScrypt2017, any, "848577e0-9ec9-5d25-bcb8-8ed913100609").WillReturnRows(sqlmock.NewRows([]string{"id"}).AddRow(1))

	// now we execute our method
	pass := coremodel.Password{
		Type:   coremodel.PassTypeScrypt2017,
		Data:   base.JSONObject{},
		Status: 2,
	}
	if err = dalPassword.New(context.Background(), tx, "848577e0-9ec9-5d25-bcb8-8ed913100609", &pass); err != nil {
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

	mock.ExpectExec(`UPDATE "core_password"`).
		WithArgs(base.StatusInactive, "848577e0-9ec9-5d25-bcb8-8ed913100609", base.StatusActive).
		WillReturnResult(sqlmock.NewResult(0, 10))

	if err = dalPassword.InvalidateForUser(context.Background(), tx, "848577e0-9ec9-5d25-bcb8-8ed913100609"); err != nil {
		t.Errorf("error was not expected while inserting password: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
