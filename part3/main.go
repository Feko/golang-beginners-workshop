package main

import "fmt"

//A normal func
func add(a int, b int) int {
	return a + b
}

//Func with returning member declared on its signature
func add2(a int, b int) (c int) {
	c = a + b
	return
}

//Func returning multiple items
func add3(a int, b int) (int, string) {
	return a + b, "OK"
}

func main() {
	fmt.Println(add(3, 4))
	fmt.Println(add2(1, 4))

	//Calling function that returns more then one member
	fmt.Println(add3(5, 5))
	result, message := add3(5, 5)
	fmt.Println(result, message)
}
