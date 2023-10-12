package main

import (
	"fmt"
	"sort"
)

func findNonDuplicate(numbers []int) []int {

	if len(numbers) == 0 {
		return numbers
	}

	sort.Ints(numbers)

	result := []int{numbers[0]}

	for i := 1; i < len(numbers); i++ {
		if numbers[i] != numbers[i-1] {
			result = append(result, numbers[i])
		}
	}
	return result

}

func main() {
	var numbers []int

	numbers = []int{1, 2, 3, 4, 1, 3, 5, 6, 2}

	result := findNonDuplicate(numbers)

	fmt.Println(result)
}
