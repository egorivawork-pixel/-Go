package main

import "fmt"

func averageAge(people map[string]int) float64 {
	sum := 0
	for _, age := range people {
		sum += age
	}
	return float64(sum) / float64(len(people))
}

func main() {
	people := map[string]int{
		"Аня":   20,
		"Игорь": 25,
		"Маша":  30,
	}

	fmt.Printf("средний возраст = %.2f\n", averageAge(people))
}
