package main

import "fmt"

func checkEvenOdd(num int) {
	if num%2 == 0 {
		fmt.Println("Number is Even")
	} else {
		fmt.Println("Number is Odd")
	}
}
func main() {
	checkEvenOdd(5)
}
