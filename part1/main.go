package main

import "fmt"

func main() {
	fmt.Println("Hello world!")
	//Variables
	i := 10
	c := 'S'

	//Arrays (Immutables)
	a := [...]int{1, 27, 11}
	var aa [3]int
	aa[1] = 5
	var aaa = [...]int{1, 2, 3, 4, 5, 6, 7, 8}

	fmt.Println(i)
	fmt.Println(c)
	fmt.Println(a)
	fmt.Println(aa)
	fmt.Println(len(aa))
	fmt.Println(len(aaa))

	//Slices - Same declaration as an Array
	var s []int
	ss := [5]int{1, 5, 7, 2, 9}
	sss := []string{"Feko", "Feeeeko", "Carpanos"}
	fmt.Println(s)
	fmt.Println(ss[1:4])
	fmt.Println(cap(s))
	fmt.Println(sss)

	//Slices are mutable, append!
	s = append(s, 1, 2)
	fmt.Println(s)

	//Maps -> Dictionaries
	var m map[string]int
	mm := make(map[string]bool)
	mm["www.google.com.br"] = true
	mm["github.com"] = true
	mm["facebook.com"] = false

	fmt.Println(m)
	fmt.Println(mm)

	//Conditionals
	if mm["github.com"] {
		fmt.Println("Are you a coder?")
	}

	//Delete from map
	delete(mm, "facebook.com")
	fmt.Println(mm)

}
