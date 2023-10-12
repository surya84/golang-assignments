package main

import "fmt"

func main() {
	var fruits = make(map[string]int)

	fruits["Orange"] = 54
	fruits["Apple"] = 32
	fruits["Pine Apple"] = 23
	fruits["Mango"] = 89
	fruits["Grapes"] = 30
	fruits["Banana"] = 90

	fmt.Println("Fruit : Quantity")
	for key, value := range fruits {
		fmt.Println(key, " : ", value)
	}
}
