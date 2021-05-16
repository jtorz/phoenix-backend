package jsonutil

import (
	"encoding/json"
	"errors"
	"io"
	"io/ioutil"
	"strings"

	"github.com/xeipuuv/gojsonschema"
)

// ReadValidateJSON reatrives the json and the schema from the reader and compares such json against the schema.
// returns the json if its well formed.
func ReadValidateJSON(jsonReader, schemaReader io.Reader) ([]byte, error) {
	schema, err := ioutil.ReadAll(schemaReader)
	if err != nil {
		return nil, err
	}
	j, err := ioutil.ReadAll(jsonReader)
	if err != nil {
		return nil, err
	}
	err = ValidateJSON(j, schema)
	if err != nil {
		return nil, err
	}
	return j, nil
}

// ValidateJSON validates the json against the schema.
func ValidateJSON(j, schema []byte) error {
	schemaLoader := gojsonschema.NewBytesLoader(schema)
	jsonLoader := gojsonschema.NewBytesLoader(j)
	result, err := gojsonschema.Validate(schemaLoader, jsonLoader)
	if err != nil {
		return err
	}
	if result.Valid() {
		return nil
	}

	errores := result.Errors()
	errStr := strings.Builder{}
	errStr.Grow(14)
	errStr.WriteString("invalid json [")
	s, _ := json.Marshal(errores[0].String())
	errStr.Grow(len(s))
	errStr.Write(s)
	for _, resultError := range errores[1:] {
		s, _ := json.Marshal(resultError.String())
		errStr.Grow(len(s) + 1)
		errStr.WriteByte(',')
		errStr.Write(s)
	}
	errStr.WriteRune(']')

	return errors.New(errStr.String())
}
