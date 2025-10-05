package internal

import "fmt"

type Doge struct {
	Name  string
	Breed string
	Age   int
}

func (d Doge) Speak() string {
	return "Woof! Woof! Name is " + d.Name + ", Breed is " + d.Breed + ", Age is " + fmt.Sprint(d.Age)
}
