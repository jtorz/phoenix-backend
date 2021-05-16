package pg

import "database/sql"

// IntZero creates a null int checking if its zero value.
func IntZero(i int) sql.NullInt64 {
	return sql.NullInt64{
		Int64: int64(i),
		Valid: i != 0,
	}
}

// StringZero creates a null string checking if its zero value.
func StringZero(s string) sql.NullString {
	return sql.NullString{
		String: s,
		Valid:  s != "",
	}
}
