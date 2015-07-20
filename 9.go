package main

import (
	"fmt"
	"math"
)

func inSlice(list []int, x int) bool {
	for _, value := range list {
		if value == x {
			return true
		}
	}
	return false
}

func appendIfNotPresent(list *[]int, x int) {
	if inSlice(*list, x) {
		return
	}
	*list = append(*list, x)
	return
}

func primeFactors(x int) []int {
	var f []int

	d := 2
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

func factors(x int) []int {
	r := int(math.Sqrt(float64(x)))
	var f []int

	for i := 2; i <= r; i++ {
		if x%i == 0 {
			f = append(f, i)
		}
	}
	return f
}

func areCoprime(x, y int) bool {
	if x > y {
		x, y = y, x
	}
	for _, value := range primeFactors(x) {
		if y%value == 0 {
			return false
		}
	}
	return true
}

// m & n are coprime
// m > n
// k is a natural number > 0
// m*m + m*n = s / (2*k)
// a = k * (m*m - n*n)
// b = k * 2*m*n
// c = k * (m*m + n*n)
func pythagoreanTriplet(sum int) (int, int, int) {
	if sum%2 != 0 {
		return 0, 0, 0
	}

	sum /= 2

	m, n, found := 1, 1, false

	for ; m <= int(math.Sqrt(float64(sum))); m++ {
		n = 1
		for ; n < m; n++ {
			if m*m+m*n == sum {
				found = true
				break
			}
		}
		if found {
			break
		}
	}

	if found {
		return m*m - n*n, 2 * m * n, m*m + n*n
	}
	fmt.Println("Not found")
	return 0, 0, 0
}

func main() {
	a, b, c := pythagoreanTriplet(1000)
	fmt.Println(a * b * c)
}
