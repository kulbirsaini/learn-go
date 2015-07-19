package main

import "fmt"

func difference(max int) int {
	sumSquares := 0
	squaresSum := 0

	for i := 0; i <= max; i++ {
		sumSquares += i * i
		squaresSum += i
	}

	return squaresSum*squaresSum - sumSquares
}

func main() {
	fmt.Println(difference(100))
}
