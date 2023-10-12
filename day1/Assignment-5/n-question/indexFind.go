package main

import "fmt"

func findIndex(numbers []int, element int) int {
	for i := 0; i < len(numbers); i++ {
		if element == numbers[i] {
			return i
		}
	}
	return -1
}

func main() {
	var numbers []int = []int{1, 2, 4, 5, 23, 12, 34, 14, 67}

	result := findIndex(numbers, 11)

	if result > 0 {
		fmt.Println("Element found at index : ", result)
	} else {
		fmt.Println("Element Not found")
	}
}
