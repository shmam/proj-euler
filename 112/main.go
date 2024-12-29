package main

import (
	"fmt"
	"math"
)

func main() {
	target := float64(0.99)
	pct := 0.
	largeSkip := 1000
	start := 21780

	for {
		pct = evalRange(start)
		if pct >= target {
			break
		}
		start += largeSkip
	}

	for {
		pct = evalRange(start)
		if pct == target {
			break
		}
		start -= 1
	}

	fmt.Println("solution: ", start, " ", pct)
}

func evalRange(topRange int) float64 {
	sum := 0.
	for i := range (topRange) + 1 {
		b := isBouncy(i)
		if b {
			sum += 1
		}
	}

	return sum / float64(topRange)
}

func isBouncy(num int) bool {
	numDigits := int(math.Floor(math.Log10(float64(num))) + 1)

	incr := false
	decr := false

	// cannot be bouncy
	if num < 100 {
		return false
	}

	lastDigit := -1
	for _ = range numDigits {
		digit := num % 10

		if lastDigit != -1 {
			if lastDigit > digit {
				incr = true
			}
			if lastDigit < digit {
				decr = true
			}
		}

		num = int(math.Floor(float64(num / 10)))
		lastDigit = digit
	}
	return incr && decr
}
