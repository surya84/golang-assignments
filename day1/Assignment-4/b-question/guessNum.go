package main

import (
	"fmt"
	"math/rand"
)

func main() {
	var num int = rand.Intn(10)
	var a int

	//num := rand.Intn(10)

	for {
		fmt.Print("Enter any number below 10 : ")
		_, err := fmt.Scanln(&a)
		if err != nil {
			fmt.Println("give correct values")
			return
		}

		if a == num {
			fmt.Println("Your Number :", a)
			fmt.Println("Random Number :", num)
			fmt.Println("you guessed it right")
			return
		} else if a > num {
			fmt.Println("too high")
		} else {
			fmt.Println("too low")
		}
	}

}
