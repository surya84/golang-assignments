package main

import (
	"fmt"
)

func compare(slice1, slice2 []int) bool {
	// sort.Ints(slice1)
	// sort.Ints(slice2)

	//var count bool = false

	if len(slice1) != len(slice2) {
		return false
	}
	for i := 0; i < len(slice1); i++ {
		if slice1[i] != slice2[i] {
			return false
		}
	}

	return true
	//return count

}

func main() {
	var slice1 []int = []int{1, 2, 3, 4, 5, 6}
	var slice2 []int = []int{1, 2, 3, 4, 5, 6}

	result := compare(slice1, slice2)

	if result {
		fmt.Println("Two slices are   equal")
	} else {
		fmt.Println("Two slices are  not equal")
	}
}
