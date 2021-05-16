package files

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFileNameExt(t *testing.T) {
	cases := []struct {
		full         string
		expectedName string
		expectedExt  string
	}{
		{"", "", ""},
		{"file.ext", "file", "ext"},
		{"file.2019.ext", "file.2019", "ext"},
		{".file.2019.ext", ".file.2019", "ext"},
		{"file", "file", ""},
		{".file", ".file", ""},
	}

	for _, c := range cases {
		actualName, actualExt := SplitNameExt(c.full)
		assert.EqualValues(t, c.expectedName, actualName, "file name comparison")
		assert.EqualValues(t, c.expectedExt, actualExt, "file extension comparison")
	}
}

func TestJoinNameExt(t *testing.T) {
	cases := []struct {
		expected string
		name     string
		ext      string
	}{
		{"", "", ""},
		{"file.ext", "file", "ext"},
		{"file.2019.ext", "file.2019", "ext"},
		{".file.2019.ext", ".file.2019", "ext"},
		{"file", "file", ""},
		{".file", ".file", ""},
	}

	for _, c := range cases {
		actual := JoinNameExt(c.name, c.ext)
		assert.EqualValues(t, c.expected, actual, "file name join")
	}
}
