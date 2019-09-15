package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func main() {
	//Casting
	a := 1.1
	b := int(a)
	fmt.Println(a)
	fmt.Println(b)

	s := strconv.Itoa(b)
	fmt.Println(s)
	fmt.Println(reflect.TypeOf(s))

	//Conditionals
	i := 1
	if i >= 0 {
		fmt.Println("Positive number")
	} else {
		fmt.Println("Negative number")
	}

	//For loop
	for idx := 10; idx < 20; idx++ {
		fmt.Println(idx)
	}

	//For loop also accepts a statement returning true
	idx := 5
	for idx > 0 {
		fmt.Printf("Idx is now %d\n", idx)
		idx--
	}

	//There's no such thing as while in Go.
	// An empty for is the same as while true. So, the code bellow would run forever!
	/*
		for{
			fmt.Print("I just don't know when to stop!")
		}
	*/

	//RANGE iterates over an array, just like FOREACH in other languages
	var pow = []int{10, 356, 23, 42, 64, 234, 24, 5, 34, 5324}
	for i, v := range pow {
		fmt.Printf("Value in position %d = %d\n", i, v)
	}

	//Pointers
	var pp *int
	p := 12
	pp = &p
	fmt.Println(pp)
	fmt.Println(*pp)
	*pp = 13
	fmt.Println(p)
	fmt.Println(&p)

}
