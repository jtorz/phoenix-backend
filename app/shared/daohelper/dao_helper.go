package daohelper

import (
	"database/sql"
	"fmt"
	"runtime"

	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/utils/pg"
)

// CheckOneRowUpdated checks that only one records was affected
func (QueryHelper) CheckOneRowUpdated(r sql.Result) error {
	n, err := r.RowsAffected()
	if err != nil {
		pc := make([]uintptr, 10)
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		file, line := f.FileLine(pc[0])
		data := fmt.Sprintf("%s:%d %s/n", file, line, f.Name())
		return fmt.Errorf("%w %s", err, data)
	}
	if n != 1 {
		pc := make([]uintptr, 10)
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		file, line := f.FileLine(pc[0])
		data := fmt.Sprintf("%s:%d %s/n", file, line, f.Name())
		if n == 0 {
			return fmt.Errorf("%s %w", data, baseerrors.ErrNotUpdated)
		}
		return fmt.Errorf("%s %w", data, baseerrors.ErrMultiUpdated)
	}
	return nil
}

// WrapIfErrDuplicated wraps the error into baseerrors.ErrDuplicated er the underlying error is due to a unique key volation.
func (QueryHelper) WrapIfErrDuplicated(err error) error {
	if pg.IsCode(err, pg.UniqueViolation) {
		return fmt.Errorf("%s %w", err.Error(), baseerrors.ErrDuplicated)
	}
	return err
}

// WrapErr wraps the error with extra information if ocurred.
func (QueryHelper) WrapErr(err error) error {
	if err != nil {
		pc := make([]uintptr, 10) // at least 1 entry needed
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		file, line := f.FileLine(pc[0])
		data := fmt.Sprintf("%s:%d %s/n", file, line, f.Name())

		return fmt.Errorf("%s %w", data, err)
	}
	return nil
}
