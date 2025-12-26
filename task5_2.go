package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Info() {
	fmt.Printf("Имя: %s, Возраст: %d лет\n", p.name, p.age)
}

func (p *Person) Birthday() {
	p.age++
}

func main() {
	people := []Person{
		{name: "Саша", age: 20},
		{name: "Маша", age: 22},
		{name: "Игорь", age: 25},
	}

	fmt.Println("Все люди:")
	for _, person := range people {
		person.Info()
	}

	people[1].Birthday()

	fmt.Println("\nПосле дня рождения Маши:")
	for _, person := range people {
		person.Info()
	}
}
