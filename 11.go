package main

import (
	"fmt"
	"strconv"
)

func productOfDigits(grid *[][]int, digits int) int64 {
	product := int64(1)

	return product
}

func largestProduct(grid [][]int, digits int) int {
	product := 1
	p := product

	width := len(grid)
	height := len(grid[0])
	for i := 0; i < width; i++ {
		for j := 0; j < height; j++ {
			if i+digits < width {
				p = 1
				for k := 0; k < digits; k++ {
					p *= grid[i+k][j]
				}

				if p > product {
					product = p
				}
			}

			if j+digits < height {
				p = 1
				for k := 0; k < digits; k++ {
					p *= grid[i][j+k]
				}

				if p > product {
					product = p
				}
			}

			if i+digits < width && j+digits < height {
				p = 1
				for k := 0; k < digits; k++ {
					p *= grid[i+k][j+k]
				}

				if p > product {
					product = p
				}
			}

			if i-digits >= 0 && j+digits < height {

				p = 1
				for k := 0; k < digits; k++ {
					p *= grid[i-k][j+k]
				}

				if p > product {
					product = p
				}
			}
		}
	}

	return product
}

func main() {
	x, y := 0, 0
	fmt.Scan(&x)
	fmt.Scan(&y)

	var str string
	var tempInt int64
	var err error
	grid := make([][]int, x)
	for i := range grid {
		grid[i] = make([]int, y)
		for j := range grid[i] {
			fmt.Scan(&str)
			tempInt, err = strconv.ParseInt(str, 10, 32)
			if err == nil {
				grid[i][j] = int(tempInt)
			}
		}
	}

	fmt.Println(largestProduct(grid, 4))
}
