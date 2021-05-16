package numbers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIToInt64(t *testing.T) {
	type s int64
	cases := []struct {
		expected      int64
		data          interface{}
		expectedError string
	}{
		{0, "asdas", `strconv.ParseInt: parsing "asdas": invalid syntax`},
		{3, "3", ""},
		{3, 3, ""},
		{64, int64(64), ""},
		{64, s(64), ""},
		{1, true, ""},
		{0, false, ""},
	}
	for _, c := range cases {
		actual, err := IToInt64(c.data)
		assert.EqualValues(t, c.expected, actual, "interface to float64")
		if c.expectedError == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, c.expectedError, "interface to float64")
		}
	}
}

func TestIToFloat64(t *testing.T) {
	cases := []struct {
		expected      float64
		data          interface{}
		expectedError string
	}{
		{0, "asdas", `strconv.ParseFloat: parsing "asdas": invalid syntax`},
		{3.45, "3.45", ""},
		{3.45, 3.45, ""},
		{3, 3, ""},
		{1, true, ""},
		{0, false, ""},
	}
	for _, c := range cases {
		actual, err := IToFloat64(c.data)
		assert.EqualValues(t, c.expected, actual, "interface to float64")
		if c.expectedError == "" {
			assert.NoError(t, err)
		} else {
			assert.EqualError(t, err, c.expectedError, "interface to float64")
		}
	}
}
