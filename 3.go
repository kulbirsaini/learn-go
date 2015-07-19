package main

import "fmt"

func primeFactors(x int) []int {
	var factors []int
	d := 2
	for x > 1 {
		for x%d == 0 {
			factors = append(factors, d)
			x /= d
		}
		d += 1
		if d*d > x {
			if x > 1 {
				factors = append(factors, x)
			}
			break
		}
	}
	return factors
}

func isValueInSlice(x int, list []int) bool {
	for _, value := range list {
		if value == x {
			return true
		}
	}
	return false
}

func uniquePrimeFactors(x int) []int {
	factors := primeFactors(x)
	var uniqueFactors []int
	for _, value := range factors {
		if !isValueInSlice(value, uniqueFactors) {
			uniqueFactors = append(uniqueFactors, value)
		}
	}
	return uniqueFactors
}

func largestPriceFactor(x int) int {
	factors := uniquePrimeFactors(x)
	largest := 1
	for _, value := range factors {
		if value > largest {
			largest = value
		}
	}
	return largest
}

func main() {
	fmt.Println(largestPriceFactor(600851475143))
}
