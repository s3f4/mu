package mu

import (
	"bytes"
	"crypto/sha256"
	"testing"

	"github.com/stretchr/testify/assert"
)

func SetText() ([]byte, []byte) {
	hash := sha256.Sum256([]byte("key"))
	key := hash[:]
	text := []byte(" $$$$ #### some text #### $$$$ (^#@!!!&*(-_+))")
	return key, text
}

//AES testing
func TestAES(t *testing.T) {
	key, text := SetText()

	encrypted, err := EncryptAES(key, text)
	assert.Nil(t, err, "encryption error %v\n", err)

	decrypted, err := DecryptAES(key, encrypted)
	assert.Nil(t, err, "decryption error %v\n", err)

	if bytes.Compare(decrypted, text) != 0 {
		t.Error("AES encryption does not work correctly.")
		t.Errorf("text:%s\n\ndecrypted:%s\n ", text, decrypted)
	}

	assert.Equal(t, " $$$$ #### some text #### $$$$ (^#@!!!&*(-_+))", string(decrypted), "Decrypted and string is not equal")
}
