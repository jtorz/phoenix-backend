package fndmodel

import (
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
)

// User struct.
type User struct {
	ID         string `json:"id"`
	Name       string `json:"name"`
	MiddleName string `json:"middleName"`
	LastName   string `json:"lastName"`
	Email      string `json:"email"`
	Username   string `json:"username"`
	//UserInformation *UserInformation `json:"userInformation"`
	Password      *Password          `json:"password"`
	Roles         Roles              `json:"role"`
	CreatedAt     time.Time          `json:"createdAt"`
	UpdatedAt     time.Time          `json:"updatedAt"`
	Status        base.Status        `json:"status"`
	RecordActions base.RecordActions `json:"recordActions"`
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
