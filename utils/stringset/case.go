package stringset

import (
	"bytes"
	"strings"
)

var uppercaseWords = map[string]struct{}{
	"ascii": {},
	"guid":  {},
	"id":    {},
	"ip":    {},
	"json":  {},
	"ram":   {},
	"uid":   {},
	"uuid":  {},
	"url":   {},
	"utf8":  {},
}

// UpperFirst transforms the first letter to upper case, the rest is left as is.
func UpperFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToUpper(string(s[0])) + s[1:]
}

// LowerFirst transforms the first letter to lower case, the rest is left as is.
func LowerFirst(s string) string {
	if s == "" {
		return ""
	}
	return strings.ToLower(string(s[0])) + s[1:]
}

// SnakeToGoCase changes a snake-case variable name
// into a go styled object variable name of "ColumnName".
// titleCase also fully uppercases "ID" components of names, for example
// "column_name_id" to "ColumnNameID".
// TitleCase changes a snake-case variable name
// into a go styled object variable name of "ColumnName".
// titleCase also fully uppercases "ID" components of names, for example
// "column_name_id" to "ColumnNameID".
//
// Note: This method is ugly because it has been highly optimized,
// we found that it was a fairly large bottleneck when we were using regexp.
//
// Code taken from github.com/volatiletech/sqlboiler/strmangle.ToTitle
func SnakeToGoCase(n string) string {

	ln := len(n)
	name := []byte(n)
	buf := &bytes.Buffer{}

	start := 0
	end := 0
	for start < ln {
		// Find the start and end of the underscores to account
		// for the possibility of being multiple underscores in a row.
		if end < ln {
			if name[start] == '_' {
				start++
				end++
				continue
				// Once we have found the end of the underscores, we can
				// find the end of the first full word.
			} else if name[end] != '_' {
				end++
				continue
			}
		}

		word := name[start:end]
		wordLen := len(word)
		var vowels bool

		numStart := wordLen
		for i, c := range word {
			vowels = vowels || (c == 97 || c == 101 || c == 105 || c == 111 || c == 117 || c == 121)

			if c > 47 && c < 58 && numStart == wordLen {
				numStart = i
			}
		}

		_, match := uppercaseWords[string(word[:numStart])]

		if match {
			// Uppercase all a-z characters
			for _, c := range word {
				if c > 96 && c < 123 {
					buf.WriteByte(c - 32)
				} else {
					buf.WriteByte(c)
				}
			}
		} else {
			if c := word[0]; c > 96 && c < 123 {
				buf.WriteByte(word[0] - 32)
				buf.Write(word[1:])
			} else {
				buf.Write(word)
			}
		}

		start = end + 1
		end = start
	}

	ret := buf.String()
	return ret
}
