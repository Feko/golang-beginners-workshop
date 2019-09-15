package main

import "fmt"

type Animal interface {
	description() string
}

//There's no reference of the interface on struct declaration
type Cat struct {
	Type  string
	Sound string
}

type Snake struct {
	Type      string
	Poisonous bool
}

//Interface implementation doesn't reference interface either!
func (c Snake) description() string {
	return fmt.Sprintf("Is it poisonous? %v", c.Poisonous)
}

func (g Cat) description() string {
	return fmt.Sprintf("Sound: %v", g.Sound)
}

func main() {
	//Interface assignment and method invocation are assured in compile-time
	var a Animal
	a = Snake{Poisonous: true}
	fmt.Println(a.description())
	a = Cat{Sound: "Meowww!!!"}
	fmt.Println(a.description())
}
