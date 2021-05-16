package base

import (
	"encoding/json"
	"fmt"
	"reflect"

	"github.com/jtorz/phoenix-backend/utils/numbers"
)

// JSONObject represents a basic json object.
type JSONObject map[string]interface{}

// GetInt gets the int value from the data
func (j JSONObject) GetInt(key string) (v int, err error) {
	i64, err := j.GetInt64(key)
	return int(i64), err
}

// GetInt64 gets the int64 value from the data
func (j JSONObject) GetInt64(key string) (v int64, err error) {
	val, exist := j[key]
	if !exist {
		return 0, fmt.Errorf("JSONObject[%#v] not found", key)
	}
	i64, err := numbers.IToInt64(val)
	if err != nil {
		return 0, fmt.Errorf("JSONObject[%#v] int cast error: %w", key, err)
	}
	return i64, nil
}

// GetString gets the string value from the data
func (j JSONObject) GetString(key string) (string, error) {
	val, exist := j[key]
	if !exist {
		return "", fmt.Errorf("JSONObject[%#v] not found", key)
	}
	return fmt.Sprint(val), nil
}

// RawMessage json marshal.
func (j JSONObject) RawMessage() (raw json.RawMessage, err error) {
	return json.Marshal(j)
}

// Scan implement the database/sql.Scanner interface.
// Scan assigns a value from a database driver.
//
// The src value will be of one of the following types:
//
//    int64
//    float64
//    bool
//    []byte
//    string
//    time.Time
//    nil - for NULL values
//
// An error should be returned if the value cannot be stored
// without loss of information.
//
// Reference types such as []byte are only valid until the next call to Scan
// and should not be retained. Their underlying memory is owned by the driver.
// If retention is necessary, copy their values before the next call to Scan.
func (j *JSONObject) Scan(value interface{}) error {
	switch v := value.(type) {
	case []byte:
		return json.Unmarshal(v, j)
	case string:
		return json.Unmarshal([]byte(v), j)
	case nil:
		*j = make(JSONObject)
		return nil
	default:
		return fmt.Errorf("error scanning modelJSONObject: unknown type %v", reflect.TypeOf(value))
	}
}
