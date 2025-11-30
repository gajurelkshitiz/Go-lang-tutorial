package main

import "fmt"

func main() {
	// zeroed values


	// int --> 0
	var num[4] int

	num[0] = 1
	// fmt.Println(len(num))
	// fmt.Println(num[0])
	fmt.Println(num)



	// bool --> false
	var vals[4] bool
	fmt.Println(vals)
	
	vals[2] = true
	fmt.Println(vals)



	// string --> empty empty i.e. " "
	var name[3] string
	fmt.Println(name)

	name[0] = "golang"
	name[1] = "is"
	name[2] = "amazing!"
	fmt.Println(name)


	// initialize value during declaration of array:
	nums := [3] int {1, 2, 3}
	fmt.Println(nums)



	// 2D Array
	my_arr := [2][2] int {{1,2},{3,4}}
	fmt.Println(my_arr)
}