package armstrong

import (
	"math"
	"strconv"
)

func IsNumber(num int) bool {
	if num < 10 {
		return true
	}

	if num > 10 && num < 20 {
		return false
	}

	total := 0
	for _, x := range strconv.Itoa(num) {
		digit, _ := strconv.Atoi(string(x))
		total += int(math.Pow(float64(digit), float64(len(strconv.Itoa(num)))))
	}

	return total == num
}
