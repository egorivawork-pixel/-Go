package main

import "fmt"

func main() {

	people := map[string]int{
		"Аня":   20,
		"Игорь": 25,
		"Маша":  30,
	}

	var name string
	fmt.Print("введите имя для удаления: ")
	fmt.Scan(&name)

	delete(people, name)

	for n, age := range people {
		fmt.Printf("имя: %s, возраст: %d\n", n, age)
	}
}
