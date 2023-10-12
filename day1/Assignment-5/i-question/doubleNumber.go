package main

import "fmt"

func main() {
	var numbers []int

	numbers = []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := 0; i < len(numbers); i++ {
		numbers[i] = numbers[i] + numbers[i]
	}

	fmt.Println("Modified slice after doubling values are : ", numbers)
}
