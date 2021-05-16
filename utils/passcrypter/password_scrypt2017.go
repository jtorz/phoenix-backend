package passcrypter

import (
	"crypto/rand"
	"fmt"
	"io"

	"golang.org/x/crypto/scrypt"
)

// ScryptCrypter golang.org/x/crypto/scrypt
//
// Key derives a key from the password, salt, and cost parameters, returning
// a byte slice of length keyLen that can be used as cryptographic key.
//
// N is a CPU/memory cost parameter, which must be a power of two greater than 1.
// r and p must satisfy r * p < 2³⁰. If the parameters do not satisfy the
// limits, the function returns a nil byte slice and an error.
//
// For example, you can get a derived key for e.g. AES-256 (which needs a
// 32-byte key) by doing:
//
//      dk, err := scrypt.Key([]byte("some password"), salt, 32768, 8, 1, 32)
//
type ScryptCrypter interface {
	Encrypt(clearText string) (hash, salt string, err error)
	Compare(clearText, hash, salt string) (bool, error)
}

func Scrypt2017() ScryptCrypter {
	return scrypt2017{}
}

type scrypt2017 struct{}

// Encrypt encrypts the clear text with the scrypt algorithm.
// The recommended parameters for interactive logins as of 2017 are N=32768, r=8
// and p=1. The parameters N, r, and p should be increased as memory latency and
// CPU parallelism increases; consider setting N to the highest power of 2 you
// can derive within 100 milliseconds. Remember to get a good random salt.
func (s scrypt2017) Encrypt(clearText string) (hash, salt string, err error) {
	salt, err = newSalt()
	if err != nil {
		return
	}
	dk, err := scrypt.Key([]byte(clearText), []byte(salt), 32768, 8, 1, 256)
	if err != nil {
		return
	}
	hash = fmt.Sprintf("%x", dk)
	return
}

func newSalt() (string, error) {
	salt := make([]byte, 128)
	_, err := io.ReadFull(rand.Reader, salt)
	if err != nil {
		return "", err
	}
	return fmt.Sprintf("%x", salt), nil
}

func (s scrypt2017) Compare(clearText, hash, salt string) (bool, error) {
	dk, err := scrypt.Key([]byte(clearText), []byte(salt), 32768, 8, 1, 256)
	if err != nil {
		return false, err
	}
	hash2 := fmt.Sprintf("%x", dk)
	return hash == hash2, nil
}
