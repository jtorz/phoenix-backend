package stringset

import (
	"encoding/hex"
	"errors"
	"hash/crc32"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

// ErrEmptyString empty string error.
var ErrEmptyString = errors.New("empty string")

// NewHashedToken genera un nombre de un token que puede ser utilizado en codigo mas un hash
//
// hellö thëré => hello_t_<HASH>
// 9 cats => n9_cats_<HASH>
func NewHashedToken(base string) string {
	if base == "" {
		panic(ErrEmptyString)
	}
	r := CleanString(base)
	if unicode.IsDigit(rune(r[0])) {
		r = "n" + r
	}
	return newToken(r, true)
}

// NewToken genera un nombre de un token que puede ser utilizado en codigo
//
// hellö thëré => hello_t
// 9 cats => n9_cats
func NewToken(base string) string {
	if base == "" {
		panic(ErrEmptyString)
	}
	r := CleanString(base)
	if unicode.IsDigit(rune(r[0])) {
		r = "n" + r
	}
	return newToken(r, false)
}

// CleanString tranforma una cadena con acento y caracteres especiales a una sin acento ni caracteres especiales.
//
// Hellö\ntHëRé! =>  hello there
func CleanString(s string) string {
	chain := transform.Chain(
		norm.NFD,
		newRemover(isSpecial),
		runes.Map(cleanSimple),
		norm.NFC)
	result, _, _ := transform.String(chain, s)
	return result
}

// SanitizePGToken elimina todos los caracteres especiales, incluyendo espacios, excepto por guion bajo (_) y pesos ($).
func SanitizePGToken(t string) string {
	if t == "" {
		panic(ErrEmptyString)
	}
	chain := transform.Chain(
		newRemover(func(r rune) bool {
			return !(unicode.IsLetter(r) || unicode.IsDigit(r) || r == '_' || r == '$')
		}))
	result, _, _ := transform.String(chain, t)
	if result[0] == '$' {
		return result[1:]
	}
	return result
}

func newToken(base string, useHash bool) string {
	ss := strings.Split(base, " ")
	sb := &strings.Builder{}
	sb.Grow(15)
	n := setBaseToken(sb, ss)
	if useHash {
		sb.WriteRune('_')
		n = 14 - n
		sb.WriteString(getHashN(time.Now().Format("010215040506"), n))
	}
	return sb.String()
}

func getHashN(s string, n int) string {
	return FirstN(getHash(s), n)
}

func getHash(s string) string {
	hasher := crc32.NewIEEE()
	hasher.Write([]byte(s))
	return hex.EncodeToString(hasher.Sum(nil))
}

func setBaseToken(sb *strings.Builder, base []string) int {
	var n, w int
	lenbase := len(base)
	if lenbase == 1 {
		n, _ = sb.WriteString(FirstN(base[0], 10))
		return n
	}

	n, _ = sb.WriteString(FirstN(base[0], 5))

	sb.WriteRune('_')
	n++

	if lenbase == 2 {
		w, _ = sb.WriteString(FirstN(base[1], 5))
		return n + w
	}

	w, _ = sb.WriteString(FirstN(base[1], 2))
	n += w
	sb.WriteRune('_')
	n++
	w, _ = sb.WriteString(FirstN(base[2], 1))
	return n + w
}

func cleanSimple(r rune) rune {
	if unicode.IsSpace(r) {
		return ' '
	}
	if unicode.IsLetter(r) {
		return unicode.ToLower(r)
	}
	if unicode.IsDigit(r) {
		return r
	}
	panic("Special char")
}

func newRemover(shouldRemove func(rune) bool) runes.Transformer {
	return runes.Remove(setFunc(shouldRemove))
}

func isSpecial(r rune) bool {
	return !unicode.IsSpace(r) && !unicode.IsLetter(r) && !unicode.IsDigit(r)
}

// setFunc implements runes.Set
// A Set is a collection of runes.
type setFunc func(rune) bool

// Contains returns true if r is contained in the set.
func (s setFunc) Contains(r rune) bool {
	return s(r)
}
