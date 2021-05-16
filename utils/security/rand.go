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

//RandStringAlfa genera un string alfabetico aleatorio del tamaño especificado
func RandStringAlfa(n int) string {
	return randString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ", n)
}

//RandStringAlfaNum genera un string alfanumerico aleatorio del tamaño especificado
func RandStringAlfaNum(n int) string {
	return randString("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", n)
}

//RandString genera un string alfanumerico con caracteres especiales aleatorio del tamaño especificado
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
