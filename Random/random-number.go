package randomnumber

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func RangeNumber(min, max, number int) []int {
	slice := make([]int, number)
	for i := 0; i < number-1; i++ {
		slice[i] = rand.Intn(max) + min
	}

	return slice
}
