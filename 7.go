package main

import (
	"fmt"
	"math"
)

func max(array []int) int {
	x := 0
	for _, value := range array {
		if value > x {
			x = value
		}
	}
	return x
}

func isPrime(x int, primes []int) bool {
	for _, value := range primes {
		if x%value == 0 {
			return false
		}
	}

	r := int(math.Sqrt(float64(x)))
	for i := max(primes) + 2; i <= r; i += 2 {
		if x%i == 0 {
			return false
		}
	}
	return true
}

func nthPrime(n int) int {
	if n == 1 {
		return 2
	}

	primes := []int{2}
	count := 1
	x := 3
	for ; ; x += 2 {
		if isPrime(x, primes) {
			primes = append(primes, x)
			count++
		}
		if count == n {
			break
		}
	}
	return x
}

func main() {
	fmt.Println(nthPrime(10001))
}
