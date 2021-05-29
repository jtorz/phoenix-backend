package baseerrors

import (
	"database/sql"
	"errors"

	"github.com/jtorz/phoenix-backend/utils/pg"
)

// ErrInvalidData is used to indicate that the format or the content of the data is not valid.
//
// Examples of ErrDataFormat is a invalid email, an age less than 0.
var ErrInvalidData = errors.New("invalid data")

// IsErrInvalidData checks if the error is ErrInvalidData.
func IsErrInvalidData(err error) bool {
	return errors.Is(err, ErrInvalidData)
}

// ErrActionNotAllowedStatus is used whe an action can't be used due to the current
// status of a record.
var ErrActionNotAllowedStatus = errors.New("action not allowed in current state")

// IsErrStatus checks if the error is ErrStatus.
func IsErrStatus(err error) bool {
	return errors.Is(err, ErrActionNotAllowedStatus)
}

// ErrPrivilege the user doesn't have the privileges to execute an operation.
var ErrPrivilege = errors.New("insufficient privileges")

// IsErrPrivilege checks if the error is ErrPrivilege.
func IsErrPrivilege(err error) bool {
	return errors.Is(err, ErrPrivilege)
}

// ErrAuth an error ocurred in the authentication o authorization.
// Can be that the user was not found, or that the password didn't match.
var ErrAuth = errors.New("auth error")

// IsErrAuth checks if the error is ErrAuth.
func IsErrAuth(err error) bool {
	return errors.Is(err, ErrAuth)
}

// ErrDuplicated while inserting a record that already exists.
var ErrDuplicated = errors.New("record duplicated")

// ErrNotUpdated the record was not updated.
var ErrNotUpdated = errors.New("record not updated")

// IsErrNotUpdated checks if the error is ErrNotUpdated.
func IsErrNotUpdated(err error) bool {
	return errors.Is(err, ErrNotUpdated)
}

// ErrMultiUpdated multiple records were updated.
var ErrMultiUpdated = errors.New("multiple records updated")

// IsErrMultiUpdated checks if the error is ErrMultiUpdated.
func IsErrMultiUpdated(err error) bool {
	return errors.Is(err, ErrMultiUpdated)
}

// IsErrDuplicated checks if the error is ErrDuplicated.
func IsErrDuplicated(err error) bool {
	if errors.Is(err, ErrDuplicated) {
		return true
	}
	return pg.IsCode(err, pg.UniqueViolation)
}

// ErrNotFound the requested information was not found.
var ErrNotFound = errors.New("not found")

// IsErrNotFound checks if the error is ErrNotFound.
func IsErrNotFound(err error) bool {
	return errors.Is(err, sql.ErrNoRows) || errors.Is(err, ErrNotFound) || pg.IsCode(err, pg.NoDataFound)
}
