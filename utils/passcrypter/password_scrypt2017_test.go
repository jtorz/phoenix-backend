package passcrypter_test

import (
	"testing"

	"github.com/jtorz/phoenix-backend/utils/passcrypter"
	"github.com/stretchr/testify/assert"
)

func TestScrypt2017(t *testing.T) {
	scrypt2017 := passcrypter.Scrypt2017()
	hash, salt, err := scrypt2017.Encrypt("hello")
	assert.Nil(t, err)
	if err != nil {
		return
	}

	ok, err := scrypt2017.Compare("hello", hash, salt)
	assert.Nil(t, err)
	assert.True(t, ok)

	ok, err = scrypt2017.Compare("hello.", hash, salt)
	assert.Nil(t, err)
	assert.False(t, ok)
}
