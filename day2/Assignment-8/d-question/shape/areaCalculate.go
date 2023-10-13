package shape

import "d-question/model"

var pi float32 = 3.14

func CalculateArea(r model.Circle) float32 {
	return pi * 2 * r.Radius * r.Radius
}
