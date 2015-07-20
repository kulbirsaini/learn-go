package main

import (
	"fmt"
	"math"
)

func inSlice(list []int64, x int64) bool {
	for _, value := range list {
		if value == x {
			return true
		}
	}
	return false
}

func appendIfNotPresent(list *[]int64, x int64) {
	if inSlice(*list, x) {
		return
	}
	*list = append(*list, x)
	return
}

func factors(x int64) []int64 {
	f := []int64{}
	d := int64(2)

	for x > 1 {
		for x%d == 0 {
			appendIfNotPresent(&f, d)
			x /= d
		}
		d++

		if d*d > x {
			if x > 1 {
				appendIfNotPresent(&f, x)
			}
			break
		}
	}
	return f
}

func isPrime(x int64) bool {
	if x == int64(2) {
		return true
	}

	if x < int64(2) || x%int64(2) == 0 {
		return false
	}

	limit := int64(math.Sqrt(float64(x)))

	for d := int64(3); d <= limit; d += int64(2) {
		if x%d == 0 {
			return false
		}
	}
	return true
}

func sumOfPrimesLessThan(x int64) int64 {
	if x < int64(3) {
		return int64(0)
	}

	var sum int64
	for i := int64(3); i < x; i += int64(2) {
		if isPrime(i) {
			sum += i
		}
	}
	return sum + int64(2)
}

func main() {
	fmt.Println(sumOfPrimesLessThan(int64(2e6)))
}
