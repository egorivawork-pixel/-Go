package main
import "fmt"

func main() {
	people := map[string]int{
		"Аня":   20,
		"Игорь": 25,
		"Маша":  30,
	}

	people["Саша"] = 22

	for name, age := range people {
		fmt.Printf("имя: %s, возраст: %d\n", name, age)
	}
}
