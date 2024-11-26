package lcmgcdhelper

import "fmt"

func CalculateGCD(x int, y int) int {
	var low int
	var high int
	if x > y {
		high = x
		low = y
	} else {
		high = y
		low = x
	}
	if low == 0 {
		return high

	}

	return CalculateGCD(low, high%low)

}

func CalculateLCM(x int, y int) int {
	defer func() {
		e := recover()
		fmt.Println(e)
	}()
	z := CalculateGCD(x, y)

	return (x / z) * (y / z) * z
}
