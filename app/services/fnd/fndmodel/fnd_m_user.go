package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// Users slice.
type Users []User

// User struct.
type User struct {
	ID         string
	Name       string
	MiddleName string
	LastName   string
	Email      string
	Username   string
	//UserInformation *UserInformation
	Password      *Password
	Roles         Roles
	CreatedAt     time.Time
	UpdatedAt     time.Time
	Status        base.Status
	RecordActions base.RecordActions
}

// UserID returns the id of the user.
func (u User) UserID() string {
	return u.ID
}

// Privileges returns the privileges that the user has.
func (u User) Privileges() Privileges {
	return u.Roles.Privileges()
}

// ToUser returns the user Struct
func (u User) ToUser() User {
	return u
}
