package bpjs

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha256"
	"encoding/base64"
	"errors"

	lzstring "github.com/daku10/go-lz-string"
)

func pkcs7Unpad(data []byte) ([]byte, error) {
	l := len(data)
	if l == 0 {
		return nil, errors.New("invalid padding")
	}
	pad := int(data[l-1])
	if pad == 0 || pad > l {
		return nil, errors.New("invalid padding")
	}
	return data[:l-pad], nil
}

func DecryptResponse(
	encrypted string,
	consID string,
	secretKey string,
	timestamp string,
	useLZ bool,
) (string, error) {

	ciphertext, err := base64.StdEncoding.DecodeString(encrypted)
	if err != nil {
		return "", errors.New("invalid base64 response")
	}

	key := consID + secretKey + timestamp
	keyHash := sha256.Sum256([]byte(key))
	iv := keyHash[:16]

	block, err := aes.NewCipher(keyHash[:])
	if err != nil {
		return "", err
	}

	if len(ciphertext)%aes.BlockSize != 0 {
		return "", errors.New("ciphertext invalid length")
	}

	mode := cipher.NewCBCDecrypter(block, iv)
	plain := make([]byte, len(ciphertext))
	mode.CryptBlocks(plain, ciphertext)

	unpad, err := pkcs7Unpad(plain)
	if err != nil {
		return "", err
	}

	result := string(unpad)

	if useLZ {
		return lzstring.DecompressFromEncodedURIComponent(result)
	}

	return result, nil
}
