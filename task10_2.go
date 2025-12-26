package main

import (
	"bufio"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"io"
	"os"
	"strings"
)

func deriveKey(keyStr string) []byte {
	// для лабы ок: приводим к 32 байтам (AES-256). В реале делали бы PBKDF2/scrypt.
	k := []byte(keyStr)
	out := make([]byte, 32)
	copy(out, k)
	return out
}

func encrypt(key []byte, plaintext []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonce := make([]byte, gcm.NonceSize())
	if _, err := io.ReadFull(rand.Reader, nonce); err != nil {
		return "", err
	}

	ciphertext := gcm.Seal(nonce, nonce, plaintext, nil) // nonce кладём в начало
	return base64.StdEncoding.EncodeToString(ciphertext), nil
}

func decrypt(key []byte, b64 string) (string, error) {
	raw, err := base64.StdEncoding.DecodeString(b64)
	if err != nil {
		return "", err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return "", err
	}

	nonceSize := gcm.NonceSize()
	if len(raw) < nonceSize {
		return "", fmt.Errorf("bad data")
	}

	nonce := raw[:nonceSize]
	ciphertext := raw[nonceSize:]

	plain, err := gcm.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		return "", err
	}
	return string(plain), nil
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Секретный ключ: ")
	keyStr, _ := in.ReadString('\n')
	keyStr = strings.TrimSpace(keyStr)
	key := deriveKey(keyStr)

	fmt.Print("Строка для шифрования: ")
	msg, _ := in.ReadString('\n')
	msg = strings.TrimRight(msg, "\r\n")

	enc, err := encrypt(key, []byte(msg))
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Зашифровано (base64):", enc)

	dec, err := decrypt(key, enc)
	if err != nil {
		fmt.Println("Ошибка расшифровки:", err)
		return
	}
	fmt.Println("Расшифровка:", dec)
}
