package temparature

func ConvertTemparature(temp int) float64 {
	result := (temp - 32) * 5 / 9
	return float64(result)
}
