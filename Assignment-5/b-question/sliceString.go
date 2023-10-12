package main

import "fmt"

func main() {
	//str string
	var str = []string{"Orange", "Apple", "Mango", "Pine Apple", "Grapes"}

	for _, value := range str {
		fmt.Print(value, " ")
	}
}
