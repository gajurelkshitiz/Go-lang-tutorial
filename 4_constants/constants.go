package main

import "fmt"

const name = "golang"
			// name := "golang"  --> X this is not valid

// var age = 23


func main() {
	// const name string= "golang"
	// const name = "golang"

	// example of constant group for initializing multiple constants

	const (
		port = 5000
		host = "localhost"
	)

	fmt.Println(name)
	// fmt.Println(age)

	fmt.Println(port, host)
}