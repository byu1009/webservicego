package jkn

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"encoding/json"
	"errors"

	lzstring "github.com/daku10/go-lz-string"
)

// PKCS7 unpad
func pkcs7Unpad(data []byte) ([]byte, error) {
	if len(data) == 0 {
		return nil, errors.New("invalid padding")
	}
	padding := int(data[len(data)-1])
	if padding > len(data) {
		return nil, errors.New("invalid padding")
	}
	return data[:len(data)-padding], nil
}

// DecryptBPJS = PHP openssl_decrypt + LZString
func DecryptBPJS(encrypted, consID, secretKey, timestamp string) (interface{}, error) {
	key := consID + secretKey + timestamp

	hash := sha256.Sum256([]byte(key))
	keyBytes := hash[:]
	iv := keyBytes[:16]

	cipherText, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(keyBytes)
	if err != nil {
		return nil, err
	}

	if len(cipherText)%aes.BlockSize != 0 {
		return nil, errors.New("cipher text invalid")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	mode.CryptBlocks(cipherText, cipherText)

	plainText, err := pkcs7Unpad(cipherText)
	if err != nil {
		return nil, err
	}

	decompressed, err := lzstring.DecompressFromEncodedURIComponent(string(plainText))
	if err != nil {
		return nil, err
	}

	// flexible decode: array atau object
	var result interface{}
	if err := json.Unmarshal([]byte(decompressed), &result); err != nil {
		return nil, err
	}

	return result, nil
}
