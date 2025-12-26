package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {

	reader := bufio.NewReader(os.Stdin)
	fmt.Print("введите строку: ")
	text, _ := reader.ReadString('\n')

	upper := strings.ToUpper(text)

	fmt.Println("строка в верхнем регистре:", upper)
}
