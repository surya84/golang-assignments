package main

import "fmt"

type Address struct {
	Street  string
	City    string
	ZipCode int
}

type Person struct {
	Name string
	Address
}

func main() {
	u1 := Person{
		Name: "Surya",
		Address: Address{
			Street:  "Vallabha Nagar",
			City:    "Bangalore",
			ZipCode: 517257,
		},
	}

	fmt.Println("Printing Person Details ...")
	fmt.Println("Name 	 : ", u1.Name)
	fmt.Println("Street  : ", u1.Street)
	fmt.Println("City 	 : ", u1.City)
	fmt.Println("ZipCode : ", u1.ZipCode)
}
