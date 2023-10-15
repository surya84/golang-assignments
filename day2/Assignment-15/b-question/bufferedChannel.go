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
	ch := make(chan int, 6)
	go sender(ch)
	go receiver(ch)
	time.Sleep(2 * time.Second)
	fmt.Println("End of main function and Program")
}
