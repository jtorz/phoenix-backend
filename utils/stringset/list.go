package stringset

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

// ListToSlice converts the list to string slice.
func ListToSlice(l string, delim byte) []string {
	if l == "" {
		return []string{}
	}
	return strings.Split(string(l), string(delim))
}

// ListFromIntSlice creates a List from the int slice
// Example: []int{4,7,8,10} --> "4,7,8,10"
func ListFromIntSlice(ns []int, delim byte) string {
	l := len(ns)
	if l == 0 {
		return ""
	}

	b := strings.Builder{}
	// Approx 2 chars per num plus the comma.
	b.Grow(len(ns) * 3)

	for i := 0; i < l-1; i++ {
		b.WriteString(strconv.Itoa(ns[i]))
		b.WriteByte(delim)
	}
	b.WriteString(strconv.Itoa(ns[l-1]))
	return b.String()
}

//ListFromSlice creates a List from the slice or array
func ListFromSlice(a interface{}, delim byte) string {
	val := reflect.ValueOf(a)
	if val.Kind() == reflect.Array || val.Kind() == reflect.Slice {
		l := val.Len()
		if l == 0 {
			return ""
		}
		if l == 1 {
			return fmt.Sprint(val.Index(0))
		}
		sb := strings.Builder{}
		sb.Grow(l * 4)
		sb.WriteString(fmt.Sprint(val.Index(0)))
		for i := 1; i < l; i++ {
			sb.WriteByte(delim)
			sb.WriteString(fmt.Sprint(val.Index(i)))
		}
		return sb.String()
	}

	return ""
}
