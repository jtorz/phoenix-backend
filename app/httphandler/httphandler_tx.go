package httphandler

import (
	"database/sql"
)

// BeginTx begins a sql transaction, if an error occurs it panics.
//
// The KeyRequestTx value is added to specify that a transaction has been initiated.
//
// If the transaction is not actively Commited or Rollbacked in the http handler function
// the transaction is rollbacked in a middleware.
// The transaction is rollbacked as well when a panic occurs.
func (c *Context) BeginTx(db *sql.DB) *Tx {
	tx, err := db.Begin()
	if err != nil {
		panic(err)
	}
	c.Set(KeyRequestTx, tx)
	return &Tx{tx}
}

// Tx sql.Tx wrapper created to handle transactions during a http handler function.
type Tx struct {
	Tx *sql.Tx
}

// Commit commits the transaction if an error occurs it panics.
//
// The KeyRequestTx value is modified to specify that the transaction has been commited.
func (tx *Tx) Commit(c *Context) {
	if err := tx.Tx.Commit(); err != nil {
		panic(err)
	}
	c.Set(KeyRequestTx, nil)
}

// Rollback commits the transaction if an error occurs it panics.
//
// The KeyRequestTx value is modified to specify that the transaction has been rollbacked.
func (tx *Tx) Rollback(c *Context) {
	if err := tx.Tx.Rollback(); err != nil {
		panic(err)
	}
	c.Set(KeyRequestTx, nil)
}
