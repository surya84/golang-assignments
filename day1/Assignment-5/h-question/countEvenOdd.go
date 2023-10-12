package main

import "fmt"

func main() {
	var numbers []int
	numbers = []int{2, 3, 4, 12, 54, 33, 45, 98, 21, 50}

	var evenCount, oddCount int

	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			evenCount++
		} else {
			oddCount++
		}
	}
	fmt.Println("The Number slice is : ", numbers)
	fmt.Println("No.of Even numbers in slice are : ", evenCount)
	fmt.Println("No.of Odd Numbers in slice are : ", oddCount)
}
