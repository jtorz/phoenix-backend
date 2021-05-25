package lex

import (
	"database/sql/driver"

	"github.com/jtorz/phoenix-backend/app/shared/lex/convert"
)

// ZeroString represents a string that may be null.
// ZeroString implements the Scanner interface.
// It asumes that a empty string is null.
type ZeroString string

// Scan implements the Scanner interface.
func (s *ZeroString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	return convert.ConvertAssign(s, value)
}

// Value implements the driver Valuer interface.
func (ns ZeroString) Value() (driver.Value, error) {
	if ns == "" {
		return nil, nil
	}
	return string(ns), nil
}
