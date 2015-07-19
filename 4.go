package main

import (
	"fmt"
	"math"
)

func isPalindrome(x int) bool {
	y, z := x, 0
	for y > 0 {
		z = z*10 + y%10
		y /= 10
	}
	return x == z
}

func numDigist(x int) int {
	digits := 0
	for x > 0 {
		x /= 10
		digits++
	}
	return digits
}

func makePalindrome(x int) int {
	y := x
	for y > 0 {
		x = x*10 + y%10
		y /= 10
	}
	return x
}

func factors(x, digits int) []int {
	d := int(math.Sqrt(float64(x)))
	min := power(10, digits-1)
	max := power(10, digits) - 1
	var f []int
	for i := min; i < d && d <= max; i++ {
		if x%i == 0 && x/i <= max && x/i >= min {
			f = append(f, i)
		}
	}
	return f
}

func power(x, y int) int {
	z := 1
	for i := 0; i < y; i++ {
		z *= x
	}
	return z
}

type Factors struct {
	number  int
	factor1 int
	factor2 int
}

func largestPalindrome(digits int) {
	if digits < 1 {
		return
	}
	n := power(10, digits) - 1
	n *= n
	d := power(10, numDigist(n)-digits) - 1
	n = n / d

	min := power(10, digits-1)
	for i := n; i > min; i-- {
		p := makePalindrome(i)
		f := factors(p, digits)
		for _, factor := range f {
			fmt.Printf("Palindrome: %d, Factor1: %d, Factor2: %d\n", p, factor, (p / factor))
			return
		}
	}
	return
}

func main() {
	for i := 0; i < 6; i++ {
		largestPalindrome(i)
	}
}
