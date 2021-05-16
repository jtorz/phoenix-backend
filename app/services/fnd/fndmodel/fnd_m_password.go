package fndmodel

import (
	"fmt"
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
	"github.com/jtorz/phoenix-backend/app/shared/baseerrors"
	"github.com/jtorz/phoenix-backend/utils/passcrypter"
)

// PasswordType pasword encription method.
type PasswordType string

const (
	//PassTypeScrypt2017 scrypt encryption algorithm.
	PassTypeScrypt2017 PasswordType = "Scrypt2017"
)

// Password user password.
type Password struct {
	ID               int
	Type             PasswordType
	Data             base.JSONObject
	InvalidationDate time.Time
	CreatedAt        time.Time
	UpdatedAt        time.Time
	Status           base.Status
}

// NewPassword creates the encrypted version of the password.
func NewPassword(passType PasswordType, password string) (*Password, error) {
	switch passType {
	case PassTypeScrypt2017:
		return newPasswordScrypt2017(passType, password)
	}
	return nil, fmt.Errorf("encryption algorithm %#v not implemented", passType)
}

// ComparePassword compares the clear  password text against the encripted information.
func (p Password) ComparePassword(passwordText string) (err error) {
	switch p.Type {
	case PassTypeScrypt2017:
		return p.comparePasswordScrypt2017(passwordText)
	}
	return fmt.Errorf("encryption algorithm %#v not implemented", p.Type)
}

// Scrypt2017 -
func newPasswordScrypt2017(passType PasswordType, password string) (*Password, error) {
	p := &Password{Type: passType}
	p.Data = make(base.JSONObject)
	scrypt2017 := passcrypter.Scrypt2017()
	hash, salt, err := scrypt2017.Encrypt(password)
	if err != nil {
		return nil, err
	}
	p.Data["password"] = hash
	p.Data["salt"] = salt
	return p, nil
}

func (p Password) comparePasswordScrypt2017(passwordText string) (err error) {
	hash, err := p.Data.GetString("password")
	if err != nil {
		return
	}
	salt, err := p.Data.GetString("salt")
	if err != nil {
		return
	}
	scrypt2017 := passcrypter.Scrypt2017()
	isValid, err := scrypt2017.Compare(passwordText, hash, salt)
	if err != nil {
		return
	}
	if !isValid {
		return baseerrors.ErrAuth
	}
	return nil
}
