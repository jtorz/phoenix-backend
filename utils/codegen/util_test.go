package codegen

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestGoCase(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{input: "core_user", expected: "CoreUser"},
		{input: "user", expected: "User"},
		{input: "mail_b_record", expected: "MailBRecord"},
	}
	for _, test := range tests {
		actual := goCase(test.input)
		assert.Equal(t, test.expected, actual)
	}
}

func TestDbTableNameToGoName(t *testing.T) {
	tests := []struct {
		input, expected string
	}{
		{input: "mail_b_record", expected: "BRecord"},
		{input: "mail__b_record", expected: "BRecord"},
		{input: "core_user", expected: "User"},
	}
	for _, test := range tests {
		actual := dbTableNameToGoName(test.input)
		assert.Equal(t, test.expected, actual)
	}

}
