package main

import "fmt"

func main() {

	var number []int

	number = []int{10, 23, 75, 54, 21}

	var max int = 0

	for i := 1; i < len(number); i++ {
		if number[i] > number[max] {
			max = i
		}
	}
	fmt.Println(number[max])
}
