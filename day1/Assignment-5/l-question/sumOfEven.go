package main

import "fmt"

func calcSum(numbers []int) int {
	var sum int
	for i := 0; i < len(numbers); i++ {
		if numbers[i]%2 == 0 {
			sum += numbers[i]
		}

	}

	return sum
}
func main() {
	var numbers []int = []int{20, 43, 12, 42, 55, 36, 97, 50}

	result := calcSum(numbers)
	fmt.Println("Numbers slice is : ", numbers)
	fmt.Println("Sum of All Even Numbers in slice is : ", result)
}
