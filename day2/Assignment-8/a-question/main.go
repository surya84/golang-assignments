package main

import (
	"fmt"
	"os"
)

func removeFile(file string) {
	err := os.Remove(file)
	if err != nil {
		fmt.Println("file error", err)
		return
	}
	fmt.Println("File Deleted")
}

func main() {

	var fileName string

	fmt.Println("Enter file name to delete : ")
	fmt.Scanln(&fileName)

	removeFile("a.txt")
}
