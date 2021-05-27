package lex

import (
	"database/sql/driver"

	"github.com/jtorz/phoenix-backend/app/shared/lex/convert"
)

// ZeroString represents a string that may be null.
// ZeroString implements the Scanner interface
// where a Zero value string is null.
type ZeroString string

// Scan implements the Scanner interface.
func (s *ZeroString) Scan(value interface{}) error {
	if value == nil {
		*s = ""
		return nil
	}
	var ss string
	if err := convert.ConvertAssign(&ss, value); err != nil {
		return err
	}
	*s = ZeroString(ss)
	return nil
}

// Value implements the driver Valuer interface.
func (ns ZeroString) Value() (driver.Value, error) {
	if ns == "" {
		return nil, nil
	}
	return string(ns), nil
}
