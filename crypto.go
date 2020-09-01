package library

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"io"
)

func GetAuthKey() []byte {
	// config := config.GetConfig()
	// /authKey := config.GetString("app.cKey")
	return Hash("authKey") // key will be set
}

//Hash for AES key
func Hash(key string) []byte {
	hash := sha256.Sum256([]byte(key))
	return hash[:]
}

//EncryptAES returns AES crypted bytes
func EncryptAES(key []byte, text []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)

	if err != nil {

		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {

		return nil, err
	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {

	}

	return gcm.Seal(nonce, nonce, text, nil), nil
}

//DecryptAES return decrypted text
func DecryptAES(key []byte, cipherText []byte) ([]byte, error) {
	c, err := aes.NewCipher(key)
	if err != nil {

		return nil, err
	}

	gcm, err := cipher.NewGCM(c)
	if err != nil {

		return nil, err
	}

	nonceSize := gcm.NonceSize()
	if len(cipherText) < nonceSize {
		return nil, errors.New("cipherText's length is less than nonce's length")
	}

	nonce, cipherText := cipherText[:nonceSize], cipherText[nonceSize:]
	plaintext, err := gcm.Open(nil, nonce, cipherText, nil)
	if err != nil {

		return nil, err
	}
	return plaintext, nil
}
