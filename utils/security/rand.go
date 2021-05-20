package security

import (
	"math/rand"
	"time"
)

const (
	letterIdxBits = 7                    // 6 bits to represent a letter index
	letterIdxMask = 1<<letterIdxBits - 1 // All 1-bits, as many as letterIdxBits
	letterIdxMax  = 63 / letterIdxBits   // # of letter indices fitting in 63 bits
)

// RandStringAlfa generates a random string with n characters with only letters.
func RandStringAlfa(n int) string {
	return randString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", n)
}

// RandStringAlfaNum generates a random string with n characters with only letters and numbers.
func RandStringAlfaNum(n int) string {
	return randString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", n)
}

// RandString generates a random string with n characters with letters, numbers and special printable characters.
func RandString(n int) string {
	return randString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789{|}~[\\]^_@:;<=>?!\"#$%&'()*+,-./", n)
}

func randString(letterBytes string, n int) string {
	var src = rand.NewSource(time.Now().UnixNano())
	b := make([]byte, n)
	// A src.Int63() generates 63 random bits, enough for letterIdxMax characters!
	for i, cache, remain := n-1, src.Int63(), letterIdxMax; i >= 0; {
		if remain == 0 {
			cache, remain = src.Int63(), letterIdxMax
		}
		if idx := int(cache & letterIdxMask); idx < len(letterBytes) {
			b[i] = letterBytes[idx]
			i--
		}
		cache >>= letterIdxBits
		remain--
	}
	return string(b)
}
