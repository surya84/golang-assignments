package main

import (
	"Assignment/calculator"
	"fmt"
)

func main() {
	fmt.Println("The Square of Number is ", calculator.Square(4))
	fmt.Println("The Addition of Two Numbers is : ", calculator.Addition(10, 20))
	fmt.Println("The Difference of two Numbers is :", calculator.Substraction(20, 10))
	fmt.Println("The Product of Two Numbers is : ", calculator.Multiplication(10, 20))
	//fmt.Printf("The Remainder of number is : %d  %d", calculator.DivMod(10, 20))
	fmt.Println(calculator.DivMod(10, 20))
}
