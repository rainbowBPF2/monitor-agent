package utill

import (
	"agent/conf"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"log"
)

// salt key length should limit to 8 char.
var saltKey []byte

func init() {
	saltKey = []byte("saltDataSaltData")
	cMap := ParseAppConf()

	if cMap != nil && cMap[conf.Encrypted] == "true" {
		saltKey = []byte(cMap[conf.EncryptedKey])
	}
}

// Encrypt data using AES algorithm
func EncryptAesData(textData string) (string, error) {
	block, err := aes.NewCipher(saltKey)

	if err != nil {
		log.Fatal(err.Error())
		return "", err
	}

	cipherText := make([]byte, aes.BlockSize+len(textData))
	initVector := cipherText[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, initVector); err != nil {
		return "", err
	}
	cipher.NewCFBEncrypter(block, initVector).XORKeyStream(cipherText[aes.BlockSize:], []byte(textData))
	return hex.EncodeToString(cipherText), nil

}

// Decrypt data using AES algorithm.
func DecryptAesData(encryptData string) (string, error) {
	cipherText, _ := hex.DecodeString(encryptData)

	block, err := aes.NewCipher(saltKey)
	if err != nil {
		return "", err
	}
	if len(cipherText) < aes.BlockSize {
		return "", errors.New("cipherText too short")
	}
	initVector := cipherText[:aes.BlockSize]
	cipherText = cipherText[aes.BlockSize:]
	cipher.NewCFBDecrypter(block, initVector).XORKeyStream(cipherText, cipherText)
	return string(cipherText), nil
}
