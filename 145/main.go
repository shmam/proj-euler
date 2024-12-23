package main


import "math"
import "fmt"

/**
Some positive integers n have the property that the sum(n + reverse(n)) consists entirely of odd (decimal) digits. 
For instance, 36 + 63 = 99 and 409 + 904 = 1313. We will call such numbers reverseable; so 36, 63, 409, and 904 are reversible. 
Leading zeroes are not allowed in either n or reverse(n).
There are 120 reversible numbers below one-thousand.
How many reversible numbers are there below one-billion 10^9?
*/

func main() {
	m := make(map[float64]int)

	sum := 0 
	for i := 1.; i < math.Pow(10.,9.); i++ {

		if m[i] != 1 {
			r := reverse(i)
			if isOdd(i + r) {
				sum++
				m[i] = 1
				m[r] = 1
			}
		}
	}

	fmt.Println("solution: ", sum)
}

func reverse(num float64) float64 {
	numDigits := math.Floor(math.Log10(num)) + 1.

	sum := 0.
	for i := numDigits; i > 0; i-- {
		d := int(num) % 10
		num = math.Floor(float64(num / 10))
		sum += float64(d) * math.Pow(10., i - 1)
	}
	return sum
}

func isOdd(sum float64) bool {
	numDigits := math.Floor(math.Log10(sum)) + 1.

	for i := numDigits; i > 0; i-- {
		d := int(sum) % 10
		if d % 2 == 0 {
			return false
		}
		sum = math.Floor(float64(sum / 10))
	}

	return true
}