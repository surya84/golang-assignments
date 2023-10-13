package main

import (
	"fmt"
	"os"
	"strings"
)

func CountWords(str []string) int {
	return len(str)
}

func main() {
	file := `C:\Users\ORR Training 2\Desktop\golang-assignments\day2\Assignment-13\sample.txt`

	readFile, err := os.ReadFile(file)
	if err != nil {
		return
	}

	var str []string = strings.Fields(string(readFile))

	fmt.Println("Number of Words in the file is : ", CountWords(str))
}
