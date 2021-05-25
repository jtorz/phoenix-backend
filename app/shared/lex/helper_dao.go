// Package lex contains the dictionary (lexicon) of the database and auxiliar functions to the daos.
//
// The elements in the package are:
//
// lex_object_names.go
//	Table names
//	View names
//	FK Constraints join expressions
//
//  lex_object_columns.go
//	* Table columns
//	* View columns
package lex

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"runtime"

	"github.com/jtorz/phoenix-backend/app/config"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/app/shared/ctxinfo"
	"github.com/jtorz/phoenix-backend/utils/pg"
)

// CheckOneRowUpdated checks that only one records was affected.
func CheckOneRowUpdated(ctx context.Context, name string, r sql.Result) error {
	n, err := r.RowsAffected()
	if err != nil {
		if ctxinfo.LogginAllowed(ctx, config.LogDebug) {
			// Added stack trace info only for debug.
			// code repeated to avoid adding info to the stack.
			pc := make([]uintptr, 10)
			runtime.Callers(2, pc)
			f := runtime.FuncForPC(pc[0])
			file, line := f.FileLine(pc[0])
			log.Printf("%s:%d %s/n", file, line, f.Name())
		}
		return err
	}
	if n == 1 {
		return nil
	}
	if n == 0 {
		err = baseerrors.ErrNotUpdated
	}
	if n > 1 {
		err = baseerrors.ErrMultiUpdated
	}
	return fmt.Errorf("%s %w", name, err)

}

// WrapIfErrDuplicated wraps the error into baseerrors.ErrDuplicated er the underlying error is due to a unique key volation.
func WrapIfErrDuplicated(err error) error {
	if pg.IsCode(err, pg.UniqueViolation) {
		return fmt.Errorf("%s %w", err.Error(), baseerrors.ErrDuplicated)
	}
	return err
}

// DebugErr logs the information of the error with extra information if ocurred.
func DebugErr(ctx context.Context, err error) {
	if err == nil {
		return
	}
	if ctxinfo.LogginAllowed(ctx, config.LogDebug) {
		// Added stack trace info only for debug.
		// code repeated to avoid adding info to the stack.
		pc := make([]uintptr, 10)
		runtime.Callers(2, pc)
		f := runtime.FuncForPC(pc[0])
		file, line := f.FileLine(pc[0])
		log.Printf("%s:%d %s/n", file, line, f.Name())
	}
}
