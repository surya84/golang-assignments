package calculator

func Square(i int) int {
	return i * i
}

func Addition(i, y int) int {
	return i + y
}

func Substraction(i, y int) int {
	return i - y
}

func Multiplication(i, y int) int {
	return i * y
}

func DivMod(i, y int) (int, int) {
	return i % y, i / y
}
