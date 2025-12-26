package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"os"
)

func savePrivate(path string, key *rsa.PrivateKey) error {
	b := x509.MarshalPKCS1PrivateKey(key)
	return os.WriteFile(path, pem.EncodeToMemory(&pem.Block{Type: "RSA PRIVATE KEY", Bytes: b}), 0600)
}

func savePublic(path string, key *rsa.PublicKey) error {
	b := x509.MarshalPKCS1PublicKey(key)
	return os.WriteFile(path, pem.EncodeToMemory(&pem.Block{Type: "RSA PUBLIC KEY", Bytes: b}), 0644)
}

func main() {
	// 1) генерим ключи и сохраняем
	priv, err := rsa.GenerateKey(rand.Reader, 2048)
	if err != nil {
		panic(err)
	}
	_ = savePrivate("private.pem", priv)
	_ = savePublic("public.pem", &priv.PublicKey)
	fmt.Println("Ключи сохранены: private.pem, public.pem")

	// 2) подписываем сообщение
	message := []byte("hello from lab10")
	hash := sha256.Sum256(message)

	signature, err := rsa.SignPSS(rand.Reader, priv, crypto.SHA256, hash[:], nil)
	if err != nil {
		panic(err)
	}
	_ = os.WriteFile("signature.bin", signature, 0644)
	fmt.Println("Подпись сохранена: signature.bin")

	// 3) проверяем подпись
	err = rsa.VerifyPSS(&priv.PublicKey, crypto.SHA256, hash[:], signature, nil)
	if err != nil {
		fmt.Println("Подпись НЕвалидна ❌", err)
		return
	}
	fmt.Println("Подпись валидна ✅")
}
