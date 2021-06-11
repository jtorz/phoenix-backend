// Package dalhelper contains the dictionary (lexicon) of the database and auxiliar functions to the dals.
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
package dalhelper

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

// WrapNotFound wrpas the error only if its a sql.ErrNoRows.
func WrapNotFound(ctx context.Context, table string, err error) error {
	if err == nil {
		return nil
	}
	if err == sql.ErrNoRows {
		return &pg.Error{
			Code:    pg.NoDataFound,
			Message: "no data found",
			Table:   table,
		}
	}
	return err
}
