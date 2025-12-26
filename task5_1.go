package main

import "fmt"

type Person struct {
	name string
	age  int
}

func (p Person) Info() {
	fmt.Printf("Имя: %s, Возраст: %d лет\n", p.name, p.age)
}

func main() {
	person1 := Person{name: "Аня", age: 20}
	person2 := Person{name: "Игорь", age: 25}

	person1.Info()
	person2.Info()
}
