package main

import "fmt"

func primeFactors(x int) map[int]int {
	d := 2
	factors := map[int]int{}
	for x > 1 {
		for x%d == 0 {
			factors[d]++
			x /= d
		}

		d++
		if d*d > x {
			if x > 1 {
				factors[x]++
			}
			break
		}
	}
	return factors
}

func smallestPositiveEvenlyDivisibleNumber(min, max int) int {
	if min < 0 || max < 0 {
		return 0
	}
	if min > max {
		min, max = max, min
	}

	factors := map[int]int{}
	for i := min; i <= max; i++ {
		for j, value := range primeFactors(i) {
			if factors[j] < value {
				factors[j] = value
			}
		}
	}

	number := 1
	for i, value := range factors {
		for j := 0; j < value; j++ {
			number *= i
		}
	}
	return number
}

func main() {
	fmt.Println(smallestPositiveEvenlyDivisibleNumber(1, 20))
}
