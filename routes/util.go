package routes

import (
	"crypto/sha512"
	"encoding/base64"
	"fmt"
)

// for decoding passwords
// not used, but it does work :)
func decryptPassword(cryptedPass string) []byte {
	decoded, err := base64.StdEncoding.DecodeString(cryptedPass)
	if err != nil {
		fmt.Println("decoder error:", err)
	}
	return decoded
}

// for encoding passwords
func encryptPassword(password []byte) string {
	h := sha512.New()
	h.Write([]byte(password))
	shaPassword := h.Sum(nil)
	encodedPassword := base64.StdEncoding.EncodeToString(shaPassword)
	return encodedPassword
}
