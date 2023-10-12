package temparature

func ConvertTemparature(temp float64) float64 {
	result := (temp - 32) * 5 / 9
	return result
}
