package main

import "fmt"
// import "time"

func main() {
	// switch statement:
	// i := 4

	// switch i {
	// case 1:
	// 	fmt.Println("one")
	// case 2:
	// 	fmt.Println("two")
	// case 3:
	// 	fmt.Println("three")
	// default:
	// 	fmt.Println("other")
	// }


	// // Multiple condition switch
	// switch time.Now().Weekday() {
	// case time.Saturday, time.Sunday:
	// 	fmt.Println("Its off day")
	// default:
	// 	fmt.Println("Its Week day")
	// }


	// type switch
	whoAmI := func(i interface{}) {
		switch i.(type) {
		case int:
			fmt.Println("Integer")
		case float32:
			fmt.Println("Float")
		case string:
			fmt.Println("String")
		case bool:
			fmt.Println("Boolean")
		case complex64:
			fmt.Println("Complex")
		default:
			fmt.Println("Invalid")
		}
	}

	whoAmI("hello")
	whoAmI(2+22i)

}