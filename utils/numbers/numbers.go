package numbers

import (
	"fmt"
	"reflect"
	"strconv"
)

// IToInt64 interface to int return the int value for the given.
func IToInt64(i interface{}) (int64, error) {
	switch v := i.(type) {
	case int64:
		return v, nil
	case string:
		return strconv.ParseInt(v, 10, 64)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	case int, int8, int16, uint, uint16, uint32, uint64, byte, rune:
		return reflect.ValueOf(i).Int(), nil
	default:
		f, err := strconv.ParseFloat(fmt.Sprint(i), 64)
		return int64(f), err
	}
}

// IToFloat64 converts the interface value to float64.
func IToFloat64(i interface{}) (float64, error) {
	switch v := i.(type) {
	case float64, float32:
		return v.(float64), nil
	case string:
		return strconv.ParseFloat(v, 64)
	case bool:
		if v {
			return 1, nil
		}
		return 0, nil
	default:
		return strconv.ParseFloat(fmt.Sprint(i), 64)
	}
}
