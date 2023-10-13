package main

import "fmt"

type Rectangle struct {
	Width  int
	Height int
}

func (r Rectangle) calculateArea() int {
	return r.Height * r.Width
}

func (r Rectangle) calculatePerimeter() int {
	return 2 * (r.Height + r.Width)
}

func main() {
	r := Rectangle{
		Width:  5,
		Height: 10,
	}

	fmt.Println("Area of Rectangle is : ", r.calculateArea())
	fmt.Println("Perimeter of Rectangle is : ", r.calculatePerimeter())
}
