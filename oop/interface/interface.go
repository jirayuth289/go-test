package oop

import "fmt"

type Speaker interface {
	Speak() string
}

type Dog struct {
	Name string
}

// Dog's implementation of the Speaker interface
func (d Dog) Speak() string {
	return "Woof!"
}

// Person struct
type Person struct {
	Name string
}

// Person's implementation of the Speaker interface
func (p Person) Speak() string {
	return "Hello!"
}

// function that accepts Speaker interface
func makeSound(s Speaker) {
	fmt.Println(s.Speak())
}

func TestInterface() {
	fmt.Println("--------------------------Abstract & Interface--------------------------")
	var animal [2]Speaker

	animal[0] = Dog{Name: "Buddy"}
	animal[1] = Person{Name: "Alice"}

	for i := 0; i < len(animal); i++ {
		makeSound(animal[i])
	}
}
