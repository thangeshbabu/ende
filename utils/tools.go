package utils

import (
	"aes_encrypt/global"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"crypto/rand"
	"encoding/hex"
	"errors"
	"io"
	"log"
	"os"
)

func createHash(key string) string {

	hasher := md5.New()

	hasher.Write([]byte(key))

	return hex.EncodeToString(hasher.Sum(nil))

}

func Decrypt(data []byte, passphrase string) []byte {

	key := []byte(createHash(passphrase))

	block, err := aes.NewCipher(key)

	if err != nil {

		panic(err.Error())

	}

	gcm, err := cipher.NewGCM(block)

	if err != nil {

		panic(err.Error())

	}

	nonceSize := gcm.NonceSize()

	nonce, ciphertext := data[:nonceSize], data[nonceSize:]

	plaintext, err := gcm.Open(nil, nonce, ciphertext, nil)

	if err != nil {

		panic(err.Error())

	}

	return plaintext

}

func Encrypt(data []byte, passphrase string) []byte {

	block, _ := aes.NewCipher([]byte(createHash(passphrase)))

	gcm, err := cipher.NewGCM(block)

	if err != nil {

		panic(err.Error())

	}

	nonce := make([]byte, gcm.NonceSize())

	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {

		panic(err.Error())

	}

	ciphertext := gcm.Seal(nonce, nonce, data, nil)

	return ciphertext

}

func If_path_exist(path string) bool {

	if _, err := os.Stat(path); err == nil {

		return true
	} else if errors.Is(err, os.ErrNotExist) {
		return false

	} else {
		Error_check(err)

	}
	return false
}

func Error_check(err error) {

	if err != nil {
		global.Errorlog.Println(err)
		global.Errorlog.Fatal("\033[32m", err, "\033[0m")
	}
}

func Logging() {

	log.SetFlags(7)

	mw := io.MultiWriter(os.Stdout)

	// defining custom loggers which writes to both os.stdout and logfile at the same time.
	global.Errorlog = log.New(mw, "ERROR: ", 0)
	global.Infolog = log.New(mw, "Info: ", 0)

}
