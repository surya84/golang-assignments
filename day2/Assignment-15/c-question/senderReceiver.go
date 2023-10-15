package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	go func() {
		for i := 1; i <= 5; i++ {
			ch <- i
		}
	}()

}

func receiver(ch chan int) {
	go func() {
		for number := range ch {
			fmt.Println("Received : ", number)
		}

	}()

	// return

}

func main() {
	ch := make(chan int)

	for i := 1; i <= 3; i++ {
		go sender(ch)
	}

	for i := 1; i <= 2; i++ {
		go receiver(ch)
	}
	time.Sleep(2 * time.Second)
	fmt.Println("End of main function and Program")
}
