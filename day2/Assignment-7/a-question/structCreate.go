package main

import "fmt"

type Employee struct {
	Name string
	Id   int
	Age  int
	City string
}

func main() {
	u := Employee{
		Name: "Ram",
		Id:   456821,
		Age:  32,
		City: "Bangalore",
	}

	u2 := Employee{
		Name: "Lakshman",
		Id:   543219,
		Age:  45,
		City: "Hyderabad",
	}

	fmt.Println("Printing Employee - 1 Details ---")
	fmt.Println("Name 	: ", u.Name)
	fmt.Println("Id 	: ", u.Id)
	fmt.Println("Age 	: ", u.Age)
	fmt.Println("City 	: ", u.City)
	fmt.Println("Printing Employee - 2 Details ---")
	fmt.Println("Name 	: ", u2.Name)
	fmt.Println("Id 	: ", u2.Id)
	fmt.Println("Age 	: ", u2.Age)
	fmt.Println("City 	: ", u2.City)

}
