package main

import "fmt"

func fibonacci() func() int {
	x, y := 1, 0

	return func() int {
		x, y = y, x+y
		return x
	}
}

func main() {
	f := fibonacci()
	sum := 0
	max := int(4e6)
	for {
		n := f()
		if n >= max {
			break
		}
		if n%2 == 0 {
			sum += n
		}
	}
	fmt.Println(sum)
}
