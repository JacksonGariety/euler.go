package main

import "fmt"

func evenFibs(n int) int {
	a, b, sum := 1, 2, 0
	for b < 4000000 {
		if b%2 == 0 {
			sum += b
		}
		a, b = b, a+b
	}
	return sum
}

func main() {
	fmt.Println(evenFibs(4000000))
}
