package main

import "fmt"

func Divide(a int, b int) int {
	if b == 0 {
		panic("denominator is zero")
	}
	return a / b
}

func safeDivide(a int, b int) int {
	defer errorRecover()

	return Divide(a, b)
}
func errorRecover() {
	err := recover()
	if err != nil {
		fmt.Println("Panic Issue...")
	}
}
func main() {

	fmt.Println(safeDivide(8, 0))
	fmt.Println(safeDivide(9, 3))
	fmt.Println("End of Main and Program")

}
