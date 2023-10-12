package main

import (
	"fmt"
	"project2/temparature"
)

func main() {
	// fmt.Print("enter temparature in fahrenheit :")
	// reader := bufio.NewReader(os.Stdin)
	// input, _ := reader.ReadString('\n')

	// num, _ := strconv.ParseInt(strings.TrimSpace(input), 0, 64)

	var fahrenheit float64

	fmt.Print("Enter temparature in Fahrenheit : ")
	_, err := fmt.Scanf("%f", &fahrenheit)

	if err != nil {
		return
	}

	fmt.Println(temparature.ConvertTemparature(fahrenheit))

}
