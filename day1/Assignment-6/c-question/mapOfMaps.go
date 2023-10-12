package main

import "fmt"

func main() {
	var studentsData = make(map[string]map[string]interface{})

	studentsData["Ram"] = map[string]interface{}{
		"Age":   23,
		"Grade": 92,
		"City":  "Hyderabad",
	}

	studentsData["Sita"] = map[string]interface{}{
		"Age":   22,
		"Grade": 85,
		"City":  "Bangalore",
	}

	studentsData["Lucky"] = map[string]interface{}{
		"Age":   34,
		"Grade": 64,
		"City":  "Pune",
	}

	studentsData["Raj"] = map[string]interface{}{
		"Age":   45,
		"Grade": 54,
		"City":  "Chennai",
	}

	for key, value := range studentsData {
		fmt.Println(key, " : ", value)
	}

}
