package jsonutil

import (
	"bytes"
	"errors"
	"io"
	"testing"

	"github.com/stretchr/testify/assert"
)

type failingReader struct{}

func (r failingReader) Read(p []byte) (n int, err error) {
	return 0, errors.New("error")
}
func TestReadValidateJSON(t *testing.T) {

	cases := []struct {
		json          io.Reader
		schema        io.Reader
		expectedJSON  []byte
		expectedError string
	}{
		{
			bytes.NewBufferString(`"hola"`),
			bytes.NewBufferString(`{"$schema": "https://json-schema.org/draft-07/schema#","type": "object"}`),
			nil,
			`invalid json ["(root): Invalid type. Expected: object, given: string"]`,
		},
		{
			bytes.NewBufferString(`{"alfa":"a"}`),
			bytes.NewBufferString(`{"$schema": "https://json-schema.org/draft-07/schema#","type": "object"}`),
			[]byte(`{"alfa":"a"}`),
			``,
		},
		{
			failingReader{},
			bytes.NewBufferString(`{"$schema": "https://json-schema.org/draft-07/schema#","type": "object"}`),
			nil,
			`error`,
		},
		{
			bytes.NewBufferString(`{"alfa":"a"}`),
			failingReader{},
			nil,
			`error`,
		},
	}

	for _, c := range cases {
		actualJSON, actualErr := ReadValidateJSON(c.json, c.schema)
		assert.EqualValues(t, c.expectedJSON, actualJSON)
		if c.expectedError == "" {
			assert.NoError(t, actualErr)
		} else {
			assert.EqualError(t, actualErr, c.expectedError)
		}
	}
}

func TestValidateJSON(t *testing.T) {
	cases := []struct {
		json          string
		schema        string
		expectedError string
	}{
		{
			`{"alfa":"a","beta":"b"}`,
			``,
			`EOF`,
		},
		{
			`{"alfa":"a","beta":"b"}`,
			`{"$schema": "https://json-schema.org/draft-07/schema#","type": "object"}`,
			``,
		},
		{
			`"hola"`,
			`{"$schema": "https://json-schema.org/draft-07/schema#","type": "object"}`,
			`invalid json ["(root): Invalid type. Expected: object, given: string"]`,
		},
		{
			`[]`,
			`{"$schema": "https://json-schema.org/draft-07/schema#","type": "object"}`,
			`invalid json ["(root): Invalid type. Expected: object, given: array"]`,
		},
		{
			`{"alfa":"a","beta":"b"}`,
			`{"$schema": "https://json-schema.org/draft-07/schema#","type": "object","properties":{"beta":{"type":"object"}},"required":["alfa","beta","gama"]}`,
			`invalid json ["(root): gama is required","beta: Invalid type. Expected: object, given: string"]`,
		},
	}

	for _, c := range cases {
		actualErr := ValidateJSON([]byte(c.json), []byte(c.schema))
		if c.expectedError == "" {
			assert.NoError(t, actualErr)
		} else {
			assert.EqualError(t, actualErr, c.expectedError)
		}
	}
}
