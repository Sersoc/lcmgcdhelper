package lcmgcdhelper

func CalculateGCD(x int, y int) int {
	if y == 0 {
		return x
	}
	return CalculateGCD(y, x%y)
}

func CalculateLCM(x int, y int) int {
	if x == 0 || y == 0 {
		return 0
	}
	z := CalculateGCD(x, y)
	return (x / z) * y
}
