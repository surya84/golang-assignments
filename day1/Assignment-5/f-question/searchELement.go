package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func search(numbers []int, element int) bool {
	for i := 0; i < len(numbers); i++ {
		if element == numbers[i] {
			return true
		}
	}
	return false
}

func main() {
	var numbers []int = []int{1, 2, 3, 4, 5, 6, 7, 8}
	fmt.Println("The number Slice is : ", numbers)
	fmt.Print("Enter Number to Search : ")
	reader := bufio.NewReader(os.Stdin)

	input, _ := reader.ReadString('\n')

	num, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

	result := search(numbers, int(num))

	if result {
		fmt.Println("Element found")
	} else {
		fmt.Println("Element Not Found")
	}
}
