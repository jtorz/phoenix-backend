package httphandler

import (
	"database/sql"
)

// Tx sql.Tx wrapper
type Tx struct {
	Tx *sql.Tx
}

// Commit commits the transaction if an error occurs it panics
func (tx *Tx) Commit(c Handler) {
	if err := tx.Tx.Commit(); err != nil {
		panic(err)
	}
	c.Set(KeyRequestTx, nil)
}

// Rollback commits the transaction if an error occurs it panics
func (tx *Tx) Rollback(c Handler) {
	if err := tx.Tx.Rollback(); err != nil {
		panic(err)
	}
	c.Set(KeyRequestTx, nil)
}
