package main

import "fmt"

func factorial(num int) int {
	if num == 1 || num == 0 {
		return num
	}
	return num * factorial(num-1)
}

func main() {
	var num int
	fmt.Print("Enter number to calculate : ")
	fmt.Scanln(&num)
	fmt.Printf("Factorial of %d is : %d", num, factorial(num))
}
