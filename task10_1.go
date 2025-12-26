package main

import (
	"bufio"
	"crypto/md5"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"os"
	"strings"
)

func hash(algo, s string) (string, error) {
	b := []byte(s)
	switch strings.ToLower(algo) {
	case "md5":
		h := md5.Sum(b)
		return hex.EncodeToString(h[:]), nil
	case "sha256":
		h := sha256.Sum256(b)
		return hex.EncodeToString(h[:]), nil
	case "sha512":
		h := sha512.Sum512(b)
		return hex.EncodeToString(h[:]), nil
	default:
		return "", fmt.Errorf("unknown algo")
	}
}

func main() {
	in := bufio.NewReader(os.Stdin)

	fmt.Print("Алгоритм (md5/sha256/sha512): ")
	algo, _ := in.ReadString('\n')
	algo = strings.TrimSpace(algo)

	fmt.Print("Строка: ")
	s, _ := in.ReadString('\n')
	s = strings.TrimRight(s, "\r\n")

	h, err := hash(algo, s)
	if err != nil {
		fmt.Println("Ошибка:", err)
		return
	}
	fmt.Println("Хэш:", h)

	fmt.Print("Проверка: введи хэш для сравнения (или Enter чтобы пропустить): ")
	userHash, _ := in.ReadString('\n')
	userHash = strings.TrimSpace(userHash)
	if userHash != "" {
		if strings.EqualFold(userHash, h) {
			fmt.Println("Совпадает ✅")
		} else {
			fmt.Println("Не совпадает ❌")
		}
	}
}
