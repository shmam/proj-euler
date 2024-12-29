package main

import (
	"fmt"
	"math"
)

func main() {
	var dsum uint64 = 0
	var nsum uint64 = 0
	var d uint64 = 1
	var n uint64 = 0
	for ; n < math.MaxInt64; n++ {
		dsum += fn(n, d)
		if n == dsum {
			nsum += n
		}
	}

	fmt.Println("s(1)", nsum)
}

func fn(n uint64, d uint64) uint64 {
	count := 0
	// take all of the digits in n
	numDigits := uint64(math.Floor(math.Log10(float64(n))) + 1)

	// count how many digits in n = d and add it to dsum
	for range numDigits {
		digit := uint64(math.Mod(float64(n), 10))

		if digit == d {
			count++
		}

		n = uint64((math.Floor(float64(n) / 10)))
	}
	return uint64(count)
}
