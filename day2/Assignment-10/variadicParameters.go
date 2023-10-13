package main

import (
	"fmt"
)

func sum(nums ...int) int {
	fmt.Print("Number slice : ", nums)
	totalSum := 0

	for _, num := range nums {
		totalSum += num
	}
	return totalSum
	//fmt.Println(totalSum)
}

func main() {
	fmt.Println(" 	sum : ", sum(1, 2, 3))
	fmt.Println("	sum : ", sum(1, 2, 4, 6, 8))

	number := []int{10, 20, 30}
	fmt.Println("	sum : ", sum(number...))
}
