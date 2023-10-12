package main

import "fmt"

func main() {

	var numbers []int

	numbers = []int{1, 2, 3, 4, 5}

	var sum int

	var avg float64

	for i := 0; i < len(numbers); i++ {
		sum += numbers[i]
	}
	avg = (float64)(sum / len(numbers))

	fmt.Println("The sum of Numbers in slice is : ", sum)
	fmt.Println("The Average of numbers in slice is : ", avg)
}
