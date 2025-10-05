package main

import (
	"fmt"
	"interface/internal"
)

func main() {

	fmt.Println("Hello, Go!")
	dog := internal.Doge{Name: "Buddy", Breed: "Golden Retriever", Age: 3}
	fmt.Println(dog.Speak())

	rect := internal.Rectangle{Width: 5, Height: 10}
	fmt.Println("Area:", rect.Area())
	fmt.Println("Perimeter:", rect.Perimeter())
	circ := internal.Circle{Radius: 7}
	fmt.Println("Area:", circ.Area())
	fmt.Println("Perimeter:", circ.Perimeter())
}
