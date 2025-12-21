package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"stringutils"
)

func main() {

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Введите строку: ")
	text, err := reader.ReadString('\n')
	if err != nil {
		fmt.Println("Ошибка ввода:", err)
		return
	}

	text = strings.TrimSpace(text)

	reversed := stringutils.Reverse(text)

	fmt.Printf("Перевернутая строка: %s\n", reversed)
}
