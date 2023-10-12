package main

import (
	"bufio"
	"fmt"
	"os"
	"project2/temparature"
	"strconv"
	"strings"
)

func main() {
	fmt.Print("enter temparature in fahrenheit :")
	reader := bufio.NewReader(os.Stdin)
	input, _ := reader.ReadString('\n')

	num, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

	fmt.Println(temparature.ConvertTemparature(int(num)))

	// var a int
	// fmt.Scanf("Enter : ")
	// fmt.Println(a)

}
