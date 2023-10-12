package main

import "fmt"

func deleteElement(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}

func main() {
	var numbers []int

	//Adding numbers to slice
	numbers = append(numbers, 5)
	fmt.Println(numbers)

	numbers = append(numbers, 10)
	fmt.Println(numbers)
	numbers = append(numbers, 15)
	fmt.Println(numbers)
	numbers = append(numbers, 20)
	fmt.Println(numbers)
	numbers = append(numbers, 25)
	fmt.Println(numbers)

	fmt.Println("length of number slice is : ", len(numbers))
	fmt.Println("Capacity of number slice is : ", cap(numbers))
	var target int
	for i := 0; i < len(numbers); i++ {
		if numbers[i] == 15 {
			target = i
		}
	}
	numbers = deleteElement(numbers, target)
	fmt.Println(numbers)

}
