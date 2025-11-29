package main

import "fmt"

func main() {

	// only looping construct in go: can be used as all other loops
	
	// to simulate while loop
	// i := 1
	// for i <= 3 {
	// 	fmt.Println("While Loop running")
	// 	i = i + 1
	// }

	// to simulate infinite loop
	// for {
	// 	fmt.Println("Infinite")
	// }


	// classical for loop
	// for i := 0; i < 5; i++ {
	// 	// fmt.Println("Classical For loop")
		

	// 	// we can use 'break' and 'continue' here.
		
	// 	if i == 2 {
	// 		continue
	// 	}
		
	// 	if i == 2 {
	// 		break
	// 	}

	// 	fmt.Println(i)
	// }

	
	// range
	for i := range 3 {
		fmt.Println(i)
	}

	fmt.Println("End of Program!")
}