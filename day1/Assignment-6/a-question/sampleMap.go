package main

import "fmt"

func main() {
	var studentGrades = make(map[string]float64)

	studentGrades["Ram"] = 9.8
	studentGrades["sita"] = 8.5
	studentGrades["lucky"] = 8.3
	studentGrades["raj"] = 6.2
	studentGrades["raghav"] = 4.3

	fmt.Println("Map Data Before Deleting one record")
	for k, v := range studentGrades {
		fmt.Println(k, " : ", v)
	}

	fmt.Println("Map Data After Deleting one record")

	delete(studentGrades, "raghav")
	for k, v := range studentGrades {
		fmt.Println(k, " : ", v)
	}
}
