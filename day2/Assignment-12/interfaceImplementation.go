package main

import "fmt"

type Animal interface {
	MakeSound() string
}

type Dog struct {
}

type Cat struct {
}

func (d Dog) MakeSound() string {
	return "Dog sounds like Bow - Bow"
}

func (c Cat) MakeSound() string {
	return "Cat sounds like Meaw - Meaw "
}

func main() {
	d := Dog{}
	c := Cat{}

	fmt.Println(d.MakeSound())
	fmt.Println(c.MakeSound())
}
