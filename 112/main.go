package main

import "math"
import "fmt"

/**
Working from left-to-right if no digit is exceeded by the digit to its left it is called an increasing number; for example, 134468.
Similarly if no digit is exceeded by the digit to its right it is called a decreasing number; for example, 66420.
We shall call a positive integer that is neither increasing nor decreasing a "bouncy" number; for example, 155349.

Clearly there cannot be any bouncy numbers below one-hundred, but just over half of the numbers below one-thousand (525) are bouncy. 
In fact, the least number for which the proportion of bouncy numbers first reaches 50% is 538.
Surprisingly, bouncy numbers become more and more common and by the time we reach 21780 the proportion of bouncy numbers is equal to 90%.
Find the least number for which the proportion of bouncy numbers is exactly 99%.
*/

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
		// fmt.Println(pct)
		start += largeSkip
	}

	for {
		pct = evalRange(start)
		if pct == target {
			break
		}
		// fmt.Println(pct)
		start -= 1
	}
	


	fmt.Println("solution: ", start, " ", pct)
	// solution:  1587000   0.99
}

func evalRange(topRange int) float64 {
	sum := 0.
	for i := range(topRange) + 1 {
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
	for _ = range(numDigits) {
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