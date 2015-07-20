package main

import "fmt"

func inSlice(x int, str string) bool {
	for _, value := range str {
		if x == int(value-'0') {
			return true
		}
	}
	return false
}

func sumOfDigits(str string) int {
	sum := 0
	for _, value := range str {
		sum += int(value - '0')
	}
	return sum
}

func productOfDigits(str string) int64 {
	product := int64(1)
	for _, value := range str {
		product *= int64(value - '0')
	}
	return product
}

func largestNdigitProduct(str string, digits int) int64 {
	product := int64(1)
	chunk := str[:digits]

	for i := 1; digits+i < len(str); i++ {
		p := productOfDigits(chunk)
		if p > product {
			product = p
		}
		chunk = str[i : digits+i]
	}
	return product
}

func main() {
	var str string
	fmt.Scan(&str)
	fmt.Println(largestNdigitProduct(str, 3))
	fmt.Println(largestNdigitProduct(str, 4))
	fmt.Println(largestNdigitProduct(str, 13))
}
