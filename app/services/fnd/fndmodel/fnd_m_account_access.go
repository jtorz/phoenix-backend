package fndmodel

import (
	"crypto/rand"
	"fmt"
	"time"

	"github.com/jtorz/phoenix-backend/app/shared/base"
	"golang.org/x/crypto/scrypt"
)

// AccountAccessType account access handled by the system.
type AccountAccessType string

const (
	// AccAccRestoreAccount Access used to activate a new account or restore a password.
	AccAccRestoreAccount AccountAccessType = "RestoreAccount"
)

// AccountAccess can be used as alternate authentication methods.
type AccountAccess struct {
	Type      AccountAccessType
	Key       string
	User      User
	CreatedAt time.Time   `json:"createdAt"`
	UpdatedAt time.Time   `json:"updatedAt"`
	Status    base.Status `json:"status"`
}

// NewAccountAccess creates a new account acces for the user
func NewAccountAccess(u User, k AccountAccessType) (*AccountAccess, error) {
	key := make([]byte, 128)
	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}
	key, err = scrypt.Key([]byte(u.ID), key, 2048, 8, 1, 64)
	if err != nil {
		return nil, err
	}
	return &AccountAccess{
		Type:   k,
		Key:    fmt.Sprintf("%x", key),
		User:   u,
		Status: base.StatusActive,
	}, nil
}
