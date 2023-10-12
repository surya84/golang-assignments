package main

import "fmt"

func main() {
	var slice1 []int
	var slice2 []int

	slice1 = []int{100, 200, 300}
	slice2 = []int{400, 500, 600}

	slice1 = append(slice1, slice2...)

	fmt.Println("The slice1 after concatinating with slice2 : ", slice1)

}
