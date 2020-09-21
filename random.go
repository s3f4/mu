package mu

import (
	"math/rand"
	"time"
)

// charsetType is a string type that choose
// which type of characters will be produced
type charsetType string

const (
	// Alpha only alpha characters
	Alpha charsetType = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	// Numeric only numeric characters
	Numeric = "0123456789"
	// AlphaNumeric alphanumeric characters
	AlphaNumeric = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

var charset string

// SetCharset with charsettype
func SetCharset(ctype charsetType) {
	charset = string(ctype)
}

func getCharset() string {
	if charset == "" {
		return string(AlphaNumeric)
	}
	return charset
}

// RandomString creates a random string
func RandomString(length int) string {
	seed := rand.New(rand.NewSource(time.Now().UnixNano()))
	randomBytes := make([]byte, length)
	charset := getCharset()
	for i := range randomBytes {
		randomBytes[i] = charset[seed.Intn(len(charset))]
	}
	return string(randomBytes)
}
