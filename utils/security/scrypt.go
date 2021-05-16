package security

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
)

// NewScrypt generates a new scrypt value from the clear text and the given values.
func NewScrypt(clearText string, N, r, p, keyLen int) (hash, salt string, err error) {
	salt, err = newSalt()
	if err != nil {
		return
	}
	dk, err := scrypt.Key([]byte(clearText), []byte(salt), N, r, p, keyLen)
	if err != nil {
		return
	}
	hash = fmt.Sprintf("%x", dk)
	return
}

// CompareScrypt compares the hash against the new generated hash from the clear text.
func CompareScrypt(hash, clearText, salt string, N, r, p, keyLen int) (bool, error) {
	dk, err := scrypt.Key([]byte(clearText), []byte(salt), N, r, p, keyLen)
	if err != nil {
		return false, err
	}
	hash2 := fmt.Sprintf("%x", dk)
	return hash == hash2, nil
}

func newSalt() (string, error) {
	salt := make([]byte, 128)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", salt), nil
}
