package main

import "fmt"

type Student struct {
	Name string
	Age  int
}

func update(s *Student) {
	s.Age += 1
}

func main() {
	s1 := Student{
		Name: "Surya",
		Age:  22,
	}

	fmt.Println("Original Data od Student is : ")
	fmt.Println("Name : ", s1.Name)
	fmt.Println("Age  : ", s1.Age)

	update(&s1)
	fmt.Println("Student Data After incrementind Age with 1")
	fmt.Println("Name : ", s1.Name)
	fmt.Println("Age  : ", s1.Age)

}
