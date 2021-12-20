package main

import (
	"fmt"
	randomnumber "randomnumber/Random"
)

func main() {
	nilai := randomnumber.RangeNumber(2, 10, 5)
	fmt.Println(nilai)
}
