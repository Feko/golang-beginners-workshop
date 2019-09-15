package main

import "fmt"

//That's a little tricky: If a struct name begins with a capital letter,
//then the struct is PUBLIC. Else, it's PRIVATE
//This rule applies to funcs as well!

//Usuario: Struct representing an user
type Usuario struct {
	Id    int
	Login string
	Email string
}

//This method belongs to the Struct, even though it is outside the Struct definition
func (u *Usuario) description() {
	fmt.Printf("User with ID %v has a login: %v ", u.Id, u.Login)
}

//Must use pointer declaration, or struct isn't passed by reference
func (u *Usuario) setId() {
	u.Id = 1000
}

func main() {
	u1 := Usuario{1, "feko", "someemail@example.com"}
	u2 := Usuario{Id: 2, Email: "justanother@email.com"}
	fmt.Println(u1)
	fmt.Println(u2)
	fmt.Println(u1.Email, u2.Email)

	u1.setId()       // Changes inside the func declaration are preserved
	u1.description() // Access struct's method normally, besides it's outside its definition

}
