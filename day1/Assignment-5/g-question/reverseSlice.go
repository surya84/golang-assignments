package main

import "fmt"

func main() {

	var numbers []int

	numbers = []int{1, 2, 3, 4, 5}

	fmt.Println("Original Number Slice is : ", numbers)

	start := 0
	end := len(numbers) - 1

	//var temp int
	for start < end {
		numbers[start], numbers[end] = numbers[end], numbers[start]
		start++
		end--
	}

	fmt.Println("Number slice after reversing : ", numbers)

}
