package stringset

import "strings"

// FirstN returns the first n characters from a non unicode string
func FirstN(s string, n int) string {
	if len(s) > n {
		return s[:n]
	}
	return s
}

// FirstNUnicode returns the first n characters from a unicode string
func FirstNUnicode(s string, n int) string {
	i := 0
	for j := range s {
		if i == n {
			return s[:j]
		}
		i++
	}
	return s
}

// FindInSlice searchs the string in the array and returns its position.
func FindInSlice(a []string, s string) (int, bool) {
	for i := range a {
		if a[i] == s {
			return i, true
		}
	}
	return 0, false
}

// UpperFirst transforms the first letter to upper case, the rest is left as is.
func UpperFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}
