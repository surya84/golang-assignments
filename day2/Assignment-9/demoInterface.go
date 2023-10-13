package main

import "fmt"

var pi float32 = 3.14

type Shape interface {
	Area()
}

type Circle struct {
	Radius float32
}

type Rectangle struct {
	length float32
	width  float32
}

func (c Circle) Area() float32 {

	return pi * c.Radius * c.Radius * 2
}

func (r Rectangle) Area() float32 {
	return r.length * r.width
}

func main() {
	c := Circle{
		Radius: 6.5,
	}

	r := Rectangle{
		length: 20.3,
		width:  20,
	}

	fmt.Println("Area of Circle is : ", c.Area())
	fmt.Println("Area of Rectangle is : ", r.Area())
}
