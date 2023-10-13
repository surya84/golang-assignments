package main

import "fmt"

func divide(num1, num2 int) int {
	//defer fmt.Println("fdm")
	result := num1 / num2
	return result

}

func safeDivide() {
	divide(10, 2)
	msg := recover()

	if msg != nil {
		fmt.Println(msg)
		fmt.Println("recovered from panic")
	}
}

func main() {
	//safeDivide()
	defer fmt.Println("fdm")
	fmt.Println("exiting from main")
}
