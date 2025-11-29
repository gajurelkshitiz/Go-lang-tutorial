package main

import "fmt"

func main() {
	// defualt variable initialization
	// var name string = "golang"

	// Alternatively, it infers the type
	var name = "golang"

	// other examples
	var age = 23
	var is_active = true


	// shorthand syntax
	full_name := "kshitiz gajurel"


	// Note: So, we should use type (because it infers by defualt)
	// Ans: No, when we declare the variable initially, and then initialize the value later then in such case we need to mention the type.

	// eg:
	var score int

	score = 73


	fmt.Println(name)
	fmt.Println(age)
	fmt.Println(is_active)
	fmt.Println(full_name)
	fmt.Println(score)
}