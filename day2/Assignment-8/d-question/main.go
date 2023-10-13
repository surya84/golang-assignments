package main

import (
	"d-question/model"
	"d-question/shape"
	"fmt"
)

func main() {
	r := model.Circle{
		Radius: 3.5,
	}

	fmt.Println("Area of Circle is : ", shape.CalculateArea(r))

}
